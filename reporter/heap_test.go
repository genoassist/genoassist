package reporter

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Pop  = "Pop"
	Push = "Push"
)

func TestIntHeap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		heap               *IntHeap
		expectedLen        int
		toPush             []int
		ops                []string
		expectedLenPostOps int
		expectedSeqPostOps []int
	}{
		{
			name:               "test_heap_empty_post_pop",
			heap:               &IntHeap{1, 2, 3},
			expectedLen:        3,
			toPush:             []int{},
			ops:                []string{Pop, Pop, Pop},
			expectedLenPostOps: 0,
			expectedSeqPostOps: []int{},
		},
		{
			name:               "test_sequence_matches_expected_sequence_post_pop",
			heap:               &IntHeap{1, 2, 3, 4, 5},
			expectedLen:        5,
			toPush:             []int{},
			ops:                []string{Pop, Pop, Pop},
			expectedLenPostOps: 2,
			expectedSeqPostOps: []int{2, 1},
		},
		{
			name:               "test_sequence_matches_expected_sequence_post_push_pop",
			heap:               &IntHeap{1, 2, 3},
			expectedLen:        3,
			toPush:             []int{4, 5},
			ops:                []string{Push, Pop, Push, Pop},
			expectedLenPostOps: 3,
			expectedSeqPostOps: []int{3, 2, 1},
		},
		{
			name:               "test_sequence_matches_expected_sequence_post_pop_push",
			heap:               &IntHeap{1, 2, 3},
			expectedLen:        3,
			toPush:             []int{4, 5, 6},
			ops:                []string{Pop, Pop, Pop, Push, Push, Push},
			expectedLenPostOps: 3,
			expectedSeqPostOps: []int{6, 5, 4},
		},
	}
	for _, tt := range tests {
		heap.Init(tt.heap)
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.heap.Len(), tt.expectedLen)
			for _, op := range tt.ops {
				switch op {
				case Pop:
					heap.Pop(tt.heap)
				case Push:
					v := tt.toPush[0]
					tt.toPush = tt.toPush[1:]
					heap.Push(tt.heap, v)
				}
			}
			postSeq := []int{}
			for tt.heap.Len() > 0 {
				postSeq = append(postSeq, heap.Pop(tt.heap).(int))
			}
			assert.EqualValues(t, tt.expectedSeqPostOps, postSeq)
			assert.Equal(t, tt.expectedLenPostOps, len(postSeq))
		})
	}
}
