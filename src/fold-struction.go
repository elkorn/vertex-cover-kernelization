package graph

import "fmt"

type structionVertex struct {
	v    Vertex
	i, j Vertex
}

func (self *structionVertex) Name() string {
	return fmt.Sprintf("v%v%v", self.i, self.j)
}

func struction(g *Graph, v0 Vertex) (result *Graph) {
	// Set of neighbors {v1, ... ,vp}
	s, sSet := g.getNeighborsWithSet(v0)
	p := sSet.Cardinality()
	Debug("Neighbors of %v: %v", v0, sSet)
	// newNodes := mapset.NewSet()
	newGraphCapacity := g.currentVertexIndex
	// TODO: this capacity is arbitrary, fix it.
	newVertices := make([]*structionVertex, 0, g.currentVertexIndex)
	newVertexLookup := make([][]Vertex, g.currentVertexIndex)
	for i, _ := range newVertexLookup {
		newVertexLookup[i] = make([]Vertex, g.currentVertexIndex)
	}

	// lookupVertex := func(i, j int) (result Vertex) {
	// 	result = newVertexLookup[i][j]
	// 	if INVALID_VERTEX == result {
	// 		result = newVertexLookup[j][i]
	// 	}
	//
	// 	return
	// }

	newDeletions := make([]bool, g.currentVertexIndex)
	copy(newDeletions, g.isVertexDeleted)

	// Remove vertices {v0, ... ,vp}
	newDeletions[v0.toInt()] = true
	for i := 0; i < p; i++ {
		vi := s[i]
		Debug("vi: %v, i: %v", vi, i)
		newDeletions[vi.toInt()] = true
		for j := i + 1; j < p; j++ {
			vj := s[j]
			Debug("vj: %v, j: %v", vj, j)
			// For each anti-edge (vi,vj) in G, where 0 < i < j <= p
			// introduce a new node vij.
			if !g.hasEdge(vi, vj) {
				newGraphCapacity++
				structionVertex := &structionVertex{
					v: Vertex(newGraphCapacity),
					i: vi,
					j: vj,
				}

				newVertices = append(newVertices, structionVertex)
				newVertexLookup[vi.toInt()][vj.toInt()] = structionVertex.v
				newVertexLookup[vj.toInt()][vi.toInt()] = structionVertex.v
			}
		}
	}

	Debug("Deletions: %v", newDeletions)

	result = MkGraphRememberingDeletedVertices(newGraphCapacity, newDeletions)

	addStructionNeighbors := func(i, v Vertex) {
		for _, neighbor := range newVertexLookup[i.toInt()] {
			if !(neighbor == INVALID_VERTEX || neighbor == v) {
				Debug("Adding struction edge %v-%v", neighbor, v)
				result.AddEdge(neighbor, v)
			}
		}
	}

	for _, newVertex := range newVertices {
		i, j := newVertex.i, newVertex.j
		Debug("i: %v, j: %v", i, j)
		addStructionNeighbors(i, newVertex.v)
		addStructionNeighbors(j, newVertex.v)
		// 1. Add an edge (vir, vjs) if i == j and g.hasEdge(vr,vs)
		// if i == j {
		// for vr, nr := range s {
		// 	for vs, ns := range s {
		// 		if nr == ns {
		// 			continue
		// 		}
		// 		if g.hasEdge(nr, ns) {
		// 			vir, vjs := lookupVertex(i, vr), lookupVertex(j, vs)
		// 			if INVALID_VERTEX == vir || INVALID_VERTEX == vjs {
		// 				// Don't know what to do here...
		// 				continue
		// 			}
		// 			Debug("i!=j, vir: (%v,%v), vjs: (%v,%v), adding edge %v-%v", i, vr, j, vs, vir, vjs)
		// 			result.AddEdge(vir, vjs)
		// 		}
		// 	}
		// }
		// } else {
		// Add an edge (vir, vjs) if i != j
		Debug("%v of %v", i, newVertex.Name())
		g.ForAllNeighbors(i, func(edge *Edge, done chan<- bool) {
			Debug("Edge %v-%v", edge.from, edge.to)
			result.AddEdge(newVertex.v, getOtherVertex(i, edge))
		})

		Debug("%v of %v", j, newVertex.Name())
		g.ForAllNeighbors(j, func(edge *Edge, done chan<- bool) {
			Debug("Edge %v-%v", edge.from, edge.to)
			result.AddEdge(newVertex.v, getOtherVertex(j, edge))
		})
		// }

		// For every vertex u not in {v0,...,vp},
		// if g.hasEdge(vi,u) or g.hasEdge(vj,u),
		// add an edge (vij, u)
		// g.ForAllVertices(func(u Vertex, done chan<- bool) {
		// 	if u == v0 || sSet.Contains(u) {
		// 		return
		// 	}
		//
		// 	if g.hasEdge(s[i], u) || g.hasEdge(s[j], u) {
		// 		result.AddEdge(newVertex.v, u)
		// 	}
		// })
	}

	return
}

func generalFold(g *Graph) *Graph {
	halt := make(chan bool, 1)
	crown := findCrown(g, halt, MAX_INT)
	if nil == crown {
		return g
	}

	result := MkGraphRememberingDeletedVertices(g.currentVertexIndex+1, g.isVertexDeleted)
	foldRoot := Vertex(result.currentVertexIndex)

	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		if crown.I.Contains(edge.from) ||
			crown.I.Contains(edge.to) ||
			crown.H.Contains(edge.from) ||
			crown.H.Contains(edge.to) {
			return
		}

		result.AddEdge(edge.from, edge.to)
	})

	// for vInter := range crown.I.Iter() {
	// 	v := vInter.(Vertex)
	// 	g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
	// 		result.AddEdge(foldRoot, getOtherVertex(v, edge))
	// 	})
	// }

	for vInter := range crown.H.Iter() {
		v := vInter.(Vertex)
		g.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			result.AddEdge(foldRoot, getOtherVertex(v, edge))
		})
	}

	reduceCrown(result, crown)
	Debug("H: %v", crown.H)

	return result
}

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
