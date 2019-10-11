package heap

import (
	"math"
	"errors"
	"fmt"
)

// TODO: Write tests for heaps

func Parent(i int) int {
	return int(math.Floor(float64(i-1)/2.0))
}

func Left(i int) int {
	return 2*i
}

func Right(i int) int {
	return 2*i+1
}

func HeapMaximum(A []int) int {
	return A[1]
}

func MaxHeapify(A []int, i int) {
	l := Left(i)
	r := Right(i)
	largest := 0

	fmt.Println(largest)
	fmt.Println(r)
	fmt.Println(l)

	if l <= len(A) && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}

	if r <= len(A) && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		temp := A[largest]
		A[largest] = A[i]
		A[i] = temp
		MaxHeapify(A, largest)
	}
}

func BuildMaxHeap(A []int) {
	A_floor := int(math.Floor(float64(len(A)))/2.0)
	for i := A_floor; i > 1; i-- {
		fmt.Println(A_floor, i)
		MaxHeapify(A, i)
	}
}

func HeapExtractMax(A []int) (int, error) {
	if len(A) < 1 {
		return 0, errors.New("heap underflow")
	}

	max := A[1]
	MaxHeapify(A, 1)
	return max, nil
}

func HeapIncreaseKey(A []int, i int, key int) error {
	if key < A[i] {
		return errors.New("new key is smaller than current key")
	}

	A[i] = key
	for i > 1 && A[Parent(i)] < A[i] {
		temp := A[Parent(i)]
		A[Parent(i)] = A[i]
		A[i] = temp

		i = Parent(i)
	}
	return nil
}

func MaxHeapInsert(A []int, key int) {
	A = append(A, -2147483648)
	HeapIncreaseKey(A, len(A), key)
}