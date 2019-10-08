package heap

import (
	"math"
)

func Parent(i int) int {
	return int(math.Floor(float64(i-1)/2.0))
}

func Left(i int) int {
	return 2*i+1
}

func Right(i int) int {
	return 2*i+2
}

func HeapMaximum(A []int) int {
	return A[1]
}

func MaxHeapify(A []int, i int) {
	l := Left(i)
	r := Right(i)
	largest := 0

	if r <= len(A) && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}

	if r<= len(A) && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		temp := A[largest]
		A[largest] = A[i]
		A[i] = temp
		MaxHeapify(A, i)
	}
}

func BuildMaxHeap(A []int) {
	for i := int(math.Floor(float64(len(A)))/2.0); i > 1; i-- {
		MaxHeapify(A, i)
	}
}