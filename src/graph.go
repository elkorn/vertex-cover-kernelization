package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices           []Vertex
	Edges              Edges
	degrees            map[Vertex]int
	currentVertexIndex int
}

func (self *Graph) hasVertex(v Vertex) bool {
	for _, b := range self.Vertices {
		if b == v {
			return true
		}
	}

	return false
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

func (g *Graph) AddVertex() error {
	g.currentVertexIndex++
	Debug("Adding %v", g.currentVertexIndex)
	g.Vertices = append(g.Vertices, Vertex(g.currentVertexIndex))
	return nil
}

func (self *Graph) RemoveVertex(v Vertex) error {
	if !self.hasVertex(v) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	self.Vertices = append(self.Vertices[0:v.toInt()], self.Vertices[v.toInt()+1:self.currentVertexIndex]...)
	positions := self.getCoveredEdgePositions(v)
	for i := len(positions) - 1; i >= 0; i-- {
		self.degrees[self.Edges[positions[i]].from] -= 1
		self.degrees[self.Edges[positions[i]].to] -= 1
		self.Edges = removeAt(self.Edges, positions[i])
	}

	return nil
}

func (self *Graph) AddEdge(a, b Vertex) error {
	if a == b {
		return errors.New(fmt.Sprintf("Cannot connect vertex %v with itself.", a))
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

	self.Edges = append(self.Edges, MkEdge(a, b))

	self.degrees[a] += 1
	self.degrees[b] += 1
	return nil
}

func (self *Graph) IsVertexCover(vertices ...Vertex) bool {
	isCovered := make(map[*Edge]bool)
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
	if !self.hasVertex(v) {
		return -1, errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	return self.degrees[v], nil
}

func MkGraph() *Graph {
	g := new(Graph)
	g.Vertices = make([]Vertex, 0)
	g.Edges = make(Edges, 0)
	g.degrees = make(map[Vertex]int)
	return g
}
