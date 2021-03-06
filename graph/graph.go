package graph

import "strconv"

// Graph directed graph struct
type Graph struct {
	Nodes map[*Node]bool
	Edges []*Edge
}

// Edge struct as edge between Child and Parent With Weight
// in a Graph.
type Edge struct {
	Parent *Node
	Child  *Node
	Weight int
}

// Node struct with Name representing the Node.
type Node struct {
	Name string
}

// AddEdge between to nodes in a graph.
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

// AddNode to the graph.
func (G *Graph) AddNode(node *Node) {
	if G.Nodes == nil {
		G.Nodes = make(map[*Node]bool)
	}

	_, isPresent := G.Nodes[node]
	if !isPresent {
		G.Nodes[node] = true
	}
}

// GetNodeEdges returns all the Edges that start with the
// specified Node. In other terms, returns all the Edges
// connecting to the Node's neighbors
func (G *Graph) GetNodeEdges(node *Node) (edges []*Edge) {
	for _, edge := range G.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

// ToString converts the Graph to string
// representation.
func (G *Graph) ToString() string {
	var s string

	s += "Edges:\n"
	for _, edge := range G.Edges {
		s += edge.Parent.Name + " -> " + edge.Child.Name + " = " + strconv.Itoa(edge.Weight)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	i := 0
	for node := range G.Nodes {
		if i == len(G.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
		i++
	}
	s += "\n"

	return s
}
