package quality_controller

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/genomagic/constants"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"

	"github.com/genomagic/config_parser"
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

// TODO: Replace this function into a constants file where this can be used by other quality control processes
// getImageID attempts to find the Docker image by given term
func getImageID(client *client.Client, ctx context.Context, term string) (string, error) {
	images, err := client.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get available Docker images, err: %v", err)
	} else if len(images) == 0 {
		return "", fmt.Errorf("getImageID found no images")
	}
	found := false
	assemblerID := ""
	for _, im := range images {
		if found {
			break
		}
		for _, tag := range im.RepoTags {
			if strings.Contains(tag, term) {
				found = true
				assemblerID = im.ID
			}
		}
	}
	if !found {
		return "", fmt.Errorf("failed to find a Docker container for the given assembler")
	}
	return assemblerID, nil
}

// Process performs the error correction process
func (e *errorCorrection) Process() (string, error) {
	// correctedFile is the placeholder filename where the corrected reads are going to be stored.
	var correctedFile = path.Join(constants.BaseOut, "corrected.fastq")

	img, err := getImageID(e.dockerCLI, e.ctx, "greatfireball/canu")
	if err != nil {
		return "", fmt.Errorf("cannot get image ID for greatfireball/canu, err: %v", err)
	}

	ctConfig := &container.Config{
		Tty: true,
		Cmd: []string{
			"-correct",
			"-d", path.Join(e.config.GenoMagic.OutputPath, "canu-corr"),
			"-p", "run1",
			fmt.Sprintf("genomeSize=%d", e.config.Assemblers.Flye.GenomeSize),
			"-nanopore-raw", e.toDecontaminate,
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
		return "", fmt.Errorf("fialed to create container, err: %v", err)
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
	return correctedFile, nil
}
