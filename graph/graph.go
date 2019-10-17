package graph

// Graph ...
type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// Edge ...
type Edge struct {
	Parent *Node
	Child  *Node
	Weight int
}

// Node ...
type Node struct {
	Name string
}

// AddEdge ...
func (G *Graph) AddEdge(parent, child *Node, weight int) {
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Weight: weight,
	}

	G.Edges = append(G.Edges, edge)
	G.AddNode(parent)
	G.AddNode(child)
}

// AddNode ...
func (G *Graph) AddNode(node *Node) {
	var isPresent bool
	for _, n := range G.Nodes {
		if n == node {
			isPresent = true
		}
	}

	if !isPresent {
		G.Nodes = append(G.Nodes, node)
	}
}

// ToString ...
func (G *Graph) ToString() {
	// TODO
}
