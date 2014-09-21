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

func (self *Crown) IsTrivial() bool {
	return self.Width() == 0 && self.I.Cardinality() == 0
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
	Debug("M1 cardinality: %v", M1.NEdges())
	Debug("Outsiders: %v", O)
	if options.Verbose {
		M1.ForAllEdges(func(edge *Edge, done chan<- bool) {
			Debug("%v", edge.Str())
		})
	}
	if M1.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 2.: Find a maximum aux. matching M2 of the edges between O and N(O)
	outsiderNeighbors := MkGraph(G.currentVertexIndex)
	for vInter := range O.Iter() {
		v := vInter.(Vertex)
		G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			outsiderNeighbors.AddEdge(edge.from, edge.to)
		})
	}

	M2 := FindMaximumMatching(outsiderNeighbors)
	Debug("M2 cardinality: %v", M2.NEdges())
	if options.Verbose {
		M2.ForAllEdges(func(edge *Edge, done chan<- bool) {
			Debug("%v", edge.Str())
		})
	}

	// If either the cardinality of M1 or M2 is > k, the process can be halted.
	if M2.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 3.: Let I0 be the set of vertices in O that are unmatched by M2.
	// TODO: handle the Invariant: I0.Cardinality() > 0
	In := mapset.NewSet()
	for vInter := range O.Iter() {
		v := vInter.(Vertex)
		if deg, _ := M2.Degree(v); deg == 0 {
			In.Add(v)
		}
	}

	if In.Cardinality() == 0 {
		Debug("I0 is empty!")
		return &Crown{
			I: In,
			H: In,
		}
	}

	n := 0
	N := 0
	Isteps := make([]mapset.Set, 0, 1999)
	Hsteps := make([]mapset.Set, 0, 1999)
	Isteps = append(Isteps, In)
	// Step 4.:Repeat the following steps until n=N so that I_(N-1)=IN
	for {
		// 4a. Let Hn = N(In)
		Hsteps = append(Hsteps, mapset.NewSet())
		Debug("n: %v, N: %v", n, N)
		Debug("Isteps: %v", Isteps)
		for vInter := range In.Iter() {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
				Hsteps[n].Add(getOtherVertex(v, edge))
			})
		}

		Debug("Hsteps: %v", Hsteps)
		if n > 0 && Isteps[N].Equal(Isteps[N-1]) {
			break
		}

		// 4b. Let I_(n+1)= In ∪ N_M2(Hn)
		neighbors := mapset.NewSet()
		for vInter := range Hsteps[n].Iter() {
			v := vInter.(Vertex)
			G.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
				w := getOtherVertex(v, edge)
				if M2.hasEdge(v, w) {
					// Adding N_M2(Hn)
					neighbors.Add(w)
				}

			})
		}

		// Adding In to I_n+1
		Isteps = append(Isteps, Isteps[n].Union(neighbors))
		Debug("In: %v", Isteps[n])
		Debug("In+1: %v", Isteps[n+1])
		Debug("Hn: %v", Hsteps[n])

		n++
		N++
	}

	// The result is (I_N, H_N).
	return &Crown{
		I: Isteps[N],
		H: Hsteps[N],
	}
}

func reduceCrown(G *Graph, crown *Crown) {
	Debug("Removing crown %v", crown)
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
