package graph

import (
	"errors"
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type Graph struct {
	Vertices                Vertices
	Edges                   Edges
	degrees                 []int
	CurrentVertexIndex      int
	IsVertexDeleted         []bool
	neighbors               [][]*Edge
	numberOfVertices        int
	numberOfEdges           int
	isRegular               bool
	needToComputeRegularity bool
}

func (self *Graph) Copy() *Graph {
	result := &Graph{
		CurrentVertexIndex: self.CurrentVertexIndex,
		numberOfVertices:   self.numberOfVertices,
		numberOfEdges:      self.numberOfEdges,
	}
	result.Vertices = make(Vertices, len(self.Vertices))
	copy(result.Vertices, self.Vertices)
	result.IsVertexDeleted = make([]bool, len(self.IsVertexDeleted))
	copy(result.IsVertexDeleted, self.IsVertexDeleted)
	result.Edges = make(Edges, 0, cap(self.Edges))
	result.neighbors = make([][]*Edge, len(self.neighbors))
	result.degrees = make([]int, len(self.degrees))

	for x := range self.neighbors {
		result.neighbors[x] = make([]*Edge, len(self.neighbors[x]))
	}

	for i, edge := range self.Edges {
		utility.Debug("%v", edge)
		if nil == edge {
			continue
		}
		result.AddEdge(edge.From, edge.To)
		result.Edges[i].isDeleted = edge.isDeleted
	}

	copy(result.degrees, self.degrees)
	result.isRegular = self.isRegular
	result.needToComputeRegularity = self.needToComputeRegularity

	return result
}

func (self *Graph) HasVertex(v Vertex) bool {
	if v.ToInt() >= self.CurrentVertexIndex {
		return false
	}

	if v.ToInt() < len(self.IsVertexDeleted) {
		return !self.IsVertexDeleted[v.ToInt()]
	}

	return true
}

func (self *Graph) getNeighborEdges(v Vertex) []*Edge {
	return self.neighbors[v.ToInt()]
}

func (self *Graph) GetNeighbors(v Vertex) Neighbors {
	result := make(Neighbors, 0, len(self.getNeighborEdges(v)))
	self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
		result = result.appendIfNotContains(GetOtherVertex(v, edge))
	})

	return result
}

func (self *Graph) GetNeighborsWithSet(v Vertex) (Neighbors, mapset.Set) {
	resultSet := mapset.NewSet()
	result := make(Neighbors, 0, len(self.getNeighborEdges(v)))
	self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
		w := GetOtherVertex(v, edge)
		if !resultSet.Contains(w) {
			resultSet.Add(w)
			result = append(result, w)
		}
	})

	return result, resultSet
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
		if self.IsVertexDeleted[vertex.ToInt()] {
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

func (self *Graph) HasEdge(a, b Vertex) bool {
	edge := self.GetEdgeByCoordinates(a.ToInt(), b.ToInt())
	return edge != nil && !edge.isDeleted
}

func (g *Graph) AddVertex() error {
	g.CurrentVertexIndex++
	utility.Debug("Adding %v", g.CurrentVertexIndex)
	g.Vertices = append(g.Vertices, Vertex(g.CurrentVertexIndex))
	g.IsVertexDeleted = append(g.IsVertexDeleted, false)
	g.degrees = append(g.degrees, 0)
	g.numberOfVertices++
	for y, neighbor := range g.neighbors {
		g.neighbors[y] = append(neighbor, nil)
	}

	g.neighbors = append(g.neighbors, make([]*Edge, g.CurrentVertexIndex))
	g.needToComputeRegularity = true

	return nil
}

func (self *Graph) RemoveVertex(v Vertex) error {
	if !self.HasVertex(v) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the ", v))
	}

	self.ForAllNeighbors(v, func(edge *Edge, done chan<- bool) {
		self.RemoveEdge(edge.From, edge.To)
	})

	self.removeVertex(v)
	return nil
}

func (self *Graph) RestoreVertex(v Vertex) {
	vi := v.ToInt()
	if !self.IsVertexDeleted[vi] {
		return
	}

	for _, edge := range self.neighbors[vi] {
		if nil == edge {
			continue
		}

		edge.isDeleted = false
		self.neighbors[GetOtherVertex(v, edge).ToInt()][vi].isDeleted = false
		self.degrees[edge.From.ToInt()]++
		self.degrees[edge.To.ToInt()]++
		self.numberOfEdges++
	}
}

func (self *Graph) removeVertex(v Vertex) {
	self.IsVertexDeleted[v.ToInt()] = true
	self.numberOfVertices--
	self.needToComputeRegularity = true
}

func (self *Graph) computeRegularity() {
	deg := -1
	isRegular := true
	self.ForAllVertices(func(v Vertex, done chan<- bool) {
		if deg == -1 {
			deg = self.Degree(v)
		} else if deg != self.Degree(v) {
			isRegular = false
			done <- true
		}
	})

	self.isRegular = isRegular
	self.needToComputeRegularity = false
}

func (self *Graph) AddEdge(a, b Vertex) error {
	if a == b {
		return errors.New(fmt.Sprintf("Cannot connect vertex %v with itself.", a))
	}

	if !self.HasVertex(a) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the ", a))
	}

	if !self.HasVertex(b) {
		return errors.New(fmt.Sprintf("Vertex %v does not exist in the ", b))
	}

	if self.HasEdge(a, b) {
		return errors.New(fmt.Sprintf("An edge between %v and %v already exists.", a, b))
	}

	edge := MkEdge(a, b)
	self.Edges = append(self.Edges, edge)
	self.neighbors[a.ToInt()][b.ToInt()] = edge
	self.neighbors[b.ToInt()][a.ToInt()] = edge

	ai := a.ToInt()
	bi := b.ToInt()
	self.degrees[ai]++
	self.degrees[bi]++
	self.numberOfEdges++
	self.needToComputeRegularity = true
	return nil
}

func (self *Graph) RemoveEdge(from, to Vertex) {
	fi, ti := from.ToInt(), to.ToInt()
	edge := self.GetEdgeByCoordinates(fi, ti)
	self.degrees[fi] -= 1
	self.degrees[ti] -= 1
	edge.isDeleted = true
	self.numberOfEdges--
	self.needToComputeRegularity = true
}

func (self *Graph) IsVertexCover(vertices ...Vertex) bool {
	n := self.NEdges()
	amountCovered := 0
	isCovered := make([][]bool, self.CurrentVertexIndex)
	for i := range isCovered {
		isCovered[i] = make([]bool, self.CurrentVertexIndex)
	}

	for _, vertex := range vertices {
		utility.Debug("Checking %v for coverage", vertex)
		self.ForAllEdges(func(edge *Edge, done chan<- bool) {
			if edge.IsCoveredBy(vertex) {
				if !isCovered[edge.From.ToInt()][edge.To.ToInt()] {
					utility.Debug("Edge %v -> Covered", *edge)
					amountCovered++
					isCovered[edge.From.ToInt()][edge.To.ToInt()] = true
				}
			}
		})
	}

	utility.Debug("Coverage map for %v: %v", vertices, isCovered)
	return amountCovered == n
}

func (self *Graph) Degree(v Vertex) int {
	if !self.HasVertex(v) {
		panic(errors.New(fmt.Sprintf("Vertex %v does not exist in the ", v)))
	}

	return self.degrees[v.ToInt()]
}

func (self *Graph) IsRegular() bool {
	if self.needToComputeRegularity {
		self.computeRegularity()
	}

	return self.isRegular
}

func mkGraph(vertices, capacity int) *Graph {
	g := new(Graph)
	g.Vertices = make(Vertices, vertices, capacity)
	for i := 0; i < vertices; i++ {
		g.Vertices[i] = MkVertex(i)
	}

	g.CurrentVertexIndex = vertices
	// NOTE: Duplicate edges are not allowed, thus the maximum number of edges in the graph is V^2
	// (when every vertex is connected to each other)
	g.Edges = make(Edges, 0, capacity*capacity)
	g.degrees = make([]int, vertices, capacity)
	g.neighbors = make([][]*Edge, capacity)
	for y := range g.neighbors {
		g.neighbors[y] = make([]*Edge, capacity)
	}

	g.IsVertexDeleted = make([]bool, vertices, capacity)
	g.numberOfVertices = vertices
	g.needToComputeRegularity = true
	return g
}

func IsIndependentSet(set mapset.Set, g *Graph) (result bool, dependent Edges) {
	dependent = make(Edges, 0, g.NEdges())
	result = true
	for vi := range set.Iter() {
		v1 := vi.(Vertex)
		for vi := range set.Iter() {
			v2 := vi.(Vertex)
			if v1 == v2 {
				continue
			}

			if g.HasEdge(v1, v2) {
				result = false
				dependent = append(dependent, MkEdge(v1, v2))
			}
		}
	}

	return
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
