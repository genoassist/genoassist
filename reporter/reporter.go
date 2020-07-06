// the reporter package computes multiple statistics that quantify the quality of a collection of contigs
package reporter

import (
	"fmt"

	"github.com/genomagic/result"
)

// Report represents the struct that holds the stats that characterize an assembly
type Report struct {
	// AssemblyName is the name of the assembly the Report represents
	AssemblyName string
	// result represents a collection of assembly results, which includes the assembly contigs
	result *result.Result
	// Processed is an indicator that represents whether the reporter process has been executed
	Processed bool
	// N50 score of the assembly
	N50 int32
	// L50 score of the assembly
	L50 int32
}

// NewReporter returns a new instance of a Report that implements the Reporter interface
func NewReporter(assemblyName string, result *result.Result) Reporter {
	return &Report{
		AssemblyName: assemblyName,
		result:       result,
		Processed:    false,
	}
}

// Process constructs the Report for the given assembler results
func (r *Report) Process() error {
	if v, err := r.getN50(); err != nil {
		return fmt.Errorf("failed to compute N50 for assembly %s, err: %v", r.result.AssemblyName, err)
	} else {
		r.N50 = v
	}
	if v, err := r.getL50(); err != nil {
		return fmt.Errorf("failed to compute L50 for assembly %s, err: %v", r.result.AssemblyName, err)
	} else {
		r.L50 = v
	}
	r.Processed = true
	return nil
}

// GetN50 returns the computed N50 value stored on the Report. An error is returned if the reporter process has not been executed
func (r *Report) GetN50() (int32, error) {
	if !r.Processed {
		return 0, fmt.Errorf("the reporter process has not been executed")
	}
	return r.N50, nil
}

// GetL50 returns the computed L50 value stored on the Report. An error is returned if the reporter process has not been executed
func (r *Report) GetL50() (int32, error) {
	if !r.Processed {
		return 0, fmt.Errorf("the reporter process has not been executed")
	}
	return r.L50, nil
}

// GetAssemblyName returns the reporter assembly name
func (r *Report) GetAssemblyName() string {
	return r.AssemblyName
}
