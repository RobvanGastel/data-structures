package bst

import (
	"testing"
)

// TODO: Implement Tests on Insert, Delete, Transplant

func BenchmarkInsert(*testing.B) {
	T := &Tree{}
	for i := 0; i < 1000; i++ {
		n := &Node{Key: i}
		T.InsertNode(n)
	}
}
