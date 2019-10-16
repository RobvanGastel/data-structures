package heap

import (
	test "data-structures/test"
	"math/rand"
	"testing"
)

func BenchmarkInsert(*testing.B) {
	// Fill max heap
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		A = append(A, rand.Intn(30))
	}

	_ = BuildMaxHeap(&A)
}

// // Valid date if it builds a valid Max Heap
func TestCheckMaxHeap(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	h := BuildMaxHeap(&A)

	validHeap := true
	prevMax, _ := h.ExtractMaximum()
	n := len(A)
	for j := 0; j < n; j++ {
		max, _ := h.ExtractMaximum()

		if prevMax < max {
			validHeap = false
			break
		}
	}

	test.AssertEqual(t, true, validHeap)
}

func TestCheckIncreaseKey(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	h := BuildMaxHeap(&A) // Build inplace

	val := 10000
	errorOrNil := h.MaxHeapInsert(&val)
	test.AssertEqual(t, nil, errorOrNil)
}
