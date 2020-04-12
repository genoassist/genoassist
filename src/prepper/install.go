// responsible for preparing GenoMagic to perform assemblies by pulling all the necessary
// Docker containers from DockerHub
package prepper

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	MegaHit = "docker.io/vout/megahit" // https://github.com/voutcn/megahit
)

// Assemblers is a mapping from an assembler DockerHub image link to an image name
// Used for checking that allowed assembler links are passed to prep()
var Assemblers = map[string]string{
	MegaHit: "megahit",
}

// prep holds the Docker client for pulling images
type prep struct {
	ctx     context.Context // context of the requests
	dClient *client.Client  // Docker client
}

// NewPrep initializes a prep struct and
func NewPrep() error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("failed to initialize Docker client, err: %v", err)
	}
	p := &prep{
		ctx:     ctx,
		dClient: cli,
	}
	return p.prep(MegaHit)
}

// prep pulls and creates the container of the given docker image link
func (p *prep) prep(dImageLink string) error {
	if Assemblers[dImageLink] == "" {
		return fmt.Errorf("passed assembler DockerHub link not recognized, allowed links: %v\n", Assemblers)
	}
	reader, err := p.dClient.ImagePull(p.ctx, dImageLink, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image from DockerHub, err: %v", err)
	}
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("failed to copy stdout to internal reader, err: %v", err)
	}

	resp, err := p.dClient.ContainerCreate(p.ctx, &container.Config{
		Tty:   true,
		Image: Assemblers[dImageLink],
	}, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container, err: %v", err)
	}

	if err := p.dClient.ContainerStart(p.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container, err: %v", err)
	}

	if status, err := p.dClient.ContainerWait(p.ctx, resp.ID); err != nil {
		return fmt.Errorf("failed to wait for container to start up, err: %v", err)
	} else if status != 1 {
		return fmt.Errorf("received status code != 1, status code: %d", status)
	}
	return nil
}
