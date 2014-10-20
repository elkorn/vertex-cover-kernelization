package graph

import "fmt"

type Edge struct {
	From      Vertex
	To        Vertex
	isDeleted bool
}

func (self *Edge) IsCoveredBy(v Vertex) bool {
	return self.From == v || self.To == v
}

func (self Edge) GetEndpoints() (Vertex, Vertex) {
	return self.From, self.To
}

func (self Edge) GetIntEndpoints() (int, int) {
	a, b := self.GetEndpoints()
	return a.ToInt(), b.ToInt()
}

func MkEdge(a, b Vertex) *Edge {
	return &Edge{a, b, false}
}

func MkEdgeFromInts(a, b int) *Edge {
	return MkEdge(MkVertex(a), MkVertex(b))
}

func MkEdgeValFromInts(a, b int) Edge {
	return Edge{MkVertex(a), MkVertex(b), false}
}

func (self *Edge) Str() string {
	return fmt.Sprintf("%v-%v", self.From, self.To)
}

type Edges []*Edge
