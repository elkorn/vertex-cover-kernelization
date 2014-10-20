package kernelization

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type nt_answer int

const (
	NT_OK nt_answer = iota
	NT_MIN_COVER_TOO_BIG
	NT_V12_NOT_INDEPENDENT
	NT_V12_TOO_BIG
)

type ntDecomposition struct {
	// Both original vertices and their copies belong to S.
	V1 mapset.Set
	// Neither original vertices nor their copies belong to S.
	V0 mapset.Set
	// Either original vertices or their copies belong to S (but not both).
	V12    mapset.Set
	K      int
	Answer nt_answer
}

func mkNtDecomposition(g *graph.Graph, k int) (result *ntDecomposition) {
	// def NT(G, k) :
	// 1. let U 1 = V (G) and let U 2 be a new set of vertices where |U 1 | = |U 2 |
	// 2. let σ : U 1 → U 2 be a bijection from U 2 to U 1
	// 3. let H = (U 1 ∪ U 2 , {{x, y} | x ∈ U 1 , y ∈ U 2 , and {x, σ(y)} ∈ E(G)})
	// 4. let S be a vertex cover of H (Ford-Fulkerson)
	border := graph.Vertex(g.CurrentVertexIndex)
	maxFlow, _ := fordFulkerson(mkNetworkFlow(g))
	S := mapset.NewSet()
	for _, edge := range maxFlow {
		if S.Contains(edge.From) {
			S.Add(edge.To)
		} else {
			S.Add(edge.From)
		}
	}

	result = &ntDecomposition{
		V1: mapset.NewSet(),
		V0: mapset.NewSet(),
	}

	// 5. let V 1 = {x | x ∈ S ∩ U 1 and σ −1 (x) ∈ S}
	U1 := g.Vertices.toSet()
	U1minusS := U1.Difference(S)
	for vInter := range S.Iter() {
		v := vInter.(graph.Vertex)
		if S.Contains(v + border) {
			if g.HasVertex(v) {
				result.V1.Add(v)
			}
		} else if U1minusS.Contains(v) {
			// 6. let V 0 = {x | x ∈ U 1 − S and σ −1 (x) %∈ S}
			result.V0.Add(v)
		}
	}

	result.K -= result.V1.Cardinality()
	// 7. let V 12 = U 1 − V 1 − V 0
	result.V12 = U1.Difference(result.V1).Difference(result.V0)
	if result.V1.Cardinality() > k {
		// 8. if |V 1 | > k: return NO1
		result.Answer = NT_MIN_COVER_TOO_BIG
		return
	} else if result.V1.Cardinality() == k {
		if indep, _ := isIndependentSet(result.V12, g); !indep {
			// 9. if |V 1 | = k and E(G[V 12 ]) % = ∅: return NO2
			result.Answer = NT_V12_NOT_INDEPENDENT
			return
		}
		// 10. if |V 1 | = k and E(G[V 12 ]) = ∅: return YES
		return
	}

	if result.V12.Cardinality() > 2*(k-result.V1.Cardinality()) {
		// 11. if |V 12 | > 2(k − |V 1 |): return N
		result.Answer = NT_V12_TOO_BIG
		return
	}

	// 12. return
	return
}

func (self *ntDecomposition) Str() string {
	return fmt.Sprintf(
		"NT Decomposition\nV1: %v\nV12: %v\nV0: %v\nk: %v\nans: %v",
		self.V1,
		self.V12,
		self.V0,
		self.K,
		self.Answer)
}
