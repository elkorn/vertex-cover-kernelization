package graph

import "github.com/deckarep/golang-set"

// A matching, M , of G is a subset of the edges E, such that no vertex
// in V is incident to more that one edge in M .
// Intuitively we can say that no two edges in M have a common vertex.

// A matching M is said to be maximal if M is not properly contained in
// any other matching.
// Formally, M !⊂ M' for any matching M' of G.
// Intuitively, this is equivalent to saying that a matching is maximal if we cannot
// add any edge to the existing set
func maximalMatching(g *Graph) (matching mapset.Set, outsiders mapset.Set) {
	matching = mapset.NewSet()
	outsiders = mapset.NewSet()
	added := make([]bool, len(g.Vertices))
	g.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if !(added[edge.from.toInt()] || added[edge.to.toInt()]) {
			matching.Add(edge)
			added[edge.from.toInt()] = true
			added[edge.to.toInt()] = true
		} else {
			outsiders.Add(edge)
		}
	})

	return matching, outsiders
}

// Given G = (V, E) and a matching M of G, a vertex v is exposed,
// if no edge of M is incident with v.
func (self Vertex) isExposed(matching mapset.Set) bool {
	for element := range matching.Iter() {
		edge := element.(Edge)
		if edge.IsCoveredBy(self) {
			return false
		}
	}

	return true
}

// An augmenting path P is an alternating path that starts and ends
// at two distinct exposed vertices.
func isAugmentingPath(path []int, matching mapset.Set) bool {
	start := MkVertex(path[0])
	end := MkVertex(path[len(path)-1])

	return start.isExposed(matching) &&
		end.isExposed(matching) &&
		isAlternatingPathWithMatching(path, matching)
}
