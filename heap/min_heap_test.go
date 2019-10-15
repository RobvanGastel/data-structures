package heap

import (
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
	BuildMinHeap(&A) // Build inplace

	validHeap := true
	prevMin, _ := ExtractMinimum(&A)
	n := len(A)
	for j := 0; j < n; j++ {
		min, _ := ExtractMinimum(&A)

		if prevMin > min {
			validHeap = false
			break
		}
	}

	assertEqual(t, true, validHeap)
}

func TestCheckIncreaseMinKey(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	BuildMinHeap(&A) // Build inplace

	val := -10
	errorOrNil := MinHeapInsert(&A, &val)
	assertEqual(t, nil, errorOrNil)
}
