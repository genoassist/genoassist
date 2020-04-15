// slave is responsible for launching and coordinating processes such as
// assembly and parsing for the master
package slave

import (
	"fmt"

	"github.com/genomagic/constants"
	"github.com/genomagic/slave/components"
	"github.com/genomagic/slave/components/assembler"
	"github.com/genomagic/slave/components/parser"
)

const (
	Assembly = "assembly"
	Parse    = "parse"
)

// the type of component that runs on a specific set of assembly files
type ComponentWorkType string

// a mapping between work types and the associated components
var WorkType = map[ComponentWorkType]func(fp, op, pr string) (components.Component, error){
	Assembly: assembler.NewAssembler,
	Parse:    parser.NewParser,
}

// slv defines the structure of a slave
type slv struct {
	description string            // name/description of the work performed by the slave
	filePath    string            // the file the slave is supposed to perform work on
	outPath     string            // a path to the location where results will be stored
	workType    ComponentWorkType // the type of work that has to be performed by the slave
}

// NewSlave creates and returns a new instance of a slave
func NewSlave(dsc, fnm, out string, wtp ComponentWorkType) *slv {
	return &slv{
		description: dsc,
		filePath:    fnm,
		outPath:     out,
		workType:    wtp,
	}
}

// Process performs the work that's dictated by the master
func (s *slv) Process() error {
	worker, err := WorkType[s.workType](s.filePath, s.outPath, constants.MegaHit)
	if err != nil {
		return fmt.Errorf("failed to initialize worker, err: %v", err)
	}
	if err := worker.Process(); err != nil {
		return fmt.Errorf("slave process failed, err: %v", err)
	}
	return nil
}
