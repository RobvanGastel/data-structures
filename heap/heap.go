package heap

import (
	"errors"
	"math"
)

// Heap is being modified inplace

// Parent returns parent node of index i
func Parent(i int) *int {
	p := int(math.Floor(float64(i-1) / 2.0))
	return &p
}

// Left returns left node of index i
func Left(i *int) *int {
	l := 2 * (*i)
	l = l + 1
	return &l
}

// Right returns right node of index i
func Right(i *int) *int {
	r := 2 * (*i)
	r = r + 2
	return &r
}

// Maximum returns maximum node in Heap
func Maximum(A *[]int) *int {
	return &(*A)[0]
}

// MaxHeapify restores Heap order from index i
func MaxHeapify(A *[]int, i *int) {
	l := Left(i)
	r := Right(i)
	largest := 0

	if *l < len(*A) && (*A)[*l] > (*A)[*i] {
		largest = *l
	} else {
		largest = *i
	}

	if *r < len(*A) && (*A)[*r] > (*A)[largest] {
		largest = *r
	}

	if largest != *i {
		// Swap indexes
		// a, b = b, a
		(*A)[largest], (*A)[*i] = (*A)[*i], (*A)[largest]
		MaxHeapify(A, &largest)
	}
}

// BuildMaxHeap in place orders the array as a heap
func BuildMaxHeap(A *[]int) {
	AFloor := int(math.Floor(float64(len(*A)))/2.0) - 1
	for i := AFloor; i >= 0; i-- {
		MaxHeapify(A, &i)
	}
}

// ExtractMaximum extracts the maximum node of a heap
func ExtractMaximum(A *[]int) (int, error) {
	if len(*A) < 1 {
		return 0, errors.New("heap underflow")
	}
	max := (*A)[0]
	i := 0

	// Set first item to last item
	(*A)[i] = (*A)[len(*A)-1]

	// Remove last index
	(*A) = (*A)[:len(*A)-1]
	MaxHeapify(A, &i)
	return max, nil
}

// IncreaseKey restores order after adding a key
func IncreaseKey(A *[]int, i *int, key *int) error {
	j := *i
	if *key < (*A)[j-1] {
		return errors.New("new key is smaller than current key")
	}

	(*A)[j] = *key
	for j > 0 && (*A)[*Parent(j)] < (*A)[j] {
		// Swap indexes
		// a, b = b, a
		(*A)[*Parent(j)], (*A)[j] = (*A)[j], (*A)[*Parent(j)]
		j = *Parent(j)
	}
	return nil
}

// MaxHeapInsert inserts a key in the heap and maintains order
func MaxHeapInsert(A *[]int, key *int) error {
	*A = append((*A), -2147483648)
	n := len(*A) - 1 // Index starts at one
	return IncreaseKey(A, &n, key)
}
