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
	N50          int32         // N50 score of the assembly
	L50          int32         // L50 score of the assembly
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
	if v, err := r.getNx(50); err != nil { //compute the N50 metric of a the assembly
		return fmt.Errorf("failed to compute N50 for assembly %s, err: %v", r.result.GetAssemblyName(), err)
	} else {
		r.N50 = v
	}
	if v, err := r.getL50(); err != nil {
		return fmt.Errorf("failed to compute L50 for assembly %s, err: %v", r.result.GetAssemblyName(), err)
	} else {
		r.L50 = v
	}
	return nil
}

// GetN50 returns the computed N50 value stored on the report
func (r *report) GetN50() int32 {
	return r.N50
}

// GetL50 returns the computed L50 value stored on the report
func (r *report) GetL50() int32 {
	return r.L50
}
