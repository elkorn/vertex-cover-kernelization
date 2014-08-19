package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	Vertices           Vertices
	Edges              Edges
	degrees            []int
	currentVertexIndex int
	isVertexDeleted    []bool
}

func (self *Graph) hasVertex(v Vertex) bool {
	index := self.indexOfVertex(v)
	return index != -1 && !self.isVertexDeleted[index]
}

func (self *Graph) indexOfVertex(v Vertex) int {
	for i, b := range self.Vertices {
		if b == v {
			return i
		}
	}

	return -1
}

func (self *Graph) ForAllEdges(fn func(*Edge, int, chan<- bool)) {
	done := make(chan bool, 1)
	for idx, edge := range self.Edges {
		if edge.isDeleted {
			continue
		}

		fn(edge, idx, done)

		select {
		case <-done:
			return
		default:
		}

	}
}

func (self *Graph) hasEdge(a, b Vertex) bool {
	result := false
	self.ForAllEdges(func(edge *Edge, i int, done chan<- bool) {
		if edge.from == a && edge.to == b || edge.from == b && edge.to == a {
			result = true
			done <- true
		}
	})

	return result
}

func (self *Graph) getCoveredEdgePositions(v Vertex) []int {
	result := make([]int, 0)
	self.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if edge.IsCoveredBy(v) {
			result = append(result, index)
		}
	})

	return result
}

func (g *Graph) addVertex() error {
	g.currentVertexIndex++
	Debug("Adding %v", g.currentVertexIndex)
	g.Vertices = append(g.Vertices, Vertex(g.currentVertexIndex))
	g.isVertexDeleted = append(g.isVertexDeleted, false)
	g.degrees = append(g.degrees, 0)
	return nil
}

func (self *Graph) RemoveVertex(v Vertex) error {
	index := self.indexOfVertex(v)
	if index == -1 || self.isVertexDeleted[index] {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	self.isVertexDeleted[index] = true
	positions := self.getCoveredEdgePositions(v)
	for i := len(positions) - 1; i >= 0; i-- {
		self.degrees[self.Edges[positions[i]].from.toInt()] -= 1
		self.degrees[self.Edges[positions[i]].to.toInt()] -= 1
		self.Edges[positions[i]].isDeleted = true
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

	self.degrees[a.toInt()] += 1
	self.degrees[b.toInt()] += 1
	return nil
}

func (self *Graph) IsVertexCover(vertices ...Vertex) bool {
	// TODO refactor to use a number instead of a map !!!
	n := len(self.Edges)
	amountCovered := 0
	isCovered := make([]bool, n)
	for _, vertex := range vertices {
		self.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
			if edge.IsCoveredBy(vertex) {
				if !isCovered[index] {
					amountCovered++
					isCovered[index] = true
				}
			}
		})
	}

	Debug("Coverage map for %v: %v", vertices, isCovered)
	return amountCovered == n
}

func (self *Graph) Degree(v Vertex) (int, error) {
	if !self.hasVertex(v) {
		return -1, errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	return self.degrees[v.toInt()], nil
}

func MkGraph(vertices int) *Graph {
	g := new(Graph)
	g.Vertices = make(Vertices, vertices)
	for i := 0; i < vertices; i++ {
		g.Vertices[i] = MkVertex(i)
	}

	g.currentVertexIndex = vertices
	g.Edges = make(Edges, 0)
	g.degrees = make([]int, vertices)
	g.isVertexDeleted = make([]bool, vertices)
	return g
}
