// slave is responsible for launching and coordinating processes such as assembly and parsing
// for the master
package slave

import (
	"github.com/genomagic/src/slave/components/parser"
	"github.com/genomagic/src/slave/components/assembler"
)

const (
	Assembly = "assembly"
	Parse    = "parse"
)

// a mapping between work types and the associated components
var WorkType = map[string]func(f string) interface{}{
	Assembly: assembler.NewAssembler,
	Parse:    parser.NewParser,
}

// slv defines the structure of a slave
type slv struct {
	description string // name/description of the work performed by the slave
	fileName    string // the file the slave is supposed to perform work on
	workType    string // the type of work that has to be performed by the slave
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
	_ := WorkType[s.workType](s.fileName)
	return nil
}
