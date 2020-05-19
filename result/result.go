// contains the definition of an Result
package result

import "github.com/biogo/biogo/seq"

// Result represents the result that the parser function returns and is used by reporting component of genomagic
type Result struct {
	// name of the assembly performed to obtain the contigs
	AssemblyName string
	// slice of assembly contigs
	Sequences []seq.Sequence
}

// New creates a new result struct and returns it
func New(assemblyName string, sequences []seq.Sequence) *Result {
	return &Result{
		AssemblyName: assemblyName,
		Sequences:    sequences,
	}
}
