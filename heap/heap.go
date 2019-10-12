package heap

import (
	"math"
	"errors"
)

// TODO: Write tests for heaps

func Parent(i int) int {
	return int(math.Floor(float64(i-1)/2.0))
}

func Left(i *int) *int {
	l := 2*(*i)
	return &l
}

func Right(i *int) *int {
	r := 2*(*i+1)
	return &r
}

func HeapMaximum(A []int) *int {
	return &A[1]
}

func MaxHeapify(A []int, i *int) {
	l := Left(i)
	r := Right(i)
	largest := 0

	if *l <= len(A) && A[*l] > A[*i] {
		largest = *l
	} else {
		largest = *i
	}

	if *r <= len(A)-1 && A[*r] > A[largest] {
		largest = *r
	}

	if largest != *i {
		temp := A[largest]
		A[largest] = A[*i]
		A[*i] = temp
		MaxHeapify(A, &largest)
	}
}

func BuildMaxHeap(A []int) {
	A_floor := int(math.Floor(float64(len(A)))/2.0)-1
	for i := A_floor; i > 1; i-- {
		MaxHeapify(A, &i)
	}
}

func HeapExtractMax(A []int) (int, error) {
	if len(A) < 1 {
		return 0, errors.New("heap underflow")
	}

	max := A[1]
	i := 1
	MaxHeapify(A, &i)
	return max, nil
}

func HeapIncreaseKey(A []int, i *int, key *int) error {
	j := *i
	if *key < A[*i] {
		return errors.New("new key is smaller than current key")
	}

	A[*i] = *key
	for j > 1 && A[Parent(j)] < A[j] {
		temp := A[Parent(j)]
		A[Parent(*i)] = A[j]
		A[j] = temp

		j = Parent(j)
	}
	return nil
}

func MaxHeapInsert(A []int, key *int) error {
	A = append(A, -2147483648)
	n := len(A)
	return HeapIncreaseKey(A, &n, key)
}