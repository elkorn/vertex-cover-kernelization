package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices map[Vertex]bool
	Edges    Edges
}

func (self *Graph) hasVertex(v Vertex) bool {
	return self.Vertices[v]
}

func (self *Graph) hasEdge(a, b Vertex) bool {
	for _, v := range self.Edges {
		if v.from == a && v.to == b || v.from == b && v.to == a {
			return true
		}
	}

	return false
}

func (self *Graph) getCoveredEdgePositions(v Vertex) []int {
	result := make([]int, 0)
	for index, edge := range self.Edges {
		if edge.IsCoveredBy(v) {
			result = append(result, index)
		}
	}

	return result
}

func (g *Graph) AddVertex(v Vertex) error {
	_, exists := g.Vertices[v]
	if exists {
		return errors.New(fmt.Sprintf("Vertex %v already in the set", v))
	}

	g.Vertices[v] = true
	return nil
}

func (self *Graph) RemoveVertex(v Vertex) error {
	if !self.hasVertex(v) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	delete(self.Vertices, v)
	positions := self.getCoveredEdgePositions(v)
	for i := len(positions) - 1; i >= 0; i-- {
		self.Edges = removeAt(self.Edges, positions[i])
	}

	return nil
}

func (self *Graph) AddEdge(a, b Vertex) error {
	if a == b {
		return errors.New(fmt.Sprintf("Connect two separate vertices."))
	}

	if !self.hasVertex(a) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", a))
	}

	if !self.hasVertex(b) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", b))
	}

	if self.hasEdge(a, b) {
		return errors.New(fmt.Sprintf("An edge between %v and %v already exists.", a, b))
	}

	self.Edges = append(self.Edges, Edge{a, b})
	return nil
}

func (self *Graph) IsVertexCover(vertices ...Vertex) bool {
	isCovered := make(map[Edge]bool)
	for _, edge := range self.Edges {
		isCovered[edge] = false
	}

	for _, vertex := range vertices {
		for _, edge := range self.Edges {
			if edge.IsCoveredBy(vertex) {
				isCovered[edge] = true
			}
		}
	}

	Debug("Coverage map for %v: %v", vertices, isCovered)
	for _, v := range isCovered {
		if v == false {
			return false
		}
	}

	return true
}

func (self *Graph) Degree(v Vertex) (int, error) {
	result := 0
	if !self.hasVertex(v) {
		return -1, errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	for _, edge := range self.Edges {
		if edge.IsCoveredBy(v) {
			result++
		}
	}

	return result, nil
}

func MkGraph() *Graph {
	g := new(Graph)
	g.Vertices = make(map[Vertex]bool)
	g.Edges = make(Edges, 0)
	return g
}
