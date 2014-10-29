package preprocessing

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type fold struct {
	root        graph.Vertex
	replacement graph.Vertex
	neighbors   graph.Neighbors
}

type coverageMap []bool

func getVerticesOfDegreeWithOnlyAdjacentNeighbors(self *graph.Graph, degree int) (graph.NeighborMap, bool) {
	result := graph.MkNeighborMap(self.CurrentVertexIndex)
	hasNeighbors := false
	self.ForAllVerticesOfDegree(degree, func(v graph.Vertex) {
		neighbors := self.GetNeighbors(v)
		for _, n1 := range neighbors {
			for _, n2 := range neighbors {
				if n1 == n2 {
					continue
				}

				if !self.HasEdge(n1, n2) {
					return
				}
			}
		}

		result[v.ToInt()] = neighbors
		hasNeighbors = true
	})

	return result, hasNeighbors
}

func getVerticesOfDegreeWithOnlyDisjointNeighbors(self *graph.Graph, degree int) (graph.NeighborMap, bool) {
	utility.Debug("===== GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	hasNeighbors := false
	result := graph.MkNeighborMap(self.CurrentVertexIndex)
	self.ForAllVerticesOfDegree(degree, func(v graph.Vertex) {
		neighbors := self.GetNeighbors(v)
		length := len(neighbors)
		hasOnlyDisjoint := true
		potentiallyToBeAdded := make(graph.Neighbors, 0, length)
		if length == 1 {
			potentiallyToBeAdded = potentiallyToBeAdded.AppendIfNotContains(neighbors[0])
		} else {
			for i := 0; i < length; i++ {
				n1 := neighbors[i]
				for j := 0; j < length && hasOnlyDisjoint; j++ {
					if i == j {
						continue
					}

					n2 := neighbors[j]

					if self.HasEdge(n1, n2) {
						utility.Debug("%v and %v are NOT disjoint", n1, n2)
						hasOnlyDisjoint = false
						break
					} else {
						utility.Debug("%v and %v are disjoint", n1, n2)
						potentiallyToBeAdded = potentiallyToBeAdded.AppendIfNotContains(n1)
						potentiallyToBeAdded = potentiallyToBeAdded.AppendIfNotContains(n2)
					}
				}
			}
		}

		if hasOnlyDisjoint {
			utility.Debug("For %v: adding %v", v, potentiallyToBeAdded)
			if len(potentiallyToBeAdded) > 0 {
				result[v.ToInt()] = potentiallyToBeAdded
				hasNeighbors = true
			}
		}
	})

	utility.Debug("%v", result)
	utility.Debug("===== END GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	return result, hasNeighbors
}

func removeOnce(g *graph.Graph, removed coverageMap) func(graph.Vertex) bool {
	return func(v graph.Vertex) bool {
		if !removed[v.ToInt()] {
			err := g.RemoveVertex(v)
			if nil != err {
				removed[v.ToInt()] = true
				return true
			}
		}

		return false
	}
}

func removeVerticesOfDegree(self *graph.Graph, degree int) int {
	removed := 0
	self.ForAllVerticesOfDegree(degree, func(v graph.Vertex) {
		self.RemoveVertex(v)
		removed++
	})

	return removed
}

func removeAllVerticesAccordingToMap(self *graph.Graph, v graph.NeighborMap) int {
	removed := 0
	performRemoval := removeOnce(self, make(coverageMap, self.CurrentVertexIndex))
	for center, neighbors := range v {
		if nil == neighbors || len(neighbors) == 0 {
			continue
		}

		performRemoval(graph.MkVertex(center))

		for _, neighbor := range neighbors {
			if performRemoval(neighbor) {
				removed++
			}
		}
	}

	return removed
}

func removeVertivesOfDegreeWithOnlyAdjacentNeighbors(self *graph.Graph, degree int) int {
	removed := 0
	neighborMap, mayRemove := getVerticesOfDegreeWithOnlyAdjacentNeighbors(self, degree)

	for mayRemove {
		removed += removeAllVerticesAccordingToMap(self, neighborMap)
		neighborMap, mayRemove = getVerticesOfDegreeWithOnlyAdjacentNeighbors(self, degree)
	}

	return removed
}

func contractEdges(self *graph.Graph, contractionMap graph.NeighborMap) {
	toRemove := make(graph.Neighbors, 0, self.NVertices())
	contractionMap.ForAll(func(vertex graph.Vertex, neighbors graph.Neighbors, done chan<- bool) {
		for _, neighbor := range neighbors {
			distantNeighbors := self.GetNeighbors(neighbor)
			utility.Debug("Neighbor: %v", neighbor)
			for _, distantNeighbor := range distantNeighbors {
				utility.Debug("Distante Neighbor: %v", distantNeighbor)
				self.AddEdge(vertex, distantNeighbor)
			}

			toRemove = toRemove.AppendIfNotContains(neighbor)
		}
	})

	for _, neighbor := range toRemove {
		self.RemoveVertex(neighbor)
	}

}

func foldVertex(g *graph.Graph, u graph.Vertex) *fold {
	neighbors := g.GetNeighbors(u)
	for _, n1 := range neighbors {
		for _, n2 := range neighbors {
			if n1 == n2 {
				continue
			}

			if g.HasEdge(n1, n2) {
				return nil
			}
		}
	}

	utility.Debug("Adding vertex")
	g.AddVertex()
	uPrime := graph.Vertex(g.CurrentVertexIndex)
	theFold := &fold{
		root:        u,
		replacement: uPrime,
		neighbors:   neighbors,
	}

	utility.Debug("Creating fold %v (%v) -> %v", theFold.root, theFold.neighbors, theFold.replacement)

	for _, n := range neighbors {
		g.ForAllNeighbors(n, func(edge *graph.Edge, done chan<- bool) {
			v := graph.GetOtherVertex(n, edge)
			g.AddEdge(uPrime, v)
		})
	}

	g.RemoveVertex(theFold.root)
	for _, n := range neighbors {
		g.RemoveVertex(n)
	}

	return theFold
}

func preprocessing4(g *graph.Graph) mapset.Set {
	folds := mapset.NewThreadUnsafeSet()
	// bound := graph.Vertex(g.CurrentVertexIndex)
	g.ForAllVerticesOfDegree(2, func(u graph.Vertex) {
		theFold := foldVertex(g, u)

		if nil != theFold {
			folds.Add(theFold)
		}
	})

	// if u′ is not in the opitmal cover of G′, then all its incident edges must
	// be covered by other vertices.
	// Therefore, v and w need not be included in an optimal vertex cover of G,
	// because {u,v}, {u,w} can be covered by u.
	// If u′ is included in the optimal vertex cover of G′, then at least some
	// of its incident edges must be covered by u′.
	// This implies that both v and w must be in the vertex cover.
	return folds
}

// func unfoldIter(theFold *fold, vc, folds mapset.Set, bound graph.Vertex) mapset.Set {
// 	utility.Debug("VC: %v", vc)
// 	if vc.Contains(theFold.replacement) {
// 		utility.Debug("vc contains replacement for %v", theFold.neighbors)
// 		for _, neighbor := range theFold.neighbors {
// 			utility.Debug("Checking neighbor %v...", neighbor)
// 			vc.Add(neighbor)
// 			if neighbor > bound {
// 				utility.Debug("It's syntetic.")
// 				for fInter := range folds.Iter() {
// 					f := fInter.(*fold)
// 					utility.Debug("checking fold with syntetic: %v", f.replacement)
// 					if f.replacement == neighbor {
// 						utility.Debug("replacement %v == neighbor %v", f.replacement, neighbor)
// 						unfoldIter(f, vc, folds, bound)
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return vc
// }

// func unfold(vc, folds mapset.Set, bound graph.Vertex) mapset.Set {
// 	for fInter := range folds.Iter() {
// 		unfoldIter(fInter.(*fold), vc, folds, bound)
// 	}

// 	return vc
// }

func Preprocessing(g *graph.Graph) (int, mapset.Set) {
	parameterReduction := 0
	// 1. Disjoint vertices cannot be in a vertex cover - remove them.
	removeVerticesOfDegree(g, 0)
	// 2. Vertices interconnected with only themselves are useless - remove them.
	removeVerticesOfDegree(g, 1)
	removeVerticesOfDegree(g, 0)

	// 3. Remove vertices with degree 2 which have connected neighbors.
	// Then, remove nodes whose degree has dropped to 0.
	red := removeVertivesOfDegreeWithOnlyAdjacentNeighbors(g, 2)
	utility.Debug("Removed %v vertices of deg. 2 with adj. neighbors.", red)
	parameterReduction += red
	removeVerticesOfDegree(g, 0)

	// 4. Contract the edges between vertices of degree 2 and their neighbors if they are not connected.
	// Repeat this step until all such vertices are eliminated.

	folds := preprocessing4(g)

	return parameterReduction, folds
}

func ComputeUnfoldedVertexCoverSize(folds, vc mapset.Set) int {
	size := vc.Cardinality()
	for foldInter := range folds.Iter() {
		fold := foldInter.(*fold)
		if vc.Contains(fold.replacement) {
			size += 2
		}
	}

	return size
}
