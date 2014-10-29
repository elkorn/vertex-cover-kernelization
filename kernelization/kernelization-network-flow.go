package kernelization

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

func assignWeightsToVertices(G *graph.Graph, H *graph.Graph, isCovered []bool) []float64 {
	border := G.CurrentVertexIndex
	weights := make([]float64, border)
	for _, Av := range G.Vertices {
		bvIndex := Av.ToInt() + border
		avIndex := Av.ToInt()
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

func start(name string, args ...interface{}) {
	utility.Debug("start: "+name, args...)
}

func end(name string, args ...interface{}) {
	utility.Debug("end: "+name, args...)
}

func KernelizationNetworkFlow(G *graph.Graph, k int) /*(*graph.Graph,*/ int /*)*/ {
	// Step 1: Convert graphg G to a bipartite graph H.
	// Step 2: Convert H to a network flow problem instance H'.
	start("mkNetworkFlow")
	hPrime := mkNetworkFlow(G)
	end("mkNetworkFlow")

	// Step 3: Find the maximum flow in H'.
	start("foldFulkerson")
	maxFlowPath, _ /* maxFlowValue */ := fordFulkerson(hPrime)
	end("foldFulkerson")

	// Step 4: The arcs in H' included in the instance of the maximum flow
	// 		   that correspond to edges in H constitute a matching set M of H.
	M := mapset.NewThreadUnsafeSet()
	matchedVertices := mapset.NewThreadUnsafeSet()
	// S is the set of all unmatched vertices in A.
	S := mapset.NewThreadUnsafeSet()
	start("Processing max flow path of length %v", len(maxFlowPath))
	for _, edge := range maxFlowPath {
		if edge.From != hPrime.source && edge.To != hPrime.sink {
			M.Add(*edge)
			if G.HasVertex(edge.From) {
				matchedVertices.Add(edge.From)
			} else {
				S.Add(edge.From)
			}

			if G.HasVertex(edge.To) {
				matchedVertices.Add(edge.To)
			} else {
				S.Add(edge.To)
			}
		}
	}
	end("Processing max flow path of length %v", len(maxFlowPath))

	// This acts as a map[int]bool.
	bipartiteCover := make([]bool, G.CurrentVertexIndex*2)
	// Step 5: From M we cand find a vertex cover of H.
	// Case 1: all vertices are matched.
	utility.Debug("Matched: %v of %v", matchedVertices.Cardinality(), hPrime.graph.NVertices())
	if matchedVertices.Cardinality() == hPrime.graph.NVertices() {
		// Vertex cover of H is either the set A or B.
		for _, vertex := range G.Vertices {
			bipartiteCover[vertex.ToInt()] = true
		}
	} else {
		// Case 2: not every vertex is included in the matching M.
		// R is the set of all vertices in A which are reachable
		// from S by alternating paths with respect to M.
		start("Getting reachable vertices")
		// TODO: this cannot work in this context when edges are compared
		// by pointer values when checking whether the path is alternating.
		// Investigate.
		R := reachableFromWithMatching(G.Vertices, S, hPrime, M)
		end("Getting reachable vertices")

		// T is the set of neighbors of R along edges in M
		T := mapset.NewThreadUnsafeSet()
		for edgeInterface := range M.Iter() {
			edge := edgeInterface.(graph.Edge)
			if R.Contains(edge.To) {
				T.Add(edge.From)
			}
			if R.Contains(edge.From) {
				T.Add(edge.To)
			}
		}

		A := mapset.NewThreadUnsafeSet()
		for _, g := range G.Vertices {
			A.Add(g)
		}

		// The vertex cover of the bipartite graph G' (this is a misprint in the paper I think - have to check other revisions)
		for v := range A.Difference(S).Difference(R).Union(T).Iter() {
			bipartiteCover[v.(graph.Vertex).ToInt()] = true
		}
	}

	// Step 6: Assign weights to all of the vertices of G, according to the vertex cover of H.
	weights := assignWeightsToVertices(G, hPrime.graph, bipartiteCover)

	utility.Debug("weight: %v\n", weights)

	// Step 7: The remaining graph will be G' = (V', E') where V' = {v|W_v = 0.5} and k'=k-x where x = len({v|W_v = 1})
	x := 0
	for vIndex, weight := range weights {
		if 0.5 != weight {
			G.RemoveVertex(graph.MkVertex(vIndex))
			if 1 == weight {
				x++
			}
		}
	}

	return /*G, */ x
}
