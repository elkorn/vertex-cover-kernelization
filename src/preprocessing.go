package graph

import (
	"errors"
	"fmt"
)

func (self *Graph) forAllVerticesOfDegree(degree int, action func(Vertex)) {
	// Create an immutable view of vertices with given degree.
	vertices := make(Vertices, 0, self.NVertices())

	self.ForAllVertices(func(vertex Vertex, done chan<- bool) {
		vDegree, err := self.Degree(vertex)
		if nil != err {
			panic(errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", vertex)))
		}

		if vDegree == degree {
			vertices = append(vertices, vertex)
		}
	})

	Debug("Vertices of degree %v: %v", degree, vertices)

	for _, vertex := range vertices {
		action(vertex)
	}
}

func (self *Graph) getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree int) NeighborMap {
	result := MkNeighborMap(self.currentVertexIndex)
	self.forAllVerticesOfDegree(degree, func(v Vertex) {
		self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
			result.AddNeighborOfVertex(v, getOtherVertex(v, edge))
		})
	})

	return result
}

func (self *Graph) getVerticesOfDegreeWithOnlyDisjointNeighbors(degree int) NeighborMap {
	Debug("===== GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
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
			result[v.toInt()] = potentiallyToBeAdded
		}
	})

	Debug("%v", result)
	Debug("===== END GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	return result
}
func (self *Graph) removeVerticesOfDegree(degree int) {
	self.forAllVerticesOfDegree(degree, func(v Vertex) {
		self.RemoveVertex(v)
	})
}

func (self *Graph) removeAllVerticesAccordingToMap(v NeighborMap) {
	performRemoval := removeOnce(self, make(coverageMap))
	for center, neighbors := range v {
		if nil == neighbors || len(neighbors) == 0 {
			continue
		}

		performRemoval(MkVertex(center))
		for _, neighbor := range neighbors {
			performRemoval(neighbor)
		}
	}
}

func (self *Graph) removeVertivesOfDegreeWithOnlyAdjacentNeighbors(degree int) {
	self.removeAllVerticesAccordingToMap(
		self.getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree))
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

func Preprocessing(g *Graph) error {
	// 1. Disjoint vertices cannot be in a vertex cover - remove them.
	g.removeVerticesOfDegree(0)
	// 2. Vertices interconnected with only themselves are useless - remove them.
	g.removeVerticesOfDegree(1)
	g.removeVerticesOfDegree(0)

	// 3. Remove vertices with degree 2 which have connected neighbors.
	// Then, remove nodes whose degree has dropped to 0.
	g.removeVertivesOfDegreeWithOnlyAdjacentNeighbors(2)
	g.removeVerticesOfDegree(0)

	// 4. Contract the edges between vertices of degree 2 and their neighbors if they are not connected.
	// Repeat this step until all such vertices are eliminated.
	for contractable := g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2); len(contractable) > 0; contractable = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2) {
		g.contractEdges(contractable)
	}

	return nil
}
