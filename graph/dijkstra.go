package graph

import "sort"

// CLRS example of Dijkstra
// Dijkstra(G,w,s)
// INITIALIZE-SINGLE-SOURCE(G, s) S←∅
// Q←V[G]
// whileQ̸=∅
// do u ← EXTRACT-MIN(Q) S ← S ∪ {u}
// for each vertex v ∈ Adj[u] do RELAX(u,v,w)

// Dijkstra shortest path algorithm for Graph g and start node s.
func Dijkstra(G *Graph, s *Node) {

	source := InitializeSingleSource(G, s)
	var visited []*Node

	for len(visited) != len(G.Nodes) {
		node := getClosestNonVisitedNode(source, visited)

		visited = append(visited, node)

		nodeEdges := G.GetNodeEdges(node)

		for _, edge := range nodeEdges {

			distanceToNeighbor := source[node] + edge.Weight
			if distanceToNeighbor < source[edge.Child] {
				source[edge.Child] = distanceToNeighbor
			}
		}
	}
}

// InitializeSingleSource needs a graph with a start Node that needs
// to be in that graph.
func InitializeSingleSource(G *Graph, s *Node) map[*Node]int {
	table := make(map[*Node]int)
	table[s] = 0

	for _, node := range G.Nodes {
		if node != s {
			table[node] = int(^uint(0) >> 1)
		}
	}

	return table
}

// getClosestNonVisitedNode for a given source and visited nodes.
func getClosestNonVisitedNode(source map[*Node]int, visited []*Node) *Node {
	type SourceSorted struct {
		Node   *Node
		Weight int
	}
	var sorted []SourceSorted

	for node, weight := range source {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, SourceSorted{node, weight})
		}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight < sorted[j].Weight
	})

	return sorted[0].Node
}
