package kernelization

import "github.com/elkorn/vertex-cover-kernelization/utility"
import "github.com/elkorn/vertex-cover-kernelization/graph"

type domination struct {
	x, y   Vertex
	g      *Graph
	almost bool
}

func (u graph.Vertex) isDominatedBy(v Vertex, g *Graph) bool {
	return v.dominates(u, g)
}

func (v graph.Vertex) dominates(u Vertex, g *Graph) bool {
	/*
		Vertex u is said to be dominated by a vertex v , or alternatively,
		a vertex v is said to dominate a vertex u, if ( u , v) is an
		edge in G and N ( u ) ⊆ N [v] .
	*/
	if !g.HasEdge(u, v) {
		return false
	}

	result := true

	g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
		// For the whole N(u)...
		wu := getOtherVertex(u, edge)
		contains := false

		// We're dealing with N[v]
		if v == wu {
			utility.Debug("%v is in N[%v]", wu, v)
			return
		}

		g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			wv := getOtherVertex(v, edge)
			if wv == wu {
				utility.Debug("%v is in N[%v]", wu, v)
				contains = true
				done <- true
			}
		})

		// some vertex from N(u) does not belong in N[v].
		if !contains {
			utility.Debug("%v does not dominate %v", v, u)
			done <- true
			result = false
		}
	})

	if result {
		utility.Debug("%v dominates %v", v, u)
	}

	return result
}

func (u graph.Vertex) isAlmostDominatedBy(v Vertex, g *Graph) bool {
	return v.almostDominates(u, g)
}

func (v graph.Vertex) almostDominates(u Vertex, g *Graph) bool {
	/*
		A vertex v is said to almost-dominate a vertex u
		if u and v are non-adjacent and | N ( u ) − N (v)| ≤ 1.
	*/

	if g.HasEdge(u, v) {
		return false
	}

	_, vNeighbors := g.getNeighborsWithSet(v)
	_, uNeighbors := g.getNeighborsWithSet(u)

	diff := uNeighbors.Difference(vNeighbors).Cardinality()
	// utility.Debug("[ad-neighbors] u(%v): %v, v(%v): %v, diff: %v", u, uNeighbors, v, vNeighbors, diff)
	result := diff <= 1
	if result {
		utility.Debug("%v almost-dominates %v", v, u)
	} else {
		utility.Debug("%v does not almost-dominate %v", v, u)
	}
	return result
}
