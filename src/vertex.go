package graph

type Vertex int

func (self Vertex) toInt() int {
	return int(self) - 1
}

func MkVertex(src int) Vertex {
	return Vertex(src + 1)
}

func (self *Graph) generateVertex() Vertex {
	candidate := Vertex(len(self.Vertices))
	for self.hasVertex(candidate) {
		candidate = Vertex(candidate + 1)
	}

	return candidate
}
