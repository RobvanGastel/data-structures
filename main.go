package main

import (
	. "data-structures/heap"
	"fmt"
)

func main() {
	A := []int{30, 20, 100, 200, 300, 5100, 510, 20, 20, 15, 10, 5, 3, 25}

	BuildMaxHeap(&A) // Build inplace

	prev_max, _ := HeapExtractMax(&A)
	fmt.Println(A)
	fmt.Println(prev_max)
	n := len(A)
	for j := 0; j < n; j++ {
		max, _ := HeapExtractMax(&A)
		fmt.Println(max)
		fmt.Println(j)
		if prev_max < max {
			fmt.Println("dsadadad")
			break
		}
	}
}
