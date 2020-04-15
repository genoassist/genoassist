// contains the definition and work associated with assemblers
package assembler

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/genomagic/constants"
	"github.com/genomagic/slave/components"
)

// structure of the assembler
type asmbler struct {
	assemblerName string          // assembler type e.g MegaHit
	filePath      string          // path to the file the assembler will operate on
	outPath       string          // path to the directory where results are stored
	dClient       *client.Client  // the Docker client the assembler will use to spin up containers
	dImageID      string          // the image ID of the struct assembler
	ctx           context.Context // context of requests performed to the Docker daemon
}

// NewAssembler returns a new assembler for the specified file
func NewAssembler(fp, op, am string) (components.Component, error) {
	if constants.AvailableAssemblers[am] == nil {
		return nil, fmt.Errorf("assembler not recognized")
	}

	a := &asmbler{
		assemblerName: am,
		filePath:      fp,
		outPath:       op,
		ctx:           context.Background(),
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Docker client, err: %v", err)
	} else {
		a.dClient = cli
	}

	if assemblerID, err := a.getImageID(); err != nil {
		return nil, err
	} else {
		a.dImageID = assemblerID
	}
	return a, nil
}

// getImageID attempts to find the Docker image of the assembler that is passed to NewAssembler
func (a *asmbler) getImageID() (string, error) {
	images, err := a.dClient.ImageList(a.ctx, types.ImageListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get available Docker images, err: %v", err)
	} else if len(images) == 0 {
		return "", fmt.Errorf("getImageID found no images")
	}
	found := false
	assemblerID := ""
	for _, im := range images {
		if found {
			break
		}
		for _, tag := range im.RepoTags {
			if strings.Contains(tag, a.assemblerName) {
				found = true
				assemblerID = im.ID
			}
		}
	}
	if !found {
		return "", fmt.Errorf("failed to find a Docker container for the given assembler")
	}
	return assemblerID, nil
}

// Process performs the work of the assembler
func (a *asmbler) Process() error {
	ctConfig := &container.Config{
		Tty:     true,
		Image:   a.dImageID,
		Cmd:     constants.AvailableAssemblers[a.assemblerName].Comm(a.filePath, a.outPath),
		Volumes: map[string]struct{}{},
	}
	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{ // Binding the file provided by the user to the docker container
				Type:		mount.TypeBind,
				Source:		a.filePath,
				Target:		"/raw_sequence_input.fastq",
				ReadOnly:	false,
			},
			{ // Binding the output directory provided by the user to the docker container
				Type:		mount.TypeBind,
				Source:		a.outPath,
				Target:		"/output",
				ReadOnly:	false,
			},
		},
	}

	resp, err := a.dClient.ContainerCreate(a.ctx, ctConfig, hostConfig, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container, err: %v", err)
	}

	if err := a.dClient.ContainerStart(a.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container, err: %v", err)
	}

	if _, err = a.dClient.ContainerWait(a.ctx, resp.ID); err != nil {
		return fmt.Errorf("failed to wait for container to start up, err: %v", err)
	}

	out, err := a.dClient.ContainerLogs(a.ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(os.Stdout, out); err != nil {
		return fmt.Errorf("failed to capture stdout from Docker assembly container, err: %v", err)
	}
	fmt.Printf("[GenoMagic] Megahit run was complete. Your output is stored at: %s",a.outPath)
	return nil
}
