package kernelization

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/matching"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

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

func findCrown(G *graph.Graph, halt chan<- bool, k int) *Crown {
	// Step 1.: Find a maximal matching M1 of the graph,
	// identify the set of all unmatched vertices as the set O of outsiders
	M1, O := matching.FindMaximalMatching(G)
	// If a maximum matching of size > k is found then
	// there is not a graph.vertex cover of size ≤ k and the graph.vertex cover problem
	// can be solved with a "no" instance.
	// If either the cardinality of M1 or M2 is > k, the process can be halted.
	utility.Debug("M1 cardinality: %v", M1.NEdges())
	utility.Debug("Outsiders: %v", O)
	// if options.Verbose {
	// 	M1.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
	// 		utility.Debug("%v", edge.Str())
	// 	})
	// }
	if M1.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 2.: Find a maximum aux. matching M2 of the edges between O and N(O)
	outsiderNeighbors := graph.MkGraph(G.CurrentVertexIndex)
	for vInter := range O.Iter() {
		v := vInter.(graph.Vertex)
		G.ForAllNeighbors(v, func(edge *graph.Edge, done chan<- bool) {
			outsiderNeighbors.AddEdge(edge.From, edge.To)
		})
	}

	M2 := matching.FindMaximumMatching(outsiderNeighbors)
	utility.Debug("M2 cardinality: %v", M2.NEdges())
	// if options.Verbose {
	// 	M2.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
	// 		utility.Debug("%v", edge.Str())
	// 	})
	// }

	// If either the cardinality of M1 or M2 is > k, the process can be halted.
	if M2.NEdges() > k {
		halt <- true
		return nil
	}

	// Step 3: If every graph.vertex in N(O) is matched by M2, then H=N(O) and I=O
	// form a straight crown, and we are done.
	straightCrown := true
	outsiderNeighbors.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		if outsiderNeighbors.IsVertexDeleted[v.ToInt()] {
			panic(fmt.Sprintf("Vertex %v is deleted!", v))
		}

		if !M2.HasVertex(v) || M2.Degree(v) == 0 {
			straightCrown = false
			done <- true
		}
	})

	if straightCrown {
		H := mapset.NewThreadUnsafeSet()
		outsiderNeighbors.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
			H.Add(v)
		})

		utility.Debug("Found a straight crown: I: %v, H: %v", O, H)
		return &Crown{
			H: H,
			I: O,
		}
	}

	// Step 4: Let I0 be the set of vertices in O that are unmatched by M2.
	In := mapset.NewThreadUnsafeSet()
	for vInter := range O.Iter() {
		v := vInter.(graph.Vertex)
		if !M2.HasVertex(v) || M2.Degree(v) == 0 {
			In.Add(v)
		}
	}

	if In.Cardinality() == 0 {
		utility.Debug("I0 is empty!")
		return &Crown{
			I: In,
			H: In,
		}
	}

	n := 0
	N := 0
	Isteps := make([]mapset.Set, 0, G.CurrentVertexIndex)
	Hsteps := make([]mapset.Set, 0, G.CurrentVertexIndex)
	Isteps = append(Isteps, In)
	// Step 5.:Repeat the following steps until n=N so that I_(N-1)=IN
	for {
		// 5a. Let Hn = N(In)
		Hsteps = append(Hsteps, mapset.NewThreadUnsafeSet())
		utility.Debug("n: %v, N: %v", n, N)
		utility.Debug("Isteps: %v", Isteps)
		for vInter := range In.Iter() {
			v := vInter.(graph.Vertex)
			G.ForAllNeighbors(v, func(edge *graph.Edge, done chan<- bool) {
				Hsteps[n].Add(graph.GetOtherVertex(v, edge))
			})
		}

		utility.Debug("Hsteps: %v", Hsteps)
		if n > 0 && Isteps[N].Equal(Isteps[N-1]) {
			break
		}

		// 5b. Let I_(n+1)= In ∪ N_M2(Hn)
		neighbors := mapset.NewThreadUnsafeSet()
		for vInter := range Hsteps[n].Iter() {
			v := vInter.(graph.Vertex)
			G.ForAllNeighbors(v, func(edge *graph.Edge, done chan<- bool) {
				w := graph.GetOtherVertex(v, edge)
				if M2.HasEdge(v, w) {
					// Adding N_M2(Hn)
					neighbors.Add(w)
				}
			})
		}

		// Adding In to I_n+1
		Isteps = append(Isteps, Isteps[n].Union(neighbors))
		utility.Debug("In: %v", Isteps[n])
		utility.Debug("In+1: %v", Isteps[n+1])
		utility.Debug("Hn: %v", Hsteps[n])

		n++
		N++
	}

	// This fixes the issues related to reducing a graph too much in some cases.
	// if Hsteps[N].Cardinality() == 0 {
	// 	utility.Debug("HN is empty!")
	// 	return &Crown{
	// 		I: Hsteps[N],
	// 		H: Hsteps[N],
	// 	}
	// }

	// Step 6: I_N, H_N form a flared crown.
	utility.Debug("Found a flared crown: I: %v, H: %v", Isteps[N], Hsteps[N])
	return &Crown{
		I: Isteps[N],
		H: Hsteps[N],
	}
}

func reduceCrown(G *graph.Graph, crown *Crown) {
	// TODO: There is a bug here - the algorithm keeps finding and reducing
	// crowns even if |I|=1 and |H|=0. This leads to a degradation of the graph
	// up to the point of having no vertices inside.
	// 1) Read Chlebik's paper for more details on how to fix this.
	//		- the paper proved to be inconclusive. Nothing is being said
	//		  about H, only that H=N(I) and I is not empty.
	// 2) Try using LP kernelization for finding crowns.
	utility.Debug("Removing crown %v", crown)
	// The graph G′ is produced by removing vertices in I and H
	// along with their adjacent edges.
	removeVerticesInSet := func(set mapset.Set) {
		for vInter := range set.Iter() {
			G.RemoveVertex(vInter.(graph.Vertex))
		}
	}

	removeVerticesInSet(crown.I)
	removeVerticesInSet(crown.H)
}

func ReduceCrown(G *graph.Graph, halt chan bool, k int) (kPrime int, partialCover mapset.Set) {
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
