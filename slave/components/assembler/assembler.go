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

	"github.com/genomagic/slave/components"
)

const (
	MegaHit = "megahit"
)

// AvailableAssemblers is a slice of assemblers that are currently supported
var AvailableAssemblers = map[string]bool{MegaHit: true}

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
	if !AvailableAssemblers[assembler] {
		return nil, fmt.Errorf("assembler not recognized, available assemblers: %v", AvailableAssemblers)
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

	if exists, err := a.imageExists(); err != nil {
		return nil, err
	} else if exists {
		a.dImageID = assembler
	} else {
		return nil, fmt.Errorf("failed to find Dockeri image for assembler: %s", assembler)
	}
	return a, nil
}

// imageExists attempts to find the Docker image of the assembler that is passed to NewAssembler
func (a *asmbler) imageExists() (bool, error) {
	images, err := a.dClient.ImageList(a.ctx, types.ImageListOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to get available Docker images, err: %v", err)
	} else if len(images) == 0 {
		return false, fmt.Errorf("imageExists found no images")
	}
	found := false
	for _, im := range images {
		if found {
			break
		}
		for _, tag := range im.RepoTags {
			if strings.Contains(tag, a.assemblerName) {
				found = true
			}
		}
	}
	if !found {
		return false, fmt.Errorf("failed to find a Docker container for the given assembler")
	}
	return found, nil
}

// Process performs the work of the assembler
func (a *asmbler) Process() error {
	// TODO: need to pass appropriate params to the MegaHit container
	resp, err := a.dClient.ContainerCreate(a.ctx, &container.Config{
		Tty:     true,
		Image:   a.assemblerName,
		Cmd:     []string{""},
		Volumes: map[string]struct{}{},
	}, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container, err: %v", err)
	}

	if err := a.dClient.ContainerStart(a.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container, err: %v", err)
	}

	if status, err := a.dClient.ContainerWait(a.ctx, resp.ID); err != nil {
		return fmt.Errorf("failed to wait for container to start up, err: %v", err)
	} else if status != 1 {
		return fmt.Errorf("received status code != 1, status code: %d", status)
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
