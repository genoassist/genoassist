package quality_controller

import (
	"fmt"

	"github.com/docker/docker/client"
	"github.com/genomagic/config_parser"
)

// qualityController is the struct representation of the quality control process
type qualityController struct {
	// config represents the GenoMagic configuration
	config *config_parser.Config
}

// NewQualityController constructs a new qualityController instances that implements the Controller interface
func NewQualityController(c *config_parser.Config) Controller {
	return &qualityController{
		config: c,
	}
}

// Process launches the trimming, decontamination, and error correction process
func (q *qualityController) Process() (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("failed to initialize Docker client, err: %s", err)
	}

	var correctedFile string
	trimmer := NewAdapterTrimming(cli, q.config)
	correctedFile, err = trimmer.Process()

}
