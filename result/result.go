// contains the definition of an rst
package result

import (
	"github.com/biogo/biogo/seq"
)

// rst represents the result that the parser function returns and is used by reporting component of genomagic
type rst struct {
	assemblyName string         // name of the assembly performed to obtain the contigs
	Sequences    []seq.Sequence // slice of assembly contigs
}

// New creates a new result struct and returns it
func New(an string, seqs []seq.Sequence) Result {
	return &rst{
		Sequences: seqs,
	}
}

// GetAssemblyName returns the name of the assembly the result was created for
func (r *rst) GetAssemblyName() string {
	return r.assemblyName
}

// GetSequences returns the contigs the result contains
func (r *rst) GetSequences() []seq.Sequence {
	return r.Sequences
}
