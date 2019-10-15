package graph

// DIJKSTRA(G,w,s)
// INITIALIZE-SINGLE-SOURCE(G, s) S←∅
// Q←V[G]
// whileQ̸=∅
// do u ← EXTRACT-MIN(Q) S ← S ∪ {u}
// for each vertex v ∈ Adj[u] do RELAX(u,v,w)