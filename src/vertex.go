package graph

type Vertex int

type Node struct {
	// This will become Vertex after refactoring.
	Vertex
	degree int
}

func (self *Graph) generateVertex() Vertex {
	candidate := Vertex(len(self.Vertices))
	for self.hasVertex(candidate) {
		candidate = Vertex(candidate + 1)
	}

	return candidate
}
