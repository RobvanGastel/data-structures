package main

import (
	. "data-structures/heap"
	"fmt"
)

// func InitTree(T *Tree) {	
// 	T.InsertNode(&Node{Key: 10})
// 	T.InsertNode(&Node{Key: 9})
// 	T.InsertNode(&Node{Key: 8})
// 	T.InsertNode(&Node{Key: 6})
// 	T.InsertNode(&Node{Key: 11})
// 	T.InsertNode(&Node{Key: 12})
// 	T.InsertNode(&Node{Key: 13})
// 	T.InsertNode(&Node{Key: 5})
// 	T.InsertNode(&Node{Key: 4})
// }

func main() {
	// T := &Tree{}
	// node := &Node{Key: 4}
	// T.InsertNode(node)

	// InitTree(T)
	// // InOrderTreeWalk(T.Root)

	// var k = T.Predecessor(node)
	// fmt.Println(k)

	A := [10]int{30, 20, 15, 10, 5, 3, 25}

	// fmt.Println(A[Parent(3)])
	// fmt.Println(A[Left(0)])
	// fmt.Println(A[Right(0)])


	BuildMaxHeap()
	fmt.Println(A)
}
