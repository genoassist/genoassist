// contains the definition and work associated with assemblers
package assembler

import (
	"fmt"

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
func NewAssembler(fnm string) (components.Component, error) {
	a := &asmbler{
		fileName: fnm,
	}
	if cli, err := client.NewEnvClient(); err != nil {
		return nil, fmt.Errorf("failed to initialize Docker client, err: %v", err)
	} else {
		a.dClient = cli
	}
	containers, err := 1, nil
	return a, nil
}

// Process performs the work of the assembler
func (a *asmbler) Process() error {
	return nil
}
