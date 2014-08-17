package graph

import "github.com/deckarep/golang-set"

func assignWeightsToVertices(G *Graph, H *Graph, isCovered []bool) []float64 {
	border := len(G.Vertices)
	inVerboseContext(func() {
		Debug("verts: %v", G.Vertices)
	})
	weights := make([]float64, border)
	for _, Av := range G.Vertices {
		bvIndex := Av.toInt() + border
		avIndex := Av.toInt()
		inVerboseContext(func() {
			Debug("A: %v, B: %v (of %v)", avIndex, bvIndex, len(isCovered))
		})
		if isCovered[avIndex] {
			if isCovered[bvIndex] {
				weights[avIndex] = 1
			} else {
				weights[avIndex] = 0.5
			}
		} else {
			if isCovered[bvIndex] {
				weights[avIndex] = 0.5
			} else {
				weights[avIndex] = 0
			}
		}
	}

	return weights
}

func networkFlowKernelization(G *Graph, k int) /*(*Graph,*/ int /*)*/ {
	// Step 1: Convert graphg G to a bipartite graph H.
	// Step 2: Convert H to a network flow problem instance H'.
	hPrime := mkNetworkFlow(G)

	// Step 3: Find the maximum flow in H'.
	maxFlowPath, _ /* maxFlowValue */ := fordFulkerson(hPrime)

	// Step 4: The arcs in H' included in the instance of the maximum flow
	// 		   that correspond to edges in H constitute a matching set M of H.
	M := mapset.NewSet()
	matchedVertices := mapset.NewSet()
	// S is the set of all unmatched vertices in A.
	S := mapset.NewSet()
	for _, edge := range maxFlowPath {
		if edge.from != hPrime.source && edge.to != hPrime.sink {
			M.Add(*edge)
			if G.hasVertex(edge.from) {
				matchedVertices.Add(edge.from)
			} else {
				S.Add(edge.from)
			}

			if G.hasVertex(edge.to) {
				matchedVertices.Add(edge.to)
			} else {
				S.Add(edge.to)
			}
		}
	}

	// This acts as a map[int]bool.
	bipartiteCover := make([]bool, len(G.Vertices)*2)
	// Step 5: From M we cand find a vertex cover of H.
	// Case 1: all vertices are matched.
	if matchedVertices.Cardinality() == len(hPrime.graph.Vertices) {
		// Vertex cover of H is either the set A or B.
		for _, vertex := range G.Vertices {
			bipartiteCover[vertex.toInt()] = true
		}
	} else {
		// Case 2: not every vertex is included in the matching M.
		// R is the set of all vertices in A which are reachable
		// from S by alternating paths with respect to M.
		R := G.Vertices.reachableFromWithMatching(S, hPrime.net, M)

		// T is the set of neighbors of R along edges in M
		T := mapset.NewSet()
		for edgeInterface := range M.Iter() {
			edge := edgeInterface.(Edge)
			if R.Contains(edge.to) {
				T.Add(edge.from)
			}
			if R.Contains(edge.from) {
				T.Add(edge.to)
			}
		}

		A := mapset.NewSet()
		for _, g := range G.Vertices {
			A.Add(g)
		}

		// The vertex cover of the bipartite graph G' (this is a misprint in the paper I think - have to check other revisions)
		for v := range A.Difference(S).Difference(R).Union(T).Iter() {
			bipartiteCover[v.(Vertex).toInt()] = true
		}
	}

	// Step 6: Assign weights to all of the vertices of G, according to the vertex cover of H.
	weights := assignWeightsToVertices(G, hPrime.graph, bipartiteCover)

	Debug("weight: %v", weights)

	// Step 7: The remaining graph will be G' = (V', E') where V' = {v|W_v = 0.5} and k'=k-x where x = len({v|W_v = 1})
	x := 0
	for vIndex, weight := range weights {
		if 0.5 != weight {
			G.RemoveVertex(MkVertex(vIndex))
			if 1 == weight {
				x++
			}
		}
	}

	return /*G, */ k - x
}
