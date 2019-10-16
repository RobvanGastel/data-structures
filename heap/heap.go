package heap

import "math"

// Heap object to store size and array
// is being modified inplace
type Heap struct {
	A        []int
	HeapSize int
}

// Parent returns parent node of index i
func (H *Heap) Parent(i int) *int {
	p := int(math.Floor(float64(i-1) / 2.0))
	return &p
}

// Left returns left node of index i
func (H *Heap) Left(i *int) *int {
	l := 2 * (*i)
	l = l + 1
	return &l
}

// Right returns right node of index i
func (H *Heap) Right(i *int) *int {
	r := 2 * (*i)
	r = r + 2
	return &r
}
