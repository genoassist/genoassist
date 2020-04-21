package reporter

import (
	"container/heap"
	"fmt"
)

// getL50 computes and returns the L50 score of the report contigs
// https://en.wikipedia.org/wiki/N50,_L50,_and_related_statistics#L50
func (r *report) getL50() (int32, error) {
	assemblyLen := 0
	ch := &IntHeap{}
	for _, seq := range r.result.GetSequences() {
		ch.Push(seq.Len())
		assemblyLen += seq.Len()
	}
	heap.Init(ch)

	halfAssemblyLen := assemblyLen / 2
	if halfAssemblyLen == 0 {
		return 0, fmt.Errorf("failed to compute N50 due to potentially missing contigs")
	}

	L50 := 0
	L50Len := 0
	for ch.Len() > 0 && L50Len <= halfAssemblyLen {
		el := ch.Pop().(int)
		if el > halfAssemblyLen {
			continue
		} else {
			L50++
			L50Len += el
		}
	}
	return int32(L50), nil
}
