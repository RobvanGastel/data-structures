package heap

import "math"

// Heap is being modified inplace,
// TODO: Improve heap to be more generic

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
