// responsible for preparing GenoMagic to perform assemblies by pulling all the necessary
// Docker containers from DockerHub
package prepper

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/constants"
)

// prep holds the Docker client for pulling images
type prep struct {
	// context of the requests
	ctx context.Context
	// Docker client
	dockerCLI *client.Client
}

// NewPrep attempts to install all the necessary Docker images for GenoMagic. NewPrep launches go routines for
// installing the necessary images and collects the errors in a channel. When the go routines are finished, an error
// channel is returned, with the consumer being responsible to report whether errors have occurred and alert users about
// whether a specific assembler will be skipped
func NewPrep(config *config_parser.Config) chan error {
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

	// we are launching the pulling of Docker containers in go routines, but we have to wait for them to finish
	// in order to return the final error channel
	var wg sync.WaitGroup
	wg.Add(len(constants.AvailableAssemblers) + len(constants.AvailableQualityControllers))

	for _, availableAssembler := range constants.AvailableAssemblers {
		go func(ad *constants.AssemblerDetails) {
			errs <- p.prep(ad)
			wg.Done()
		}(availableAssembler)
	}

	for _, availableQualityController := range constants.AvailableQualityControllers {
		go func(aqc *constants.QualityControllerDetails) {
			errs <- p.prep(aqc)
			wg.Done()
		}(availableQualityController)
	}
	wg.Wait()

	return errs
}

// prep pulls and creates the container of the given docker image link
func (p *prep) prep(d constants.Details) error {
	if d == nil {
		return fmt.Errorf("prep given nil details")
	}
	reader, err := p.dockerCLI.ImagePull(p.ctx, d.GetDockerURL(), types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image from DockerHub, err: %s", err)
	}
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		return fmt.Errorf("failed to copy stdout to internal reader, err: %s", err)
	}
	if err := reader.Close(); err != nil {
		return fmt.Errorf("failed to close the reader, err: %s", err)
	}
	return nil
}
