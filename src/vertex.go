package graph

type Vertex struct {
	// This will become Vertex after refactoring.
	id     int
	degree int
}

func mkVertex(id int) Vertex {
	return Vertex{id, 0}
}

func (self Vertex) eq(other *Vertex) bool {
	return other != nil && self.id == other.id
}

func (self *Graph) generateVertex() Vertex {
	candidate := len(self.Vertices)
	for self.hasVertex(candidate) {
		candidate += 1
	}

	return mkVertex(candidate)
}
