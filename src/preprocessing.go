package graph

import (
	"errors"
	"fmt"
)

// TODO: copy the graph instead of mutating.
func (self *Graph) forAllVerticesOfDegree(degree int, action func(Vertex) error) (result error) {
	self.ForAllVertices(func(vertex Vertex, index int, done chan<- bool) {
		vDegree, err := self.Degree(vertex)
		if nil != err {
			result = errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", vertex))
			done <- true
		}

		if vDegree == degree {
			err = action(vertex)
			if nil != err {
				result = err
				done <- true
			}
		}
	})

	return result
}

func (self *Graph) getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree int) NeighborMap {
	result := make(NeighborMap, len(self.Vertices))
	self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		self.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
			result.AddNeighborOfVertex(v, getOtherVertex(v, edge))
		})
		// neighbors := self.getNeighbors(v)
		// length := len(neighbors)
		// for i := 0; i < length; i++ {
		// 	for j := 0; j < length; j++ {
		// 		if i == j {
		// 			continue
		// 		}
		//
		// 		n1, n2 := neighbors[i], neighbors[j]
		//
		//
		// 		if self.hasEdge(n1, n2) {
		// 			result.AddNeighborOfVertex(v, n1)
		// 			result.AddNeighborOfVertex(v, n2)
		// 			break
		// 		}
		// 	}
		// }

		return nil
	})

	return result
}

func (self *Graph) getVerticesOfDegreeWithOnlyDisjointNeighbors(degree int) NeighborMap {
	Debug("===== GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	result := make(NeighborMap, len(self.Vertices))
	self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		neighbors := self.getNeighbors(v)
		length := len(neighbors)
		Debug("OPERATION FOR %v (neighbors: %v)", v, neighbors)
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

		return nil
	})

	Debug("%v", result)
	Debug("===== END GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	return result
}
func (self *Graph) removeVerticesOfDegree(degree int) error {
	return self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		return self.RemoveVertex(v)
	})
}

func (self *Graph) removeAllVerticesAccordingToMap(v NeighborMap) {
	performRemoval := removeOnce(self, make(coverageMap))
	for center, neighbors := range v {
		if nil == neighbors {
			continue
		}

		performRemoval(MkVertex(center))
		for _, neighbor := range neighbors {
			performRemoval(neighbor)
		}
	}
}

func (self *Graph) removeVertivesOfDegreeWithOnlyAdjacentNeighbors(degree int) {
	self.removeAllVerticesAccordingToMap(self.getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree))
}

// func (self *Graph) rewireEdge(from, to, newAnchor Vertex) {
// 	fromIndex := from.toInt()
// 	toIndex := to.toInt()
// 	newAnchorIndex := newAnchor.toInt()

// 	edge := self.getEdgeByCoordinates(fromIndex, toIndex)
// 	if nil == edge {
// 		return
// 	}

// 	edge.from = newAnchor
// 	self.neighbors[newAnchorIndex][toIndex] = edge
// 	self.neighbors[toIndex][newAnchorIndex] = edge
// 	self.degrees[newAnchorIndex]++
// 	self.degrees[fromIndex]--
// }

func (self *Graph) contractEdges(contractionMap NeighborMap) {
	// toRemove is probably redundant given the circumstances under which this will be called.
	// TODO Reason about this redundancy and introduce changes.
	toRemove := make(Neighbors, 0, self.NVertices())
	contractionMap.ForAll(func(vertex Vertex, neighbors Neighbors, done chan<- bool) {
		// TODO @refactor rewire edges from neighbor->distantNeighbor
		// When refactoring, keep in mind the cases when two neighbors of one
		// vertex are connected.
		// TODO @optimize
		for _, neighbor := range neighbors {
			distantNeighbors := self.getNeighbors(neighbor)
			Debug("Neighbor: %v", neighbor)
			for _, distantNeighbor := range distantNeighbors {
				Debug("Distante Neighbor: %v", distantNeighbor)
				// self.rewireEdge(neighbor, distantNeighbor, vertex)
				// self.neighbors[vertex][distantNeighbor] = self.neighbors[neighbor][distantNeighbor]
				// self.neighbors[distantNeighbor][vertex] = self.neighbors[distantNeighbor][neighbor]
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
	err := g.removeVerticesOfDegree(0)
	if nil != err {
		return err
	}

	// 2. Vertices interconnected with only themselves are useless - remove them.
	err = g.removeVerticesOfDegree(1)
	if nil != err {
		return err
	}

	err = g.removeVerticesOfDegree(0)
	if nil != err {
		return err
	}

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
