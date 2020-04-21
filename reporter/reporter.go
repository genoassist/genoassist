// the reporter package computes multiple statistics that quantify the quality of a collection of contigs
package reporter

import (
	"fmt"

	"github.com/genomagic/result"
)

// report represents the struct that holds the stats that characterize an assembly
type report struct {
	assemblyName string        // name of the assembly the report represents
	result       result.Result // a collection of assembly results, which includes the assembly contigs
	n50          float32       // n50 score of the assembly
	l50          float32       // l50 score of the assembly
}

// New returns a new instance of a report
func New(an string, rs result.Result) Reporter {
	return &report{
		assemblyName: an,
		result:       rs,
	}
}

// Process constructs the report for the given assembler results
func (r *report) Process() error {
	if v, err := r.getL50(); err != nil {
		return fmt.Errorf("failed to compute L50 for assembly: %s", r.result.GetAssemblyName())
	} else {
		r.l50 = v
	}
	if v, err := r.getN50(); err != nil {
		return fmt.Errorf("failed to compute N50 for assembly: %s", r.result.GetAssemblyName())
	} else {
		r.n50 = v
	}
	return nil
}
