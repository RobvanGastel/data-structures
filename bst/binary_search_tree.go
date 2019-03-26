package bst

import "fmt"

const (
	RED = 0
	BLACK = 1
)

type Node struct {
	Left, Right, Parent *Node
	Key int
}

type Tree struct {
	Root *Node
}

func TreeInsert(T *Tree, z *Node) {
	var y *Node
	var x = T.Root
	
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == nil {
		T.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

func TreeSearch(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return TreeSearch(x.Left, k)
	} else {
		return TreeSearch(x.Right, k)
	}
}

func Transplant(T *Tree, u, v *Node) {
	if u.Parent == nil {
		T.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func TreeMinimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func TreeDelete(T *Tree, z *Node) {
	if z.Left == nil {
		Transplant(T, z, z.Right)
	} else if z.Right == nil {
		Transplant(T, z, z.Right)
	} else {
		y := TreeMinimum(z.Right)
		if y.Parent != z {
			Transplant(T, y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		Transplant(T, z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

func InOrderTreeWalk(x *Node) {
	if x != nil {
		InOrderTreeWalk(x.Left)
		fmt.Println(x.Key)
		InOrderTreeWalk(x.Right)
	}
}

func InitBinarySearchTree(T *Tree) {	
	TreeInsert(T, &Node{Key: 10})
	TreeInsert(T, &Node{Key: 8})
	TreeInsert(T, &Node{Key: 12})
	TreeInsert(T, &Node{Key: 6})
	TreeInsert(T, &Node{Key: 14})
	TreeInsert(T, &Node{Key: 4})
	TreeInsert(T, &Node{Key: 9})
}
