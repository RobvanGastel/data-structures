package heap

import (
	"math/rand"
	"reflect"
	"testing"
)

// TODO: Create generic method, Checks if values are equal
func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}

	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

// TODO: Improve benchmark
func BenchmarkInsert(*testing.B) {
	A := make([]int, 0)
	// Fill max heap
	for i := 0; i < 1000; i++ {
		A = append(A, rand.Intn(30))
	}

	BuildMaxHeap(&A) // Build inplace
}

// Valid date if it builds a valid Max Heap
func TestCheckMaxHeap(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	BuildMaxHeap(&A) // Build inplace

	validHeap := true
	prevMax, _ := HeapExtractMax(&A)
	n := len(A)
	for j := 0; j < n; j++ {
		max, _ := HeapExtractMax(&A)

		if prevMax < max {
			validHeap = false
			break
		}
	}

	assertEqual(t, true, validHeap)
}

func TestCheckIncreaseKeyEmptyHeap(t *testing.T) {
	// TODO
}

func TestCheckIncreaseKey(t *testing.T) {
	A := make([]int, 0)
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	BuildMaxHeap(&A) // Build inplace

	val := 10000
	errorOrNil := MaxHeapInsert(&A, &val)
	assertEqual(t, nil, errorOrNil)
}
