// contains the definition of an rst
package result

import (
	"github.com/biogo/biogo/seq"
)

// rst represents the result that the parser function returns and is used by reporting component of genomagic
type rst struct {
	sequences []seq.Sequence
}

// New creates a new result struct and returns it
func New(seqs []seq.Sequence) Result {
	return &rst{
		sequences: seqs,
	}
}
