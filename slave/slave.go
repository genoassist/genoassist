// slave is responsible for launching and coordinating processes such as
// assembly and parsing for the master
package slave

import (
	"fmt"

	"github.com/genomagic/constants"
	"github.com/genomagic/result"
	"github.com/genomagic/slave/components/assembler"
	"github.com/genomagic/slave/components/parser"
)

const (
	Assembly = "assembly"
	Parse    = "parse"
)

// the type of component that runs on a specific set of assembly files
type ComponentWorkType string

// slv defines the structure of a slave
type slv struct {
	description string            // name/description of the work performed by the slave
	filePath    string            // the file the slave is supposed to perform work on
	outPath     string            // a path to the location where results will be stored
	numThreads  int               // number of threads to use for slave processes
	workType    ComponentWorkType // the type of work that has to be performed by the slave
}

// New creates and returns a new instance of a slave
func New(dsc, fnm, out string, thr int, wtp ComponentWorkType) Slave {
	return &slv{
		description: dsc,
		filePath:    fnm,
		outPath:     out,
		numThreads:  thr,
		workType:    wtp,
	}
}

// Process performs the work that's dictated by the master
func (s *slv) Process() ([]result.Result, error) {
	if s.workType == Assembly {
		for k := range constants.AvailableAssemblers {
			assemblerWorker, err := assembler.New(s.filePath, s.outPath, k, s.numThreads)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize assembler worker, err %v", err)
			}

			_, err = assemblerWorker.Process()
			if err != nil {
				return nil, fmt.Errorf("assembler slave process failed, err: %v", err)
			}
		}
		return nil, nil
	} else {
		var results []result.Result
		for k := range constants.AvailableAssemblers {
			parserWorker, err := parser.New(s.filePath, s.outPath, k)
			if err != nil {
				return nil, fmt.Errorf("failed to initialize parser worker, err: %v", err)
			}
			res, err := parserWorker.Process()
			if err != nil {
				return nil, fmt.Errorf("parser slave process failed, err: %v", err)
			}
			results = append(results, res)
		}
		return results, nil
	}
}
