package graph

import (
	"github.com/deckarep/golang-set"
)

type fold struct {
	root        Vertex
	replacement Vertex
	neighbors   Neighbors
}

func (self *Graph) forAllVerticesOfDegree(degree int, action func(Vertex)) {
	// Create an immutable view of vertices with given degree.
	vertices := make(Vertices, 0, self.NVertices())

	self.ForAllVertices(func(vertex Vertex, done chan<- bool) {
		if self.Degree(vertex) == degree {
			vertices = append(vertices, vertex)
		}
	})

	for _, vertex := range vertices {
		action(vertex)
	}
}

func (self *Graph) getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree int) (NeighborMap, bool) {
	result := MkNeighborMap(self.currentVertexIndex)
	hasNeighbors := false
	self.forAllVerticesOfDegree(degree, func(v Vertex) {
		neighbors := self.getNeighbors(v)
		for _, n1 := range neighbors {
			for _, n2 := range neighbors {
				if n1 == n2 {
					continue
				}

				if !self.hasEdge(n1, n2) {
					return
				}
			}
		}

		result[v.toInt()] = neighbors
		hasNeighbors = true
	})

	return result, hasNeighbors
}

func (self *Graph) getVerticesOfDegreeWithOnlyDisjointNeighbors(degree int) (NeighborMap, bool) {
	Debug("===== GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	hasNeighbors := false
	result := MkNeighborMap(self.currentVertexIndex)
	self.forAllVerticesOfDegree(degree, func(v Vertex) {
		neighbors := self.getNeighbors(v)
		length := len(neighbors)
		hasOnlyDisjoint := true
		potentiallyToBeAdded := make(Neighbors, 0, length)
		if length == 1 {
			potentiallyToBeAdded = potentiallyToBeAdded.appendIfNotContains(neighbors[0])
		} else {
			for i := 0; i < length; i++ {
				n1 := neighbors[i]
				for j := 0; j < length && hasOnlyDisjoint; j++ {
					if i == j {
						continue
					}

					n2 := neighbors[j]

					if self.hasEdge(n1, n2) {
						Debug("%v and %v are NOT disjoint", n1, n2)
						hasOnlyDisjoint = false
						break
					} else {
						Debug("%v and %v are disjoint", n1, n2)
						potentiallyToBeAdded = potentiallyToBeAdded.appendIfNotContains(n1)
						potentiallyToBeAdded = potentiallyToBeAdded.appendIfNotContains(n2)
					}
				}
			}
		}

		if hasOnlyDisjoint {
			Debug("For %v: adding %v", v, potentiallyToBeAdded)
			if len(potentiallyToBeAdded) > 0 {
				result[v.toInt()] = potentiallyToBeAdded
				hasNeighbors = true
			}
		}
	})

	Debug("%v", result)
	Debug("===== END GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	return result, hasNeighbors
}
func (self *Graph) removeVerticesOfDegree(degree int) int {
	removed := 0
	self.forAllVerticesOfDegree(degree, func(v Vertex) {
		self.RemoveVertex(v)
		removed++
	})

	return removed
}

func (self *Graph) removeAllVerticesAccordingToMap(v NeighborMap) int {
	removed := 0
	performRemoval := removeOnce(self, make(coverageMap))
	for center, neighbors := range v {
		if nil == neighbors || len(neighbors) == 0 {
			continue
		}

		performRemoval(MkVertex(center))

		for _, neighbor := range neighbors {
			if performRemoval(neighbor) {
				removed++
			}
		}
	}

	return removed
}

func (self *Graph) removeVertivesOfDegreeWithOnlyAdjacentNeighbors(degree int) int {
	removed := 0
	neighborMap, mayRemove := self.getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree)

	for mayRemove {
		removed += self.removeAllVerticesAccordingToMap(neighborMap)
		neighborMap, mayRemove = self.getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree)
	}

	return removed
}

func (self *Graph) contractEdges(contractionMap NeighborMap) {
	toRemove := make(Neighbors, 0, self.NVertices())
	contractionMap.ForAll(func(vertex Vertex, neighbors Neighbors, done chan<- bool) {
		for _, neighbor := range neighbors {
			distantNeighbors := self.getNeighbors(neighbor)
			Debug("Neighbor: %v", neighbor)
			for _, distantNeighbor := range distantNeighbors {
				Debug("Distante Neighbor: %v", distantNeighbor)
				self.AddEdge(vertex, distantNeighbor)
			}

			toRemove = toRemove.appendIfNotContains(neighbor)
		}
	})

	for _, neighbor := range toRemove {
		self.RemoveVertex(neighbor)
	}

}

func (g *Graph) fold(u Vertex) *fold {
	neighbors := g.getNeighbors(u)
	for _, n1 := range neighbors {
		for _, n2 := range neighbors {
			if n1 == n2 {
				continue
			}

			if g.hasEdge(n1, n2) {
				return nil
			}
		}
	}

	Debug("Adding vertex")
	g.addVertex()
	uPrime := Vertex(g.currentVertexIndex)
	theFold := &fold{
		root:        u,
		replacement: uPrime,
		neighbors:   neighbors,
	}

	Debug("Creating fold %v (%v) -> %v", theFold.root, theFold.neighbors, theFold.replacement)

	for _, n := range neighbors {
		g.ForAllNeighbors(n, func(edge *Edge, done chan<- bool) {
			v := getOtherVertex(n, edge)
			g.AddEdge(uPrime, v)
		})
	}

	g.RemoveVertex(theFold.root)
	for _, n := range neighbors {
		g.RemoveVertex(n)
	}

	return theFold
}

func preprocessing4(g *Graph) mapset.Set {
	folds := mapset.NewSet()
	// bound := Vertex(g.currentVertexIndex)
	g.forAllVerticesOfDegree(2, func(u Vertex) {
		theFold := g.fold(u)

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

// func unfoldIter(theFold *fold, vc, folds mapset.Set, bound Vertex) mapset.Set {
// 	Debug("VC: %v", vc)
// 	if vc.Contains(theFold.replacement) {
// 		Debug("vc contains replacement for %v", theFold.neighbors)
// 		for _, neighbor := range theFold.neighbors {
// 			Debug("Checking neighbor %v...", neighbor)
// 			vc.Add(neighbor)
// 			if neighbor > bound {
// 				Debug("It's syntetic.")
// 				for fInter := range folds.Iter() {
// 					f := fInter.(*fold)
// 					Debug("checking fold with syntetic: %v", f.replacement)
// 					if f.replacement == neighbor {
// 						Debug("replacement %v == neighbor %v", f.replacement, neighbor)
// 						unfoldIter(f, vc, folds, bound)
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return vc
// }

// func unfold(vc, folds mapset.Set, bound Vertex) mapset.Set {
// 	for fInter := range folds.Iter() {
// 		unfoldIter(fInter.(*fold), vc, folds, bound)
// 	}

// 	return vc
// }

func Preprocessing(g *Graph) (int, mapset.Set) {
	parameterReduction := 0
	// 1. Disjoint vertices cannot be in a vertex cover - remove them.
	g.removeVerticesOfDegree(0)
	// 2. Vertices interconnected with only themselves are useless - remove them.
	parameterReduction += g.removeVerticesOfDegree(1)
	g.removeVerticesOfDegree(0)
	Debug("Removed %v pendant vertices.", parameterReduction)

	// 3. Remove vertices with degree 2 which have connected neighbors.
	// Then, remove nodes whose degree has dropped to 0.
	//
	red := g.removeVertivesOfDegreeWithOnlyAdjacentNeighbors(2)
	Debug("Removed %v vertices of deg. 2 with adj. neighbors.", red)
	parameterReduction += red
	g.removeVerticesOfDegree(0)

	// 4. Contract the edges between vertices of degree 2 and their neighbors if they are not connected.
	// Repeat this step until all such vertices are eliminated.

	folds := preprocessing4(g)

	return parameterReduction, folds
}

func computeUnfoldedVertexCoverSize(folds, vc mapset.Set) int {
	size := vc.Cardinality()
	for foldInter := range folds.Iter() {
		fold := foldInter.(*fold)
		if vc.Contains(fold.replacement) {
			size += 2
		}
	}

	return size
}
