package graph

type Vertex int

func (self *Graph) generateVertex() Vertex {
	candidate := Vertex(len(self.Vertices))
	for self.hasVertex(candidate) {
		candidate = Vertex(candidate + 1)
	}

	return candidate
}
