package quality_controller

import (
	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
)

// decontamination is the representation of the decontamination process
type decontamination struct {
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// config is the GenoMagic global configuration
	config *config_parser.Config
	// toDecontaminate represents the path to the file to decontaminate
	toDecontaminate string
}

// NewDecontamination constructs and returns a new decontamination struct, which implements the Controller interface
func NewDecontamination(dockerCli *client.Client, config *config_parser.Config, fileToDecontaminate string) Controller {
	return &decontamination{
		dockerCLI:       dockerCli,
		config:          config,
		toDecontaminate: fileToDecontaminate,
	}
}

// Process launches the decontamination process
func (d *decontamination) Process() (string, error) {
	return "", nil
}
