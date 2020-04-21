// the reporter package computes multiple statistics that quantify the quality of a collection of contigs
package reporter

import "github.com/genomagic/result"

// report represents the struct that holds the stats that characterize an assembly
type report struct {
	assemblyName string         // name of the assembly the report represents
	result       *result.Result // a collection of assembly results, which includes the assembly contigs
}

// New returns a new instance of a report
func New(an string, rs *result.Result) *report {
	return &report{
		assemblyName: an,
		result:       rs,
	}
}
