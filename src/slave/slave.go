// slave is responsible for launching and coordinating processes such as
// assembly and parsing for the master
package slave

import (
	"fmt"

	"github.com/genomagic/src/slave/components/parser"
	"github.com/genomagic/src/slave/components/assembler"
	"github.com/genomagic/src/slave/components"
)

const (
	Assembly = "assembly"
	Parse    = "parse"
)

// the type of component that runs on a specific set of assembly files
type ComponentType string

// a mapping between work types and the associated components
var WorkType = map[ComponentType]func(f string) components.Component{
	Assembly: assembler.NewAssembler,
	Parse:    parser.NewParser,
}

// slv defines the structure of a slave
type slv struct {
	// name/description of the work performed by the slave
	description string
	// the file the slave is supposed to perform work on
	fileName string
	// the type of work that has to be performed by the slave
	workType ComponentType
}

// NewSlave creates and returns a new instance of a slave
func NewSlave(dsc, fnm, wtp string) *slv {
	return &slv{
		description: dsc,
		fileName:    fnm,
		workType:    wtp,
	}
}

// Process performs the work that's dictated by the master
func (s *slv) Process() error {
	worker := WorkType[s.workType](s.fileName)
	if err := worker.Process(); err != nil {
		return fmt.Errorf("slave process failed, err: %v", err)
	}
	return nil
}
