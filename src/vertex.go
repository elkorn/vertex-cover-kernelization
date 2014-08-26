package graph

type Vertex int

var INVALID_VERTEX Vertex = Vertex(0)

func (self Vertex) toInt() int {
	return int(self) - 1
}

func MkVertex(src int) Vertex {
	return Vertex(src + 1)
}

func (self *Graph) generateVertex() Vertex {
	candidate := Vertex(self.currentVertexIndex)
	for self.hasVertex(candidate) {
		candidate = Vertex(candidate + 1)
	}

	return candidate
}
