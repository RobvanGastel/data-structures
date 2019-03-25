package main

import "fmt"

const (
	RED = 0
	BLACK = 1
)

type Node struct {
	left, right, parent *Node
	key int
}

type Tree struct {
	root *Node
}

func TreeInsert(T *Tree, z *Node) {
	var y *Node
	var x = T.root
	
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		T.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func InOrderTreeWalk(x *Node) {
	if x != nil {
		InOrderTreeWalk(x.left)
		fmt.Println(x.key)
		InOrderTreeWalk(x.right)
	}
}

func InitBinarySearchTree(T *Tree) {	
	TreeInsert(T, &Node{key: 10})
	TreeInsert(T, &Node{key: 8})
	TreeInsert(T, &Node{key: 12})
	TreeInsert(T, &Node{key: 6})
	TreeInsert(T, &Node{key: 14})
	TreeInsert(T, &Node{key: 4})
	TreeInsert(T, &Node{key: 9})
}


func main() {
	T := &Tree{}
	
	InitBinarySearchTree(T)
	InOrderTreeWalk(T.root)
}
