package graph

import "github.com/deckarep/golang-set"

func assignWeightsToVertices(G *Graph, H *Graph, isCovered []bool) []float64 {
	border := len(G.Vertices)
	weights := make([]float64, border)
	for Av := range G.Vertices {
		Bv := MkVertex(Av.toInt() + border)
		avIndex := Av.toInt()
		if isCovered[Av] {
			if isCovered[Bv] {
				weights[avIndex] = 1
			} else {
				weights[avIndex] = 0.5
			}
		} else {
			if isCovered[Bv] {
				weights[avIndex] = 0.5
			} else {
				weights[avIndex] = 0
			}
		}
	}

	return weights
}

func networkFlowKernelization(G *Graph, k int) (*Graph, int) {
	// Step 1: Convert graphg G to a bipartite graph H.
	// Step 2: Convert H to a network flow problem instance H'.
	hPrime := mkNetworkFlow(G)

	// Step 3: Find the maximum flow in H'.
	maxFlowPath, maxFlowValue := fordFulkerson(hPrime)

	// Step 4: The arcs in H' included in the instance of the maximum flow
	// 		   that correspond to edges in H constitute a matching set M of H.
	M := mapset.NewSet()
	numberOfMatchedOriginalVertices := 0
	for _, edge := range maxFlowPath {
		// TODO not sure if this condition is really taken care of in the maxFlow algorithm.
		M.Add(edge)
		if G.hasVertex(edge.from) || G.hasVertex(edge.to) {
			numberOfMatchedOriginalVertices++
		}
	}

	bipartiteCover := make([]bool, G.currentVertexIndex)

	// Step 5: From M we cand find a vertex cover of H.
	// Case 1: all vertices are matched.
	if numberOfMatchedOriginalVertices == len(hPrime.graph.Vertices) {
		// Vertex cover of H is either the set A or B.
		// TODO refactor when the map of Vertices gets switched to an array.
		for vertex := range G.Vertices {
			bipartiteCover[vertex.toInt()] = true
		}
	} else {
		// Case 2: not every vertex is included in the matching.
		// S is the set of all unmatched vertices in A.
		S := mapset.NewSet()
		for v := range G.Vertices {
			if !M.Contains(v) {
				S.Add(v)
			}
		}

	}

	// Step 6: Assign weights to all of the vertices of G, according to the vertex cover of H.
	weights := assignWeightsToVertices(G, hPrime.graph, bipartiteCover)

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

	return G, k - x
}
