package quality_controller

import (
	"context"

	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
)

// adapterTrimming is the struct representation of the adapter trimming process
type adapterTrimming struct {
	// context of the process
	ctx context.Context
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// config is the GenoMagic global configuration
	config *config_parser.Config
}

// NewAdapterTrimming constructs and returns a new instance of adapterTrimming, which implements the Controller interface
func NewAdapterTrimming(ctx context.Context, dockerCli *client.Client, config *config_parser.Config) Controller {
	return &adapterTrimming{
		ctx:       ctx,
		dockerCLI: dockerCli,
		config:    config,
	}
}

// Process performs the adapter trimming process by launching a Docker container that uses Trimmomatic
func (a *adapterTrimming) Process() (string, error) {
	return "", nil
}
