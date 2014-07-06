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

func (self *Graph) removeVerticesOfDegree(degree int) error {
	return self.forAllVerticesOfDegree(degree, func(v Vertex) error {
		return self.RemoveVertex(v)
	})
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
					Debug(fmt.Sprintf("For %v", v))
					if result[v] == nil {
						result[v] = Neighbors{n1, n2}
					} else {
						if !contains(result[v], n1) {
							Debug(fmt.Sprintf("%v contains %v == %v", result[v], n1, contains(result[v], n1)))
							result[v] = append(result[v], n1)
						}

						if !contains(result[v], n2) {
							Debug(fmt.Sprintf("%v contains %v == %v", result[v], n2, contains(result[v], n2)))
							result[v] = append(result[v], n2)
						}
					}

					break
				}
			}
		}
		return nil
	})
	return result
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
	// TODO
	// 4. Contract the edges between vertices of degree 2 and their neighbors if they are not connected.
	return nil
}