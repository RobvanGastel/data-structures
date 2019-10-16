package heap

import (
	"errors"
	"math"
)

// Uses main class heap.go, for generic struct Heap

// BuildMinHeap in place orders the array as a heap
func BuildMinHeap(A *[]int) Heap {
	size := len(*A) - 1
	H := Heap{A: *A, HeapSize: size}

	AFloor := int(math.Floor(float64(len(*A)))/2.0) - 1
	for i := AFloor; i >= 0; i-- {
		H.MinHeapify(&i)
	}
	return H
}

// Minimum returns minimum node in Heap
func (H *Heap) Minimum() *int {
	return &H.A[0]
}

// MinHeapify restores Heap order from index i
func (H *Heap) MinHeapify(i *int) {
	l := *H.Left(i)
	r := *H.Right(i)
	smallest := 0

	if l < H.HeapSize && H.A[l] < H.A[*i] {
		smallest = l
	} else {
		smallest = *i
	}

	if r < H.HeapSize && H.A[r] < H.A[smallest] {
		smallest = r
	}

	if smallest != *i {
		// Swap indexes
		// a, b = b, a
		H.A[smallest], H.A[*i] = H.A[*i], H.A[smallest]
		H.MinHeapify(&smallest)
	}
}

// ExtractMinimum extracts the minimum node of a heap
func (H *Heap) ExtractMinimum() (int, error) {
	if H.HeapSize < 1 {
		return 0, errors.New("heap underflow")
	}
	min := H.A[0]
	i := 0

	// Set first item to last item
	H.A[i] = H.A[H.HeapSize]
	H.HeapSize = H.HeapSize - 1

	// Remove last index
	H.A = H.A[:H.HeapSize-1]
	H.MinHeapify(&i)
	return min, nil
}

// IncreaseMinKey restores order after adding a key
func (H *Heap) IncreaseMinKey(key *int) error {
	j := H.HeapSize
	if *key > H.A[j] {
		return errors.New("new key is smaller than current key")
	}

	H.A[j] = *key
	for j > 0 && H.A[*H.Parent(j)] > H.A[j] {
		// Swap indexes
		// a, b = b, a
		H.A[*H.Parent(j)], H.A[j] = H.A[j], H.A[*H.Parent(j)]
		j = *H.Parent(j)
	}
	return nil
}

// MinHeapInsert inserts a key in the heap and maintains order
func (H *Heap) MinHeapInsert(A *[]int, key *int) error {
	H.HeapSize = H.HeapSize + 1
	H.A = append((H.A), 2147483648)
	return H.IncreaseMinKey(key)
}
