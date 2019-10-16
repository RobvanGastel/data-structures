package sort

import (
	heap "data-structures/heap"
)

// HeapSort sorts an array asc or desc by using a heap
func HeapSort(A *[]int, order string) {

	if order == "asc" {
		h := heap.BuildMinHeap(A)

		for i := len(*A) - 1; i > 0; i-- {
			(*A)[i], (*A)[0] = (*A)[0], (*A)[i]

			h.HeapSize--
			j := 0
			h.MinHeapify(&j)
		}
		return
	}
	// Else default use "desc"
	h := heap.BuildMaxHeap(A)

	for i := len(*A) - 1; i > 0; i-- {
		(*A)[i], (*A)[0] = (*A)[0], (*A)[i]

		h.HeapSize--
		j := 0
		h.MaxHeapify(&j)
	}
}
