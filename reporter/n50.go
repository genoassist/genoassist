package reporter

import (
	"fmt"
)

// getN50 computes and returns the N50 score of the report contigs
// https://en.wikipedia.org/wiki/N50,_L50,_and_related_statistics#N50
func (r *report) getN50() (int32, error) {
	assemblyLen := 0
	contigLens := make([]int, len(r.result.GetSequences()))
	for i, seq := range r.result.GetSequences() {
		contigLens[i] = seq.Len()
		assemblyLen += seq.Len()
	}
	halfAssemblyLen := assemblyLen / 2
	sumToHalf := 0
	for i, cl := range contigLens {
		if sumToHalf == halfAssemblyLen || sumToHalf > halfAssemblyLen {
			return contigLens[i+1], nil
		}
		sumToHalf += cl
	}
	return 0, fmt.Errorf("failed to compute N50 for an unknown reason")
}
