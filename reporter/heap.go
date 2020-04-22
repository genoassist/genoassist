package reporter

// An IntHeap is a max-heap of ints
type IntHeap []int

// Len returns the number of items in the heap
func (h IntHeap) Len() int {
	return len(h)
}

// Less tells whether elements at position j is less than the one at position i
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap swaps two elements in the heap
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop returns the element at the top of the heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
