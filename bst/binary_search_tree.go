package bst

import "fmt"

// The dynamic-set operations Search, Minimum, Maximum, Successor and
// Predecessor runs in O(h), height h.

// Node of the tree.
type Node struct {
	Left, Right, Parent *Node
	Key                 int
}

// Tree only storing the root node.
type Tree struct {
	Root *Node
}

// InsertNode into the Tree.
func (T *Tree) InsertNode(z *Node) {
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

// DeleteNode from the Tree.
func (T *Tree) DeleteNode(z *Node) {
	if z.Left == nil {
		T.Transplant(z, z.Right)
	} else if z.Right == nil {
		T.Transplant(z, z.Left)
	} else {
		var y = T.Minimum(z.Right)
		if y.Parent != z {
			T.Transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		T.Transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

// Transplant rotate on certain nodes, to
// reduce height.
func (T *Tree) Transplant(u, v *Node) {
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

// Successor retrieves sucessor of given Node.
func (T *Tree) Successor(x *Node) *Node {
	if x.Right != nil {
		return T.Minimum(x.Right)
	}

	var y = x.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

// Predecessor retrieves predecessor of given
// Node.
func (T *Tree) Predecessor(x *Node) *Node {
	if x.Left != nil {
		return T.Maximum(x.Left)
	}

	var y = x.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

// Minimum receives the minimum Node in the
// Tree.
func (T *Tree) Minimum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

// Maximum receives the maximum Node in the
// Tree.
func (T *Tree) Maximum(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

// Search the binary search Tree for given
// value.
func (T *Tree) Search(k int) *Node {
	return search(T.Root, k)
}

func search(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return search(x.Left, k)
	}
	return search(x.Right, k)
}

// TODO: Refactor tree walks

func InOrderTreeWalk(x *Node) {
	if x != nil {
		InOrderTreeWalk(x.Left)
		fmt.Println(x.Key)
		InOrderTreeWalk(x.Right)
	}
}

func PreOrderTreeWalk(x *Node) {
	if x != nil {
		fmt.Println(x.Key)
		PreOrderTreeWalk(x.Left)
		PreOrderTreeWalk(x.Right)
	}
}

func PostOrderTreeWalk(x *Node) {
	if x != nil {
		PostOrderTreeWalk(x.Left)
		PostOrderTreeWalk(x.Right)
		fmt.Println(x.Key)
	}
}
