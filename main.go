package main

import (
	. "data-structures/bst"
)

func main() {
	T := &Tree{}
	InitBinarySearchTree(T)
	InOrderTreeWalk(T.Root)

	x := TreeSearch(T.Root, 9)
	TreeDelete(T, x)

	InOrderTreeWalk(T.Root)
}
