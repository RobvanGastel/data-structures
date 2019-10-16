package heap

import (
	test "data-structures/test"
	"math/rand"
	"testing"
)

// Valid date if it builds a valid Min Heap
func TestCheckMinHeap(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	h := BuildMinHeap(&A) // Build inplace

	validHeap := true
	prevMin, _ := h.ExtractMinimum()
	n := len(A)
	for j := 0; j < n; j++ {
		min, _ := h.ExtractMinimum()

		if prevMin > min {
			validHeap = false
			break
		}
	}

	test.AssertEqual(t, true, validHeap)
}
