// responsible for preparing GenoMagic to perform assemblies by pulling all the necessary
// Docker containers from DockerHub
package prepper

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/genomagic/constants"
)

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
	return p.prep(constants.AvailableAssemblers[constants.MegaHit])
}

// prep pulls and creates the container of the given docker image link
func (p *prep) prep(a *constants.AssemblerDetails) error {
	if a == nil {
		return fmt.Errorf("prep given nil AssemblerDetails")
	}
	reader, err := p.dClient.ImagePull(p.ctx, a.DHubURL, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image from DockerHub, err: %v", err)
	}
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("failed to copy stdout to internal reader, err: %v", err)
	}
	return nil
}
