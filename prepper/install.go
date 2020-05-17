// responsible for preparing GenoMagic to perform assemblies by pulling all the necessary
// Docker containers from DockerHub
package prepper

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/genomagic/config_parser"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/genomagic/constants"
)

// prep holds the Docker client for pulling images
type prep struct {
	// context of the requests
	ctx context.Context
	// Docker client
	dockerCLI *client.Client
}

// New attempts to install all the necessary Docker images for GenoMagic. New launches go routines for installing
// the necessary images and collects the errors in a channel. When the go routines are finished, an error channel
// is returned, with the consumer being responsible to report whether errors have occurred and alert users about
// whether a specific assembler will be skipped
func New(config *config_parser.Config) chan error {
	ctx := context.Background()
	errs := make(chan error, len(constants.AvailableAssemblers))
	if !config.GenoMagic.Prep {
		return errs
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		errs <- fmt.Errorf("failed to initialize Docker client, err: %v", err)
		return errs
	}
	p := &prep{
		ctx:       ctx,
		dockerCLI: cli,
	}

	// TODO: iterate over necessary quality control Docker images as well
	// we are launching the pulling of Docker containers in go routines, but we have to wait for them to finish
	// in order to return the final error channel
	var wg sync.WaitGroup
	wg.Add(len(constants.AvailableAssemblers))
	for _, aa := range constants.AvailableAssemblers {
		go func(a *constants.AssemblerDetails) {
			errs <- p.prep(a)
			wg.Done()
		}(aa)
	}
	wg.Wait()
	return errs
}

// prep pulls and creates the container of the given docker image link
func (p *prep) prep(a *constants.AssemblerDetails) error {
	if a == nil {
		return fmt.Errorf("prep given nil AssemblerDetails")
	}
	reader, err := p.dockerCLI.ImagePull(p.ctx, a.DHubURL, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image from DockerHub, err: %v", err)
	}
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("failed to copy stdout to internal reader, err: %v", err)
	}
	return nil
}
