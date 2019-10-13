package main

import (
	. "data-structures/heap"
	"fmt"
)

func main() {
	A := []int{30, 20, 15, 10, 5, 3, 25}

	BuildMaxHeap(&A) // Build inplace
	fmt.Println(A)

	val := 10000
	_ = MaxHeapInsert(&A, &val)
	fmt.Println(A)

	i, _ := HeapExtractMax(&A)
	fmt.Println(i)

	fmt.Println(A)
	i, _ = HeapExtractMax(&A)
	fmt.Println(i)
	fmt.Println(A)
}
