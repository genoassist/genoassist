// contains the definition and work associated with assemblers
package assembler

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"github.com/genomagic/src/slave/components"
)

// structure of the assembler
type asmbler struct {
	// name of the file the assembler will operate on
	fileName string
	// the Docker client the assembler will use to spin up containers
	dClient *client.Client
}

// NewAssembler returns a new assembler for the specified file
// TODO: this needs to take in an assembler name
func NewAssembler(fnm string) (components.Component, error) {
	a := &asmbler{
		fileName: fnm,
	}
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Docker client, err: %v", err)
	} else {
		a.dClient = cli
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get available Docker containers, err: %v", err)
	}
	// need to find whether the megahit container is available
	// TODO: this needs to find the given assembler container and call it with appropriate params
	for _, im := range images {
		fmt.Printf("%s %v\n", im.ID, im.RepoTags)
	}
	return a, nil
}

// Process performs the work of the assembler
func (a *asmbler) Process() error {
	return nil
}
