package graph

import (
	"errors"
	"fmt"
	"log"
)

type Options struct {
	Verbose bool
}

var options Options

func SetOptions(opts Options) {
	options = opts
}

func debug(msg string) {
	if options.Verbose {
		log.Print(msg)
	}
}

func extend(slice []Edge, element Edge) []Edge {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]Edge, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func append(slice []Edge, items ...Edge) []Edge {
	for _, item := range items {
		slice = extend(slice, item)
	}

	return slice
}

type Vertex int

type Edge struct {
	from Vertex
	to   Vertex
}

func (self *Edge) IsCoveredBy(v Vertex) bool {
	return self.from == v || self.to == v
}

type Graph struct {
	Vertices map[Vertex]bool
	Edges    []Edge
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

func (g *Graph) AddVertex(v Vertex) error {
	_, exists := g.Vertices[v]
	if exists {
		return errors.New(fmt.Sprintf("Vertex %v already in the set", v))
	}

	g.Vertices[v] = true
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

	debug(fmt.Sprintf("Coverage map for %v: %v", vertices, isCovered))
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
	g.Edges = make([]Edge, 0)
	return g
}
