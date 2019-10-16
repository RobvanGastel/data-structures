package heap

import (
	"errors"
	"math"
)

// Uses main class heap.go, for generic struct Heap

// BuildMaxHeap in place orders the array as a heap
func BuildMaxHeap(A *[]int) Heap {
	size := len(*A) - 1
	H := Heap{A: *A, HeapSize: size}

	AFloor := int(math.Floor(float64(len(*A)))/2.0) - 1
	for i := AFloor; i >= 0; i-- {
		H.MaxHeapify(&i)
	}
	return H
}

// Maximum returns maximum node in Heap
func (H *Heap) Maximum() *int {
	return &H.A[0]
}

// MaxHeapify restores Heap order from index i
func (H *Heap) MaxHeapify(i *int) {
	l := *H.Left(i)
	r := *H.Right(i)
	largest := 0

	if l < H.HeapSize && H.A[l] > H.A[*i] {
		largest = l
	} else {
		largest = *i
	}

	if r < H.HeapSize && H.A[r] > H.A[largest] {
		largest = r
	}

	if largest != *i {
		// Swap indexes
		// a, b = b, a
		H.A[largest], H.A[*i] = H.A[*i], H.A[largest]
		H.MaxHeapify(&largest)
	}
}

// ExtractMaximum extracts the maximum node of a heap
func (H *Heap) ExtractMaximum() (int, error) {
	if H.HeapSize < 1 {
		return 0, errors.New("heap underflow")
	}
	max := H.A[0]
	i := 0

	// Set first item to last item
	H.A[i] = H.A[H.HeapSize]
	H.HeapSize = H.HeapSize - 1

	// Remove last index
	H.A = H.A[:H.HeapSize-1]
	H.MaxHeapify(&i)
	return max, nil
}

// IncreaseMaxKey restores order after adding a key
func (H *Heap) IncreaseMaxKey(key *int) error {
	j := H.HeapSize
	if *key < H.A[j] {
		return errors.New("new key is smaller than current key")
	}

	H.A[j] = *key
	for j > 0 && H.A[*H.Parent(j)] < H.A[j] {
		// Swap indexes
		// a, b = b, a
		H.A[*H.Parent(j)], H.A[j] = H.A[j], H.A[*H.Parent(j)]
		j = *H.Parent(j)
	}
	return nil
}

// MaxHeapInsert inserts a key in the heap and maintains order
func (H *Heap) MaxHeapInsert(key *int) error {
	H.HeapSize = H.HeapSize + 1
	H.A = append((H.A), -2147483648)
	return H.IncreaseMaxKey(key)
}
