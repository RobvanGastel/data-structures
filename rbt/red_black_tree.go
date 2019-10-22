package rbt

// TODO: Implement all dynamic set operations tests

const (
	// RED color node constant
	RED = 0
	// BLACK color node constant
	BLACK = 1
)

// Node of the tree.
type Node struct {
	Left, Right, Parent *Node
	Color               int
	Key                 int
}

// Tree only storing the root node.
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
