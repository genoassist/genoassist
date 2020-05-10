// contains the definition and work associated with assemblers
package assembler

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/genomagic/constants"
	"github.com/genomagic/result"
	"github.com/genomagic/slave/components"
)

// structure of the assembler
type asmbler struct {
	assemblerName string          // assembler type e.g MegaHit
	filePath      string          // path to the file the assembler will operate on
	outPath       string          // path to the directory where results are stored
	numThreads    int             // number of threads to use for assemblers
	dClient       *client.Client  // the Docker client the assembler will use to spin up containers
	dImageID      string          // the image ID of the struct assembler
	ctx           context.Context // context of requests performed to the Docker daemon
}

// New returns a new assembler for the specified file
func New(fp, op, am string, thrds int) (components.Component, error) {
	if constants.AvailableAssemblers[am] == nil {
		return nil, fmt.Errorf("assembler not recognized")
	}

	a := &asmbler{
		assemblerName: am,
		filePath:      fp,
		outPath:       op,
		numThreads:    thrds,
		ctx:           context.Background(),
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
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

// getImageID attempts to find the Docker image of the assembler that is passed to New
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
func (a *asmbler) Process() (result.Result, error) {
	ctConfig := &container.Config{
		Tty:     true,
		Image:   a.dImageID,
		Cmd:     constants.AvailableAssemblers[a.assemblerName].Comm(),
		Volumes: map[string]struct{}{},
	}

	// we have assemblers that have special conditions, such as creating non-existing folders
	// run the condition functions before mounting
	if err := a.applyConditions(); err != nil {
		return nil, err
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{ // Binding the file provided by the user to the docker container
				Type:     mount.TypeBind,
				Source:   a.filePath,
				Target:   constants.RawSeqIn,
				ReadOnly: false,
			},
			{ // Binding the output directory provided by the user to the docker container
				Type:     mount.TypeBind,
				Source:   a.outPath,
				Target:   constants.BaseOut,
				ReadOnly: false,
			},
		},
	}

	resp, err := a.dClient.ContainerCreate(a.ctx, ctConfig, hostConfig, nil, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create container, err: %v", err)
	}

	if err := a.dClient.ContainerStart(a.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, fmt.Errorf("failed to start container, err: %v", err)
	}

	statCh, errCh := a.dClient.ContainerWait(a.ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return nil, fmt.Errorf("failed to wait for container to start up, err: %v", err)
		}
	case <-statCh:
	}

	out, err := a.dClient.ContainerLogs(a.ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return nil, fmt.Errorf("failed to get container logs, err: %v", err)
	}

	if _, err := io.Copy(os.Stdout, out); err != nil {
		return nil, fmt.Errorf("failed to capture stdout from Docker assembly container, err: %v", err)
	}
	return nil, nil
}

// applyConditions iterates over the conditions of a specific assemblers and attempts to fulfil the specific conditions
func (a *asmbler) applyConditions() error {
	if constants.AvailableAssemblers[a.assemblerName].ConditionsPresent {
		for i, f := range constants.AvailableAssemblers[a.assemblerName].Conditions {
			switch f {
			case constants.CreateDir:
				out := constants.AvailableAssemblers[a.assemblerName].OutputDir
				if err := os.Mkdir(path.Join(a.outPath, out), 0755); err != nil {
					return fmt.Errorf("failed to fulfil condition %d for assembler %s, err: %v", i, a.assemblerName, err)
				}
			}
		}
	}
	return nil
}
