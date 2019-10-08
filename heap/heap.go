package heap

import (
	"math"
	"errors"
)

// TODO: Write tests for heaps

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

func HeapExtractMax(A []int) (int, error) {
	if len(A) < 1 {
		return 0, errors.New("heap underflow")
	}

	max := A[1]
	MaxHeapify(A, 1)
	return max, nil
}

func HeapIncreaseKey(A []int, i int, key int) error {
	if key < A[i] {
		return errors.New("new key is smaller than current key")
	}

	A[i] = key
	for i > 1 && A[Parent(i)] < A[i] {
		temp := A[Parent(i)]
		A[Parent(i)] = A[i]
		A[i] = temp

		i = Parent(i)
	}
	return nil
}

func MaxHeapInsert(A []int, key int) {
	A = append(A, -2147483648)
	HeapIncreaseKey(A, len(A), key)
}