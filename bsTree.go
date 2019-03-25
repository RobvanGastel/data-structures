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

func TreeSearch(x *Node, k int) *Node {
	if x == nil || k == x.key {
		return x
	}
	if k < x.key {
		return TreeSearch(x.left, k)
	} else {
		return TreeSearch(x.right, k)
	}
}

func Transplant(T *Tree, u, v *Node) {
	if u.parent == nil {
		T.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func TreeMinimum(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func TreeDelete(T *Tree, z *Node) {
	if z.left == nil {
		Transplant(T, z, z.right)
	} else if z.right == nil {
		Transplant(T, z, z.right)
	} else {
		y := TreeMinimum(z.right)
		if y.parent != z {
			Transplant(T, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		Transplant(T, z, y)
		y.left = z.left
		y.left.parent = y
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

	x := TreeSearch(T.root, 9)
	TreeDelete(T, x)

	InOrderTreeWalk(T.root)
}
