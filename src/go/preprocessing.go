package graph

import (
	"errors"
	"fmt"
)

// TODO: copy the graph instead of mutating.
func (self *Graph) forAllVerticesOfDegree(degree int, action func(Vertex) error) error {
	for vertex := range self.Vertices {
		vDegree, err := self.Degree(vertex)
		if nil != err {
			return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", vertex))
		}

		if vDegree == degree {
			err := action(vertex)
			if nil != err {
				return err
			}
		}
	}

	return nil
}

func (self *Graph) getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree int) NeighborMap {
	result := make(NeighborMap)
	self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		neighbors := self.getNeighbors(v)
		length := len(neighbors)
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if i == j {
					continue
				}

				n1, n2 := neighbors[i], neighbors[j]

				if self.hasEdge(n1, n2) {
					result.AddNeighborOfVertex(v, n1)
					result.AddNeighborOfVertex(v, n2)
					break
				}
			}
		}

		return nil
	})

	return result
}

func (self *Graph) getVerticesOfDegreeWithOnlyDisjointNeighbors(degree int) NeighborMap {
	Debug("===== GET DISJOINT NEIGHBORS OF DEGREE %v =====", degree)
	result := make(NeighborMap)
	self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		neighbors := self.getNeighbors(v)
		length := len(neighbors)
		Debug("OPERATION FOR %v (neighbors: %v)", v, neighbors)
		hasOnlyDisjoint := true
		potentiallyToBeAdded := Neighbors{}
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
			result[v] = potentiallyToBeAdded
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
		performRemoval(center)
		for _, neighbor := range neighbors {
			performRemoval(neighbor)
		}
	}
}

func (self *Graph) removeVertivesOfDegreeWithOnlyAdjacentNeighbors(degree int) {
	self.removeAllVerticesAccordingToMap(self.getVerticesOfDegreeWithOnlyAdjacentNeighbors(degree))
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
	// TODO
	return nil
}
