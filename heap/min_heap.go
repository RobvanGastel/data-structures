package heap

import (
	"errors"
	"math"
)

// Minimum returns minimum node in Heap
func Minimum(A *[]int) *int {
	return &(*A)[0]
}

// MinHeapify restores Heap order from index i
func MinHeapify(A *[]int, i *int) {
	l := Left(i)
	r := Right(i)
	smallest := 0

	if *l < len(*A) && (*A)[*l] < (*A)[*i] {
		smallest = *l
	} else {
		smallest = *i
	}

	if *r < len(*A) && (*A)[*r] < (*A)[smallest] {
		smallest = *r
	}

	if smallest != *i {
		// Swap indexes
		// a, b = b, a
		(*A)[smallest], (*A)[*i] = (*A)[*i], (*A)[smallest]
		MinHeapify(A, &smallest)
	}
}

// BuildMinHeap in place orders the array as a heap
func BuildMinHeap(A *[]int) {
	AFloor := int(math.Floor(float64(len(*A)))/2.0) - 1
	for i := AFloor; i >= 0; i-- {
		MinHeapify(A, &i)
	}
}

// ExtractMinimum extracts the minimum node of a heap
func ExtractMinimum(A *[]int) (int, error) {
	if len(*A) < 1 {
		return 0, errors.New("heap underflow")
	}
	min := (*A)[0]
	i := 0

	// Set first item to last item
	(*A)[i] = (*A)[len(*A)-1]

	// Remove last index
	(*A) = (*A)[:len(*A)-1]
	MinHeapify(A, &i)
	return min, nil
}

// IncreaseKey restores order after adding a key
func IncreaseMinKey(A *[]int, i *int, key *int) error {
	j := *i
	if *key > (*A)[j-1] {
		return errors.New("new key is bigger than current key")
	}

	(*A)[j] = *key
	for j > 0 && (*A)[*Parent(j)] > (*A)[j] {
		// Swap indexes
		// a, b = b, a
		(*A)[*Parent(j)], (*A)[j] = (*A)[j], (*A)[*Parent(j)]
		j = *Parent(j)
	}
	return nil
}

// MinHeapInsert inserts a key in the heap and maintains order
func MinHeapInsert(A *[]int, key *int) error {
	*A = append((*A), 2147483648)
	n := len(*A) - 1 // Index starts at one
	return IncreaseMinKey(A, &n, key)
}
