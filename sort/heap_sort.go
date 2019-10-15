package sort

import (
	. "data-structures/heap"
)

// TODO: Add asc or desc,
// TODO: Define a heap-size nessecary
// to keep order in rest of array
func HeapSort(A *[]int) {
	BuildMaxHeap(A)

	for i := len(*A) - 1; i > 1; i-- {
		(*A)[i], (*A)[0] = (*A)[0], (*A)[i]

		// heap-size[ A] ← heap-size[ A] − 1
		j := 0
		MaxHeapify(A, &j)
	}
}
