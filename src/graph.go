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
	result.Vertices = make(Vertices, len(self.Vertices))
	copy(result.Vertices, self.Vertices)
	result.isVertexDeleted = make([]bool, len(self.isVertexDeleted))
	copy(result.isVertexDeleted, self.isVertexDeleted)
	result.Edges = make(Edges, 0, cap(self.Edges))
	result.neighbors = make([][]*Edge, len(self.neighbors))
	result.degrees = make([]int, len(self.degrees))

	for x := range self.neighbors {
		result.neighbors[x] = make([]*Edge, len(self.neighbors[x]))
	}

	for i, edge := range self.Edges {
		Debug("%v", edge)
		if nil == edge {
			continue
		}
		result.AddEdge(edge.from, edge.to)
		result.Edges[i].isDeleted = edge.isDeleted
	}

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
	self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
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

func (self *Graph) ForAllVertices(fn func(Vertex, chan<- bool)) {
	done := make(chan bool, 1)
	for _, vertex := range self.Vertices {
		if self.isVertexDeleted[vertex.toInt()] {
			continue
		}

		fn(vertex, done)

		select {
		case <-done:
			return
		default:
		}
	}
}

func (self *Graph) ForAllEdges(fn func(*Edge, chan<- bool)) {
	done := make(chan bool, 1)
	for _, edge := range self.Edges {
		if nil == edge || edge.isDeleted {
			continue
		}

		fn(edge, done)

		select {
		case <-done:
			return
		default:
		}
	}
}

func (self *Graph) ForAllNeighbors(v Vertex, fn func(*Edge, chan<- bool)) {
	done := make(chan bool, 1)
	for _, edge := range self.getNeighborEdges(v) {
		if nil == edge || edge.isDeleted {
			continue
		}

		fn(edge, done)
		select {
		case <-done:
			return
		default:
		}
	}
}

func (self *Graph) hasEdge(a, b Vertex) bool {
	edge := self.getEdgeByCoordinates(a.toInt(), b.toInt())
	return edge != nil && !edge.isDeleted
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

	self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
		self.RemoveEdge(edge.from, edge.to)
	})

	self.removeVertex(v)
	return nil
}

func (self *Graph) RestoreVertex(v Vertex) {
	vi := v.toInt()
	if !self.isVertexDeleted[vi] {
		return
	}

	for _, edge := range self.neighbors[vi] {
		if nil == edge {
			continue
		}

		edge.isDeleted = false
		self.neighbors[getOtherVertex(v, edge)][vi].isDeleted = false
		self.degrees[edge.from.toInt()]++
		self.degrees[edge.to.toInt()]++
		self.numberOfEdges++
	}
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
	ai := a.toInt()
	bi := b.toInt()
	self.degrees[ai]++
	self.degrees[bi]++
	self.numberOfEdges++
	return nil
}

func (self *Graph) RemoveEdge(from, to Vertex) {
	fi, ti := from.toInt(), to.toInt()
	edge := self.getEdgeByCoordinates(fi, ti)
	self.degrees[fi] -= 1
	self.degrees[ti] -= 1
	edge.isDeleted = true
	self.numberOfEdges--
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
		self.ForAllEdges(func(edge *Edge, done chan<- bool) {
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

func mkGraph(vertices, capacity int) *Graph {
	g := new(Graph)
	g.Vertices = make(Vertices, vertices, capacity)
	for i := 0; i < vertices; i++ {
		g.Vertices[i] = MkVertex(i)
	}

	g.currentVertexIndex = vertices
	// NOTE: Duplicate edges are not allowed, thus the maximum number of edges in the graph is V^2
	// (when every vertex is connected to each other)
	g.Edges = make(Edges, 0, capacity*capacity)
	g.degrees = make([]int, vertices, capacity)
	g.neighbors = make([][]*Edge, capacity)
	for y := range g.neighbors {
		g.neighbors[y] = make([]*Edge, capacity)
	}

	g.isVertexDeleted = make([]bool, vertices, capacity)
	g.numberOfVertices = vertices
	return g
}

func MkGraph(vertices int) *Graph {
	return mkGraph(vertices, vertices)
}

func MkGraphWithCapacity(vertices, capacity int) *Graph {
	return mkGraph(vertices, capacity)
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
