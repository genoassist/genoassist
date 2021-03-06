package reporter

import (
	"fmt"
)

// getN50 computes and returns the N50 score of the Report contigs
// https://en.wikipedia.org/wiki/N50,_L50,_and_related_statistics#N50
func (r *Report) getN50() (int32, error) {

	// immediately return the contig len if it is the only contig in the assembly
	if len(r.result.Sequences) == 1 {
		return int32(r.result.Sequences[0].Len()), nil
	}

	assemblyLen := 0
	contigLens := make([]int, len(r.result.Sequences))

	for i, seq := range r.result.Sequences {
		contigLens[i] = seq.Len()
		assemblyLen += seq.Len()
	}
	halfAssemblyLen := assemblyLen / 2
	if halfAssemblyLen == 0 {
		return 0, fmt.Errorf("failed to compute N50 due to potentially missing contigs")
	}
	sumToHalf := 0
	for i, cl := range contigLens {
		if sumToHalf >= halfAssemblyLen {
			return int32(contigLens[i]), nil
		}
		sumToHalf += cl

	}
	return 0, fmt.Errorf("failed to compute N50 for an unknown reason")
}
