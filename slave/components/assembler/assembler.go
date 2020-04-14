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
	"github.com/docker/docker/client"

	"github.com/genomagic/constants"
	"github.com/genomagic/slave/components"
)

// structure of the assembler
type asmbler struct {
	// assembler type e.g MegaHit
	assemblerName string
	// path to the file the assembler will operate on
	filePath string
	// the Docker client the assembler will use to spin up containers
	dClient *client.Client
	// the image ID of the struct assembler
	dImageID string
	// context of requests performed to the Docker daemon
	ctx context.Context
}

// NewAssembler returns a new assembler for the specified file
func NewAssembler(filePath, assembler string) (components.Component, error) {
	if constants.AvailableAssemblers[assembler] == nil {
		return nil, fmt.Errorf("assembler not recognized")
	}

	a := &asmbler{
		assemblerName: assembler,
		filePath:      filePath,
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
		Cmd:     []string{""}, // TODO: need to pass appropriate params to the MegaHit container
		Volumes: map[string]struct{}{},
	}
	resp, err := a.dClient.ContainerCreate(a.ctx, ctConfig, nil, nil, "")
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
	return nil
}
