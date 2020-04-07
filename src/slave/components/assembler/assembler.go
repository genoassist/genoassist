// contains the definition and work associated with assemblers
package assembler

import (
	"github.com/genomagic/src/slave/components"
)

// structure of the assembler
type asmbler struct {
	// name of the file the assembler will operate on
	fileName string
}

// NewAssembler returns a new assembler for the specified file
func NewAssembler(fnm string) components.Component {
	return &asmbler{
		fileName: fnm,
	}
}

// Process performs the work of the assembler
func (a *asmbler) Process() error {
	return nil
}
