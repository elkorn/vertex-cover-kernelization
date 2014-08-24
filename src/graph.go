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
	neighbors          [][]*Edge
	numberOfVertices   int
	numberOfEdges      int
}

func (self *Graph) Copy() *Graph {
	result := &Graph{
		currentVertexIndex: self.currentVertexIndex,
		numberOfVertices:   self.numberOfVertices,
		numberOfEdges:      self.numberOfEdges,
	}
	copy(result.Vertices, self.Vertices)
	copy(result.Edges, self.Edges)
	copy(result.isVertexDeleted, self.isVertexDeleted)
	copy(result.degrees, self.degrees)

	return result
}

func (self *Graph) hasVertex(v Vertex) bool {
	return v.toInt() < self.currentVertexIndex && !self.isVertexDeleted[v.toInt()]
}

func (self *Graph) getNeighborEdges(v Vertex) []*Edge {
	return self.neighbors[v.toInt()]
}

func (self *Graph) getNeighbors(v Vertex) Neighbors {
	result := make(Neighbors, 0, len(self.getNeighborEdges(v)))
	self.ForAllNeighbors(v, func(edge *Edge, idx int, done chan<- bool) {
		Debug("Found neighbor edge of %v: %v", v, edge)

		result = result.appendIfNotContains(getOtherVertex(v, edge))
	})

	return result
}

func (self *Graph) NVertices() int {
	return self.numberOfVertices
}

func (self *Graph) NEdges() int {
	return self.numberOfEdges
}

func (self *Graph) ForAllVertices(fn func(Vertex, int, chan<- bool)) {
	done := make(chan bool, 1)
	for idx, vertex := range self.Vertices {
		if self.isVertexDeleted[vertex.toInt()] {
			continue
		}

		fn(vertex, idx, done)

		select {
		case <-done:
			return
		default:
		}
	}
}

func (self *Graph) ForAllEdges(fn func(*Edge, int, chan<- bool)) {
	done := make(chan bool, 1)
	for idx, edge := range self.Edges {
		if nil == edge || edge.isDeleted {
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

func (self *Graph) ForAllNeighbors(v Vertex, fn func(*Edge, int, chan<- bool)) {
	done := make(chan bool, 1)
	for idx, edge := range self.getNeighborEdges(v) {
		if nil == edge || edge.isDeleted {
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
	g.numberOfVertices++
	for y, neighbor := range g.neighbors {
		Debug("Before append : %v", g.neighbors[y])
		g.neighbors[y] = append(neighbor, nil)
		Debug("After append : %v", g.neighbors[y])
	}

	g.neighbors = append(g.neighbors, make([]*Edge, g.currentVertexIndex))
	return nil
}

func (self *Graph) RemoveVertex(v Vertex) error {
	if !self.hasVertex(v) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", v))
	}

	positions := self.getCoveredEdgePositions(v)
	for i := len(positions) - 1; i >= 0; i-- {
		self.degrees[self.Edges[positions[i]].from.toInt()] -= 1
		self.degrees[self.Edges[positions[i]].to.toInt()] -= 1
		self.Edges[positions[i]].isDeleted = true
		self.numberOfEdges--
	}

	self.removeVertex(v)
	return nil
}

func (self *Graph) removeVertex(v Vertex) {
	self.isVertexDeleted[v.toInt()] = true
	self.numberOfVertices--
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

	edge := MkEdge(a, b)
	self.Edges = append(self.Edges, edge)
	self.neighbors[a.toInt()][b.toInt()] = edge
	self.neighbors[b.toInt()][a.toInt()] = edge

	Debug("Added neighbor %v->%v: %v", a, b, *edge)
	self.degrees[a.toInt()] += 1
	self.degrees[b.toInt()] += 1
	self.numberOfEdges++
	return nil
}

func (self *Graph) IsVertexCover(vertices ...Vertex) bool {
	n := self.NEdges()
	amountCovered := 0
	isCovered := make([][]bool, self.currentVertexIndex)
	for i := range isCovered {
		isCovered[i] = make([]bool, self.currentVertexIndex)
	}

	for _, vertex := range vertices {
		Debug("Checking %v for coverage", vertex)
		self.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
			if edge.IsCoveredBy(vertex) {
				if !isCovered[edge.from.toInt()][edge.to.toInt()] {
					Debug("Edge %v -> Covered", *edge)
					amountCovered++
					isCovered[edge.from.toInt()][edge.to.toInt()] = true
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
	// NOTE: Duplicate edges are not allowed, thus the maximum number of edges in the graph is V^2
	// (when every vertex is connected to each other)
	g.Edges = make(Edges, vertices*vertices)
	g.degrees = make([]int, vertices)
	g.neighbors = make([][]*Edge, vertices)
	for y := range g.neighbors {
		g.neighbors[y] = make([]*Edge, vertices)
	}

	g.isVertexDeleted = make([]bool, vertices)
	g.numberOfVertices = vertices
	return g
}

func MkGraphRememberingDeletedVertices(vertices int, deletedReference []bool) *Graph {
	g := MkGraph(vertices)
	for i, isDeleted := range deletedReference {
		if isDeleted {
			g.removeVertex(MkVertex(i))
		}
	}

	return g
}
