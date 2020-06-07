package quality_controller

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
	"github.com/genomagic/constants"
)

// errorCorrection is a representation of the error correction process
type errorCorrection struct {
	// context of the process
	ctx context.Context
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// config is the GenoMagic global configuration
	config *config_parser.Config
	// toDecontaminate represents the path to the file to decontaminate
	toDecontaminate string
}

// NewErrorCorrection constructs and returns an errorCorrection struct, which implements the Controller interface
func NewErrorCorrection(ctx context.Context, dockerCli *client.Client, config *config_parser.Config, fileToCorrect string) Controller {
	return &errorCorrection{
		ctx:             ctx,
		dockerCLI:       dockerCli,
		config:          config,
		toDecontaminate: fileToCorrect,
	}
}

// Process performs the error correction process
func (e *errorCorrection) Process() (string, error) {
	// correctedFile is the placeholder filename where the corrected reads are going to be stored.
	correctedOutuptFile := "run1.correctedReads.fasta.gz"

	img, err := getImageID(e.dockerCLI, e.ctx, "greatfireball/canu")
	if err != nil {
		return "", fmt.Errorf("cannot get image ID for greatfireball/canu, err: %v", err)
	}

	ctConfig := &container.Config{
		Tty: true,
		Cmd: []string{
			"-correct",
			"-d", path.Join(constants.BaseOut, "canu-corr"),
			"-p", "run1",
			fmt.Sprintf("genomeSize=%s", e.config.Assemblers.Flye.GenomeSize),
			"-nanopore-raw", path.Join(constants.BaseOut, "trimmed.fastq"),
		},
		Image: img,
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{ // Binding the output directory path provided by the user for saving error-corrected file in.
				Type:   mount.TypeBind,
				Source: e.config.GenoMagic.OutputPath,
				Target: constants.BaseOut,
			},
		},
	}

	resp, err := e.dockerCLI.ContainerCreate(e.ctx, ctConfig, hostConfig, nil, "")
	if err != nil {
		return "", fmt.Errorf("failed to create error correction container, err: %s", err)
	}

	if err := e.dockerCLI.ContainerStart(e.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container, err: %v", err)
	}

	statCh, errCh := e.dockerCLI.ContainerWait(e.ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", fmt.Errorf("failed to wait for container to start up, err: %v", err)
		}
	case <-statCh:
	}

	out, err := e.dockerCLI.ContainerLogs(e.ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", fmt.Errorf("failed to get container log, err: %v", err)
	}

	if _, err := io.Copy(os.Stdout, out); err != nil {
		return "", fmt.Errorf("failed to capture stdout from Docker assembly container, err: %v", err)
	}
	return path.Join(e.config.GenoMagic.OutputPath, "canu-corr", correctedOutuptFile), nil
}
