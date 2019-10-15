package main

import (
	"fmt"
	"math/rand"

	. "./sort"
)

func main() {
	A := make([]int, 0)
	for i := 0; i < 10; i++ {
		randm := rand.Intn(1000)
		A = append(A, randm)
	}
	fmt.Println(A)
	HeapSort(&A) // Build inplace
	fmt.Println(A)
}
