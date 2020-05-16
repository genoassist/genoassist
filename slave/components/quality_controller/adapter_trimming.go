package quality_controller

import (
	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
)

// adapterTrimming is the struct representation of the adapter trimming process
type adapterTrimming struct {
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// genoConfig is the GenoMagic global configuration
	genoConfig *config_parser.Config
}

// NewAdapterTrimming constructs and returns a new instance of adapterTrimming, which implements the Controller interface
func NewAdapterTrimming(d *client.Client, c *config_parser.Config) Controller {
	return &adapterTrimming{
		dockerCLI:  d,
		genoConfig: c,
	}
}

// Process performs the adapter trimming process by launching a Docker container that uses Trimmomatic
func (a *adapterTrimming) Process() (string, error) {
	return "", nil
}
