package graph

type domination struct {
	x, y   Vertex
	g      *Graph
	almost bool
}

func (u Vertex) isDominatedBy(v Vertex, g *Graph) bool {
	return v.dominates(u, g)
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
			Debug("%v is in N[%v]", wu, v)
			return
		}

		g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			wv := getOtherVertex(v, edge)
			if wv == wu {
				Debug("%v is in N[%v]", wu, v)
				contains = true
				done <- true
			}
		})

		// some vertex from N(u) does not belong in N[v].
		if !contains {
			Debug("%v does not dominate %v", v, u)
			done <- true
			result = false
		}
	})

	if result {
		Debug("%v dominates %v", v, u)
	}

	return result
}

func (u Vertex) isAlmostDominatedBy(v Vertex, g *Graph) bool {
	return v.almostDominates(u, g)
}

func (v Vertex) almostDominates(u Vertex, g *Graph) bool {
	/*
		A vertex v is said to almost-dominate a vertex u
		if u and v are non-adjacent and | N ( u ) − N (v)| ≤ 1.
	*/

	if g.hasEdge(u, v) {
		return false
	}

	_, vNeighbors := g.getNeighborsWithSet(v)
	_, uNeighbors := g.getNeighborsWithSet(u)

	diff := uNeighbors.Difference(vNeighbors).Cardinality()
	// Debug("[ad-neighbors] u(%v): %v, v(%v): %v, diff: %v", u, uNeighbors, v, vNeighbors, diff)
	result := diff <= 1
	if result {
		Debug("%v almost-dominates %v", v, u)
	} else {
		Debug("%v does not almost-dominate %v", v, u)
	}
	return result
}
