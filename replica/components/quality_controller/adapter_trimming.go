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

	"github.com/genoassist/config_parser"
	"github.com/genoassist/constants"
)

// adapterTrimming is the struct representation of the adapter trimming process
type adapterTrimming struct {
	// dockerCLI is used for launching a Docker container that perform adapter trimming
	dockerCLI *client.Client
	// config is the GenoAssist global configuration
	config *config_parser.Config
	// the context of the process
	ctx context.Context
}

// NewAdapterTrimming function creates a adapterTrimming instance containing necessary config for adapter trimming job
func NewAdapterTrimming(ctx context.Context, dockerCli *client.Client, config *config_parser.Config) Controller {
	return &adapterTrimming{
		dockerCLI: dockerCli,
		config:    config,
		ctx:       ctx,
	}
}

// Process does the adapter trimming on raw input data
func (a *adapterTrimming) Process() (string, error) {

	// trimmedOutputFilename is the filename where the trimmed reads are going to be stored.
	trimmedOutputFilename := "trimmed.fastq"

	img, err := getImageID(a.dockerCLI, a.ctx, "replikation/porechop")
	if err != nil {
		return "", fmt.Errorf("cannot get image ID for replikation/porechop, err: %v", err)
	}

	ctConfig := &container.Config{
		Tty: true,
		Cmd: []string{
			"-i", constants.RawSeqIn,
			"-o", path.Join(constants.BaseOut, trimmedOutputFilename),
		},
		Image: img,
	}

	hostConfig := &container.HostConfig{
		Mounts: []mount.Mount{
			{ // Binding the input raw sequence file provided by the user
				Type:   mount.TypeBind,
				Source: a.config.GenoAssist.InputFilePath,
				Target: constants.RawSeqIn,
			},
			{ // Binding the output directory path provided by the user for saving trimmed file in.
				Type:   mount.TypeBind,
				Source: a.config.GenoAssist.OutputPath,
				Target: constants.BaseOut,
			},
		},
	}

	resp, err := a.dockerCLI.ContainerCreate(a.ctx, ctConfig, hostConfig, nil, "")
	if err != nil {
		return "", fmt.Errorf("fialed to create container, err: %v", err)
	}

	if err := a.dockerCLI.ContainerStart(a.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start adapter trimming container, err: %s", err)
	}

	statCh, errCh := a.dockerCLI.ContainerWait(a.ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", fmt.Errorf("failed to wait for container to start up, err: %v", err)
		}
	case <-statCh:
	}

	out, err := a.dockerCLI.ContainerLogs(a.ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return "", fmt.Errorf("failed to get container log, err: %v", err)
	}

	if _, err := io.Copy(os.Stdout, out); err != nil {
		return "", fmt.Errorf("failed to capture stdout from Docker assembly container, err: %v", err)
	}
	return path.Join(a.config.GenoAssist.OutputPath, trimmedOutputFilename), nil

}
