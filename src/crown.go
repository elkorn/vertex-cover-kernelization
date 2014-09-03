package graph

import "github.com/deckarep/golang-set"

type Crown struct {
	I Vertices // An independent subset of G
	H Vertices // Head of the crown: N(I) - neighbors of I
	// Also, there must exist a matching M on the edges connecting I and H
	// such that all elements of H are matched.
}

func (self *Crown) Width() int {
	return len(self.H)
}

func findCrown(G *Graph) *Crown {
	// Step 1.: Find a maximal matching M1 of the graph,
	// identify the set of all unmatched vertices as the set O of outsiders
	M1, O := FindMaximalMatching(G)
	// Step 2.: Find a maximum aux. matching M2 of the edges between O and N(O)
	outsiderNeighbors := MkGraph(G.currentVertexIndex)
	for v := range O.Iter() {
		G.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
			if !outsiderNeighbors.hasEdge(edge.from, edge.to) {
				outsiderNeighbors.AddEdge(edge.from, edge.to)
			}
		})
	}

	M2 := FindMaximumMatching(outsiderNeighbors)
	// Step 3.: Let I0 be the set of vertices in O that are unmatched by M2.
	I0 := mapset.NewSet()
	for v := range O.Iter() {
		if M2.Degree(v) == 0 {
			I0.Add(v)
		}
	}

	n := 0
	// Step 4.:Repeat the following steps until n=N so that I_(N-1)=IN
	I1 := mapset.NewSet()
	for {
		// 4a. Let Hn = N(In)
		Hn := mapset.NewSet()
		for vInter := range I0 {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
				Hn.Add(getOtherVertex(v, edge))
			})
		}

		// 4b. Let I_(n+1)= In ∪ N_M2(Hn)
		neighbors := mapset.NewSet()
		for vInter := range Hn.Iter() {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
				w := getOtherVertex(v, edge)
				if !M2.hasEdge(v, w) {
					return
				}

				neighbors.Add(w)
			})
		}

		I1 = I0.Union(neighbors)
		if I1 == I0 {
			break
		}

		I0 = I1
		n++
	}
}
