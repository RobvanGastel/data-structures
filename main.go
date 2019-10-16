package main

import (
	"fmt"
	"math/rand"

	// . "./heap"
	. "./sort"
)

func main() {
	A := make([]int, 0)
	for i := 0; i < 10; i++ {
		randm := rand.Intn(10000)
		A = append(A, randm)
	}
	fmt.Println(A)
	HeapSort(&A, "asc") // Build inplace
	fmt.Println(A)

	// Fill max heap
	// A := make([]int, 0)
	// for i := 0; i < 1000; i++ {
	// 	A = append(A, rand.Intn(30))
	// }
	// fmt.Println(A)
	// heap := BuildMaxHeap(&A)
	// fmt.Println(A)
}
