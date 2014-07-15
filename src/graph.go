package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices           map[int]*Vertex
	Edges              Edges
	currentVertexIndex int
}

func (self *Graph) hasVertex(v int) bool {
	_, res := self.Vertices[v]
	return res
}

func (self *Graph) hasEdge(a, b int) bool {
	for _, v := range self.Edges {
		if v.from == self.Vertices[a] && v.to == self.Vertices[b] || v.from == self.Vertices[b] && v.to == self.Vertices[a] {
			return true
		}
	}

	return false
}

func (self *Graph) getCoveredEdgePositions(v int) []int {
	result := make([]int, 0)
	for index, edge := range self.Edges {
		if edge.IsCoveredBy(self.Vertices[v]) {
			result = append(result, index)
		}
	}

	return result
}

func (g *Graph) AddVertex() {
	Debug("Adding %v", g.currentVertexIndex)
	vertex := g.generateVertex()
	g.Vertices[vertex.id] = &vertex
}

func (self *Graph) RemoveVertex(v int) error {
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

func (self *Graph) AddEdge(a, b int) error {
	if !self.hasVertex(a) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", a))
	}

	if !self.hasVertex(b) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", b))
	}

	if self.Vertices[a] == self.Vertices[b] {
		return errors.New(fmt.Sprintf("Connect two separate vertices."))
	}

	if self.hasEdge(a, b) {
		return errors.New(fmt.Sprintf("An edge between %v and %v already exists.", a, b))
	}

	self.Edges = append(self.Edges, Edge{self.Vertices[a], self.Vertices[b]})
	self.Vertices[a].degree += 1
	self.Vertices[b].degree += 1
	return nil
}

func (self *Graph) IsVertexCover(vertices ...int) bool {
	isCovered := make(map[Edge]bool)
	for _, edge := range self.Edges {
		isCovered[edge] = false
	}

	for _, vertex := range vertices {
		for _, edge := range self.Edges {
			if edge.IsCoveredBy(self.Vertices[vertex]) {
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

func (self *Graph) Degree(v int) (int, error) {
	result := 0
	if !self.hasVertex(v) {
		return -1, errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	for _, edge := range self.Edges {
		if edge.IsCoveredBy(self.Vertices[v]) {
			result++
		}
	}

	return result, nil
}

func MkGraph() *Graph {
	g := new(Graph)
	g.Vertices = make(map[int]*Vertex)
	g.Edges = make(Edges, 0)
	return g
}
