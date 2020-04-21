// contains the definition of an rst
package result

import (
	"github.com/biogo/biogo/seq"
)

// rst represents the result that the parser function returns and is used by reporting component of genomagic
type rst struct {
	Sequences []seq.Sequence // slice of assembly contigs
}

// New creates a new result struct and returns it
func New(seqs []seq.Sequence) Result {
	return &rst{
		Sequences: seqs,
	}
}
