package graph

import "github.com/deckarep/golang-set"

type Crown struct {
	I mapset.Set // An independent, non-empty subset of G
	H mapset.Set // Head of the crown: N(I) - neighbors of I
	// Also, there must exist a matching M on the edges connecting I and H
	// such that all elements of H are matched.
}

func (self *Crown) Width() int {
	return self.H.Cardinality()
}

func (self *graphVisualizer) highlightCrown(crown *Crown) {
	for vInter := range crown.I.Iter() {
		self.HighlightVertex(vInter.(Vertex), "lightgray")
	}

	for vInter := range crown.H.Iter() {
		self.HighlightVertex(vInter.(Vertex), "yellow")
	}
}

func findCrown(G *Graph, halt chan<- bool, k int) *Crown {
	// Step 1.: Find a maximal matching M1 of the graph,
	// identify the set of all unmatched vertices as the set O of outsiders
	M1, O := FindMaximalMatching(G)
	// If a maximum matching of size > k is found then
	// there is not a vertex cover of size ≤ k and the vertex cover problem
	// can be solved with a "no" instance.
	// If either the cardinality of M1 or M2 is > k, the process can be halted.
	if M1.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 2.: Find a maximum aux. matching M2 of the edges between O and N(O)
	outsiderNeighbors := MkGraph(G.currentVertexIndex)
	for vInter := range O.Iter() {
		v := vInter.(Vertex)
		G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			if !outsiderNeighbors.hasEdge(edge.from, edge.to) {
				outsiderNeighbors.AddEdge(edge.from, edge.to)
			}
		})
	}

	M2 := FindMaximumMatching(outsiderNeighbors)
	// If either the cardinality of M1 or M2 is > k, the process can be halted.
	if M2.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 3.: Let I0 be the set of vertices in O that are unmatched by M2.
	// TODO: handle the Invariant: I0.Cardinality() > 0
	if O.Cardinality() == 0 {
		Debug("Outsiders is empty!")
	}
	I0 := mapset.NewSet()
	for vInter := range O.Iter() {
		v := vInter.(Vertex)
		if deg, _ := M2.Degree(v); deg == 0 {
			I0.Add(v)
		}
	}

	n := 0
	// Step 4.:Repeat the following steps until n=N so that I_(N-1)=IN
	I1 := mapset.NewSet()
	var Hn mapset.Set
	for {
		// 4a. Let Hn = N(In)
		Hn = mapset.NewSet()
		for vInter := range I0.Iter() {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
				Hn.Add(getOtherVertex(v, edge))
			})
		}

		// 4b. Let I_(n+1)= In ∪ N_M2(Hn)
		neighbors := mapset.NewSet()
		for vInter := range Hn.Iter() {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
				w := getOtherVertex(v, edge)
				if !M2.hasEdge(v, w) {
					return
				}

				neighbors.Add(w)
			})
		}

		I1 = I0.Union(neighbors)
		Debug("I0: %v", I0)
		Debug("I1: %v", I1)
		Debug("Hn: %v", Hn)
		if I1.Equal(I0) {
			break
		}

		I0 = I1
		n++
	}

	return &Crown{
		I: I1,
		H: Hn,
	}
}

func reduceCrown(G *Graph, crown *Crown) {
	// The graph G′ is produced by removing vertices in I and H
	// along with their adjacent edges.
	removeVerticesInSet := func(set mapset.Set) {
		for vInter := range set.Iter() {
			G.RemoveVertex(vInter.(Vertex))
		}
	}

	removeVerticesInSet(crown.I)
	removeVerticesInSet(crown.H)
}

func ReduceCrown(G *Graph, halt chan bool, k int) (kPrime int, partialCover mapset.Set) {
	crown := findCrown(G, halt, k)
	select {
	case <-halt:
		halt <- true
		return -1, nil
	default:
	}

	reduceCrown(G, crown)

	// The problem size becomes n′= n - (|I|+|H|)
	// and the parameter size is k′ = k - |H|.
	kPrime = k - crown.Width()
	partialCover = crown.H
	return kPrime, partialCover
}
