package quality_controller

import (
	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
)

// errorCorrection is a representation of the error correction process
type errorCorrection struct {
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// config is the GenoMagic global configuration
	config *config_parser.Config
	// toDecontaminate represents the path to the file to decontaminate
	toDecontaminate string
}

// NewErrorCorrection constructs and returns an errorCorrection struct, which implements the Controller interface
func NewErrorCorrection(dockerCli *client.Client, config *config_parser.Config, fileToDecontaminate string) Controller {
	return &errorCorrection{
		dockerCLI:       dockerCli,
		config:          config,
		toDecontaminate: fileToDecontaminate,
	}
}

// Process performs the error correction process
func (e *errorCorrection) Process() (string, error) {
	return "", nil
}
