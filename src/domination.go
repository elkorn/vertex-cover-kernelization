package graph

type domination struct {
	x, y   Vertex
	g      *Graph
	almost bool
}

func (v Vertex) dominates(u Vertex, g *Graph) bool {
	/*
		Vertex u is said to be dominated by a vertex v , or alternatively,
		a vertex v is said to dominate a vertex u, if ( u , v) is an
		edge in G and N ( u ) ⊆ N [v] .
	*/
	if !g.hasEdge(u, v) {
		return false
	}

	result := true

	g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
		// For the whole N(u)...
		wu := getOtherVertex(u, edge)
		contains := false

		// We're dealing with N[v]
		if v == wu {
			done <- true
			return
		}

		g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			wv := getOtherVertex(v, edge)
			if wv == wu {
				contains = true
				done <- true
			}
		})

		// some vertex from N(u) does not belong in N[v].
		if !contains {
			done <- true
			result = false
		}
	})

	return result
}

func (v Vertex) almostDominates(u Vertex, g *Graph) bool {
	/*
		1 A vertex u is said to be
		almost-dominated by a vertex v , or alternatively,
		a vertex v is said to almost-dominate a vertex u,
		if u and v are non-adjacent and | N ( u ) − N (v)| ≤ 1.
	*/

	if g.hasEdge(u, v) {
		return false
	}

	vNeighbors, uNeighbors := g.getNeighbors(v), g.getNeighbors(u)
	Debug("[Neighbors] u: %v, v: %v", uNeighbors, vNeighbors)
	return IntAbs(len(uNeighbors)-len(vNeighbors)) <= 1
}
