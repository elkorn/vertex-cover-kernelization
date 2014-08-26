package graph

type Edge struct {
	from      Vertex
	to        Vertex
	isDeleted bool
}

func (self *Edge) IsCoveredBy(v Vertex) bool {
	return self.from == v || self.to == v
}

func (self Edge) GetEndpoints() (Vertex, Vertex) {
	return self.from, self.to
}

func (self Edge) GetIntEndpoints() (int, int) {
	a, b := self.GetEndpoints()
	return a.toInt(), b.toInt()
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

type Edges []*Edge
