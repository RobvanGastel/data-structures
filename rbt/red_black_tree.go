package rbt

import "fmt"

// TODO: Implement all dynamic set operations

const (
	RED = 0
	BLACK = 1
)

type Node struct {
	Left, Right, Parent *Node
	Color int
	Key int
}

type Tree struct {
	Root *Node
}

// TODO: Translate to Red Black tree insert
func (T *Tree) InsertNode(z *Node) {
	var y *Node
	var x = T.Root
	
	if x == nil {
		z.Color = BLACK
		T.Root = z
	}

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

func InOrderTreeWalk(x *Node) {
	if x != nil {
		InOrderTreeWalk(x.Left)
		fmt.Println(x.Key)
		InOrderTreeWalk(x.Right)
	}
}

