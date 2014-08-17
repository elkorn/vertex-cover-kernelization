package graph

type Edge struct {
	from Vertex
	to   Vertex
}

func (self *Edge) IsCoveredBy(v Vertex) bool {
	return self.from == v || self.to == v
}

func MkEdge(a, b Vertex) *Edge {
	return &Edge{a, b}
}

func MkEdgeFromInts(a, b int) *Edge {
	return MkEdge(MkVertex(a), MkVertex(b))
}

type Edges []*Edge
