package rbt

import (
	"testing"
)

func BenchmarkInsert(*testing.B) {
	T := &Tree{}
	for i := 0; i< 1000; i++ {
		n := &Node{Key: i}
		T.InsertNode(n)
	}
}