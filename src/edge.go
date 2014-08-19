package graph

type Edge struct {
	from      Vertex
	to        Vertex
	isDeleted bool
}

func (self *Edge) IsCoveredBy(v Vertex) bool {
	return self.from == v || self.to == v
}

func MkEdge(a, b Vertex) *Edge {
	return &Edge{a, b, false}
}

func MkEdgeFromInts(a, b int) *Edge {
	return MkEdge(MkVertex(a), MkVertex(b))
}

type Edges []*Edge
