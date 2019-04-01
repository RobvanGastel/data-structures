package graph

// Directed and Undirected Graph
// Adjacency Matrix / Adjacency List

type Graph struct {
	nodes []*Node
}

type Node struct {
	edges []*Edge
	Index int
	Value int
}

type Edge struct {
	Weight int
	Start Node
	End Node
}

func (G *Graph) MakeEdge(from, to Node) {

}

func (G *Graph) MakeNode() {

}
