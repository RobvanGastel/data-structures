package heap

import (
	"testing"
	"math/rand"
	"reflect"
)

// TODO: Create generic method, Checks if values are equal
func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	// debug.PrintStack()
	t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}

// TODO: Implement Tests on Insert, Delete, Transplant

func BenchmarkInsert(*testing.B) {
	A := []int{}

	// Fill max heap
	for i := 0; i< 1000; i++ {
		A = append(A, rand.Intn(30))
	}

	BuildMaxHeap(A) // Build inplace
}


func TestCheckMaxHeap(t *testing.T) {
	A := make([]int, 0)

	// Fill max heap
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	BuildMaxHeap(A) // Build inplace

	valid_heap := true
	prev_max, _ := HeapExtractMax(A)
	for j := 0; j < len(A); j++ {
		max, _ := HeapExtractMax(A)

		if prev_max > max {
			valid_heap = false
			break
		}
	}

    AssertEqual(t, true, valid_heap)
}

func TestCheckIncreaseKeyEmptyHeap(t *testing.T) {
	// A := make([]int, 0)

	// // Fill max heap
	// for i := 0; i < 1000; i++ {
	// 	randm := rand.Intn(1000)
	// 	A = append(A, randm)
	// }

	// fmt.Println(len(A))
	// BuildMaxHeap(A) // Build inplace

    // AssertEqual(t, true, valid_heap)
}

func TestCheckIncreaseKey(t *testing.T) {
	A := make([]int, 0)

	// Fill max heap
	for i := 0; i < 1000; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	BuildMaxHeap(A) // Build inplace

	val := 10000
	errorOrNil := MaxHeapInsert(A, &val)

    AssertEqual(t, nil, errorOrNil)
}