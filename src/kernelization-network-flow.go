package graph

import "github.com/deckarep/golang-set"

func isAlternatingPathWithMatching(path []int, matching mapset.Set) bool {
	previous := Edge{path[0], path[1]}
	// P is an alternating path if:
	// 	- P is a path in G
	// 	- for each subsequent pair of edges from P one of them
	//    belongs to M and the other one does not.

	for i := 2; i < len(path); i++ {
		current := Edge{path[i-1], path[i]}
		inVerboseContext(func() {
			Debug("Checking %v:%v", previous, current)
		})

		if matching.Contains(e1) == matching.Contains(e2) {
			return false
		}

		previous := current
	}

	return true
}

func isReachableWithMatching(from, to Vertex, net Net, matching mapset.Set) bool {
	exists, path, _ := shortestPath(net, from, to)
	if len(path) <= 2 {
		// TODO I'm not sure whether in case where only a single edge is in the path,
		// the mere fact of it existing is enough - does it have to belong/not belong to M?
		return exists
	}

	if exists {
		return isAlternatingPathWithMatching(path, matching)
	}

	return false
}

func reachableFrom(toReach []Vertex, reachFrom mapset.Set, net Net, matching mapset.Set) mapset.Set {
	result := mapset.NewSet()
	for from := range reachFrom.Iter() {
		for _, to := range toReach {
			if result.Contains(to) {
				continue
			}

			if isReachableWithMatching(from, to, net, matching) {
				result.Add(to)
			}
		}
	}

	return result
}

func assignWeightsToVertices(G *Graph, H *Graph, isCovered []bool) []float64 {
	border := len(G.Vertices)
	weights := make([]float64, border)
	for _, Av := range G.Vertices {
		Bv := MkVertex(Av.toInt() + border)
		avIndex := Av.toInt()
		if isCovered[avIndex] {
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
	maxFlowPath, _ /* maxFlowValue */ := fordFulkerson(hPrime)

	Debug("Max. flow path: %v", maxFlowPath)
	// Step 4: The arcs in H' included in the instance of the maximum flow
	// 		   that correspond to edges in H constitute a matching set M of H.
	M := mapset.NewSet()
	matchedVertices := mapset.NewSet()
	// S is the set of all unmatched vertices in A.
	S := mapset.NewSet()
	inVerboseContext(func() {
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

		for m := range M.Iter() {
			Debug("m: %v", m)
		}
	})

	var bipartiteCover []bool
	// Step 5: From M we cand find a vertex cover of H.
	// Case 1: all vertices are matched.
	if matchedVertices.Cardinality() == len(hPrime.graph.Vertices) {
		bipartiteCover = make([]bool, len(G.Vertices))
		// Vertex cover of H is either the set A or B.
		for _, vertex := range G.Vertices {
			bipartiteCover[vertex.toInt()] = true
		}
	} else {
		bipartiteCover = make([]bool, 0)
		// Case 2: not every vertex is included in the matching.
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

	return G, k - x
}
