package kernelization

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type structionVertex struct {
	v    Vertex
	i, j Vertex
}

func (self *structionVertex) Name() string {
	return fmt.Sprintf("v%v%v", self.i, self.j)
}

func struction(g *graph.Graph, v0 Vertex) *graph.Graph {
	// Set of neighbors {v1, ... ,vp}
	s, sSet := g.getNeighborsWithSet(v0)
	result, _ := structionWithGivenNeighbors(g, v0, s, sSet)
	return result
}

func structionWithGivenNeighbors(g *graph.Graph, v0 Vertex, s Neighbors, sSet mapset.Set) (result *graph.Graph, reduction int) {
	p := sSet.Cardinality()
	newGraphCapacity := g.CurrentVertexIndex
	newVertices := make([]*structionVertex, 0, g.CurrentVertexIndex)
	newVertexLookup := make([][]Vertex, g.CurrentVertexIndex)
	for i, _ := range newVertexLookup {
		newVertexLookup[i] = make([]Vertex, g.CurrentVertexIndex)
	}
	newDeletions := make([]bool, g.CurrentVertexIndex)
	copy(newDeletions, g.IsVertexDeleted
)

	// Remove vertices {v0, ... ,vp}
	newDeletions[v0.ToInt()] = true
	reduction = 1
	for i := 0; i < p; i++ {
		vi := s[i]
		utility.Debug("vi: %v, i: %v", vi, i)
		newDeletions[vi.ToInt()] = true
		reduction++
		for j := i + 1; j < p; j++ {
			vj := s[j]
			utility.Debug("vj: %v, j: %v", vj, j)
			// For each anti-edge (vi,vj) in G, where 0 < i < j <= p
			// introduce a new node vij.
			if !g.HasEdge(vi, vj) {
				newGraphCapacity++
				reduction--
				structionVertex := &structionVertex{
					v: Vertex(newGraphCapacity),
					i: vi,
					j: vj,
				}

				newVertices = append(newVertices, structionVertex)
				newVertexLookup[vi.ToInt()][vj.ToInt()] = structionVertex.v
				newVertexLookup[vj.ToInt()][vi.ToInt()] = structionVertex.v
			}
		}
	}

	utility.Debug("Deletions: %v", newDeletions)

	result = MkGraphRememberingDeletedVertices(newGraphCapacity, newDeletions)
	for _, newVertex1 := range newVertices {
		for _, newVertex2 := range newVertices {
			if newVertex1 == newVertex2 {
				continue
			}

			// Add an edge (vir, vjs) if i == j and g.HasEdge(vr,vs)
			if newVertex1.i == newVertex2.i &&
				g.HasEdge(newVertex1.j, newVertex2.j) ||
				// Add an edge (vir, vjs) if i != j
				newVertex1.i != newVertex2.i {
				result.AddEdge(newVertex1.v, newVertex2.v)
			}
		}

		// For every vertex u not in {v0,...,vp},
		// if g.HasEdge(vi,u) or g.HasEdge(vj,u),
		// add an edge (vij, u)
		g.ForAllVertices(func(u Vertex, done chan<- bool) {
			if sSet.Contains(u) || u == v0 {
				return
			}

			for _, newVertex := range newVertices {
				if g.HasEdge(newVertex.i, u) || g.HasEdge(newVertex.j, u) {
					result.AddEdge(newVertex.v, u)
				}
			}
		})
	}

	return
}

func (self Vertex) isStructionApplicable(g *graph.Graph, neighbors mapset.Set) bool {
	p := neighbors.Cardinality()
	maxAllowedAntiEdges := p - 1
	numAntiEdges := 0
	for n := range neighbors.Iter() {
		n1 := n.(Vertex)
		for n := range neighbors.Iter() {
			n2 := n.(Vertex)
			if n1 == n2 {
				continue
			}

			if !g.HasEdge(n1, n2) {
				numAntiEdges++
				if numAntiEdges > maxAllowedAntiEdges {
					return false
				}
			}
		}
	}

	return true
}

// func kernelizeIfHasCoverOfSize(g *graph.Graph, k int) (hasCover bool, reduction int) {
// 	// Based on J. F. Buss and J. Goldsmith, SIAM 22, (1993), pp. 560-572.
// 	// 1.1. Let U be the set of vertices of degree more than k.
// 	U := mapset.NewSet()
// 	g.forAllVerticesOfDegreeGeq(k+1, func(v Vertex) {
// 		U.Add(v)
// 	})

// 	u := U.Cardinality()

// 	if u == 0 {

// 		return 0
// 	}

// 	utility.Debug("%v vertices of degree > %v", u, k)
// 	// 1.2. If |U| > k, then reject; there is no cover of size k or less.
// 	if u > k {
// 		hasCover = false
// 		reduction = 0
// 		return
// 	}

// 	// 2.1. Let G’ be the subgraph of G induced by V⧵U.
// 	// Every k-cover of G consists of U together with a k(k-|U|)-cover of G’.
// 	for v := range U.Iter() {
// 		g.RemoveVertex(v.(Vertex))
// 	}

// 	// 2.2. If G’ has more than k(k-|U|) edges, then reject; G’ has no (k-|U|)-cover.
// 	if g.NEdges() > k*(k-u) {
// 		utility.Debug("More than %v edges in subgraph, rejecting", k*(k-u))
// 		for v := range U.Iter() {
// 			g.RestoreVertex(v.(Vertex))
// 		}

// 		return 0
// 	}

// 	// 3. If G’ has a cover of size k IU[, then accept; otherwise reject.
// 	return k - u - kernelizeIfHasCoverOfSize(g, k-u)
// }

func generalFold(g *graph.Graph, halt chan bool, k int) (*graph.Graph, int) {
	// If the subroutine General-Fold() is not applicable to the
	// graph it means that its application does not change the structure of the
	// graph.
	kPrime := k
	var crown *Crown
	// Apply the NT-decomposition to G until the application of this
	// decomposition is trivial.
	// A decomposition is trivial when the crown is (∅, ∅).
	for {
		crown = findCrown(g, halt, k)
		// gv := MkGraphVisualizer(g)
		// gv.HighlightCrown(crown)
		// gv.Display()
		select {
		case <-halt:
			halt <- true
			return nil, k
		default:
		}

		if crown.IsTrivial() {
			break
		}

		reduceCrown(g, crown)
		kPrime = k - crown.Width()
	}

	// At this point, the graph is free of any non-trivial crowns.
	return reduceAlmostCrown(g, halt, kPrime)
}

func reduceAlmostCrown(g *graph.Graph, halt chan<- bool, kPrime int) (*graph.Graph, int) {
	// G′ has an almost-crown if and only if there exists a vertex v ∈ G′ such
	// that G′ ⧵ {v} has an equal crown.
	// For every vertex v in G′ , check if G′ ⧵ {v} has a crown.
	var crown *Crown
	var almostCrownVertex Vertex
	if g.NVertices() == 0 {
		utility.Debug("No vertices to search!")
		return g, kPrime
	}
	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		g.RemoveVertex(v)
		utility.Debug("Removed vertex %v, looking for a crown.", v)
		crown = findCrown(g, halt, kPrime)
		utility.Debug("restoring vertex %v", v)
		g.RestoreVertex(v)
		if !crown.IsTrivial() {
			// If the NT-decomposition yields a crown,
			// then this crown must be an equal crown.
			// Otherwise, the graph would not be crown-free.
			// Hence, we have constructed an almost-crown structure in G′.
			// According to the paper, this should take O(k^3\sqrt(k)), so there
			// shoul exist one such structure.
			almostCrownVertex = v
			done <- true
		}
	})
	// let G′ be the graph obtained from G by removing I ∪ N (I)
	// and adding a vertex u_I,
	// then connecting u_I to every vertex v ∈ G′
	// such that v was a neighbor of a vertex u ∈ N (I) in G.
	if almostCrownVertex != INVALID_VERTEX {
		// TODO: This is how it should look like as the result. Not sure if such
		// treatment is correct.
		// It might be OK, since as the result we are always removing I and N(I).
		// crown.H = crown.I.Clone()
		// crown.I.Clear()
		// crown.I.Add(almostCrownVertex)
		utility.Debug("Found a non-trivial almost-crown! H: %v, I: %v", crown.H, crown.I)
		g.addVertex()
		foldRoot := Vertex(g.CurrentVertexIndex)
		foldAndRemove := func(v Vertex) {
			g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
				g.AddEdge(foldRoot, getOtherVertex(v, edge))
			})

			g.RemoveVertex(v)
		}
		for vInter := range crown.H.Iter() {
			foldAndRemove(vInter.(Vertex))
		}

		for vInter := range crown.I.Iter() {
			foldAndRemove(vInter.(Vertex))
		}

		foldAndRemove(almostCrownVertex)
	}
	// Then τ( G\prime)=τ(G)−|I|.
	return g, kPrime - crown.I.Cardinality()
}

func foldStructionVC(G *graph.Graph, T *StructurePriorityQueueProxy, k int) {
	/*
			Therefore, when the algorithm branches on z, on the side of the branch where z is included,
		we can restrict our search to a minimum vertex cover that excludes at least two neighbors of N ( z ) , and we know that this
		is safe because if such a minimum vertex cover does not exist, then on the other side of the branch where N ( z ) has been
		included the algorithm will still be able to find a minimum vertex cover. Consequently, on the side of the branch where z is
		included, we can work under the assumption that at least two vertices in N ( z ) must be excluded. This working assumption
		will be stipulated by creating the tuple ( N ( z ), q = 2 ) .
	*/
}

/*
	---------------- IMPORTANT! ----------------
	It seems as though the algorithm creates *tuples*, yet T contains
	*structures*. This leads me to conclude that T should be initialized with
	a set of good pairs and vertices of degree \geq 7. The priorities have to be
	used accordingly.
	--------------------------------------------

	One question remains - if and how does the tuple updating operation affect
	other structures in T?

	T has structures -> algo creates tuples -> algo generates 2-tuples ->
	algo processes 2-tuples.

	Tuple
	What is a working assumption?

	What does 'stipulate', 'vacuously' mean?

	What happens to G when branching? It should be modified - check that in the
	pseudocode. This is to measure what 'implicit' means in the paper.

	How do structures relate to tuples?
	A tuple, a good pair or a vertex v with d(v) \geq 7 will be referred to as
	a structure.

	Conditional_Struction and Conditional_GeneralFold are applied when the
	reduction in k surpasses that resulting from branching on a certain
	tuple ( !!! in case it exists !!! ) -> that's what the paper says. Does
	that mean that when no tuple exists, struction and folding are applied
	to somehow get new tuples?

	---------------- TECH ----------------

	In the priority queue of tuples, it would be good to maintaina an `index`
	property for each one. This will be useful when updating tuples after each
	operation of the algorithm.

	A max. degree variable must be maintained within the graph - this will allow
	searching for degrees with d(v) \geq 7.

	StructurePriorityQueue should support an `update` operation, since the
	priority of structures will be changed dynamically.

	Golang enumeration - the structures should be enumerated by their type. The
		enum should also provide priority values after casting to int.
*/

// VC(G, T , k)
// 	Input: a graph G, a set T of tuples, and a positive integer k.
// 	Output: the size of a minimum vertex cover of G if the size is bounded by k;
// 	report failure otherwise.
// 	0. if |G| > 0 and k = 0 then reject;
// 	1. apply Reducing;
// 	2. pick a structure Γ of highest priority;
// 	3. if (Γ is a 2-tuple ({u, z}, 1)) or (Γ is a good pair (u, z) where z is
// 		almost-dominated by a vertex v ∈ N (u)) or (Γ is a vertex z with d(z) ≥ 7)
// 	then return
// 		min{1+VC(G − z, T ∪ (N (z), 2), k − 1), d(z)+ VC(G − N [z], T , k − d(z))};
// 	else /* Γ is a good pair (u, z) where z is not almost-dominated by by any
// 		vertex in N (u) */
// 		return
// 		min{1+VC(G − z, T , k − 1), d(z)+ VC(G − N [z], T ∪ (N (u), 2), k − d(z))};
//
// Reducing
// 	a. for each tuple (S, q) ∈ T do
// 		a.1. if |S| < q then reject;
// 		a.2. for every vertex u ∈ S do T = T ∪ {(S − {u}, q − 1)};
// 		a.3. if S is not an S independent set then
// 			T = T ∪ ( (u,v)∈E,u,v∈S {(S − {u, v}, q − 1)});
// 		a.4. if there exists v ∈ G such that |N (v) ∩ S| ≥ |S| − q + 1 then
// 			return (1+VC(G − v, T , k − 1)); exit;
// 	b. if Conditional General Fold(G) or Conditional Struction(G) in the
// 		given order is applicable then apply it; exit;
// 	c. if there are vertices u and v in G such that v dominates u then
// 		return (1+ VC(G − v, T , k − 1)); exit;
//
// Conditional General Fold
// 	if there exists a strong 2-tuple ({u, z}, 1) in T then
// 	if the repeated application of General Fold reduces the parameter by at
// 		least 2 then apply it repeatedly;
// 	else if the application of General-Fold reduces the parameter by 1 and
// 		(d(u) < 4)
// 	then apply it until it is no longer applicable;
// 	else apply General-Fold until it is no longer applicable;
//
// Conditional Struction
// 	if there exists a strong 2-tuple {u, v} in T then
// 	if there exists w ∈ {u, v} such that d(w) = 3 and the Struction is
// 		applicable to w then apply it;
// 	else if there exists a vertex u ∈ G where d(u) = 3 or d(u) = 4 and such that
// 		the Struction is applicable to u then apply it;
