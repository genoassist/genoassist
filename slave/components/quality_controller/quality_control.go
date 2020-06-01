package quality_controller

import (
	"context"
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
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", fmt.Errorf("failed to initialize Docker client, err: %s", err)
	}

	trimmer := NewAdapterTrimming(ctx, cli, q.config)
	trimmedFile, err := trimmer.Process()
	if err != nil {
		return "", fmt.Errorf("failed to perform raw sequence adapter trimming, err: %s", err)
	}

	//decontaminator := NewDecontamination(ctx, cli, q.config, trimmedFile)
	//decontaminatedFile, err := decontaminator.Process()
	//if err != nil {
	//	return "", fmt.Errorf("failed to perform trimmed file decontamination, err: %s", err)
	//}
	//
	//corrector := NewErrorCorrection(ctx, cli, q.config, decontaminatedFile)
	//correctedFile, err := corrector.Process()
	//if err != nil {
	//	return "", fmt.Errorf("failed to perform error correction on the decontaminated file, err: %s", err)
	//}
	return trimmedFile, nil
}
