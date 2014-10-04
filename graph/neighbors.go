package graph

type Neighbors Vertices

func (self Neighbors) appendIfNotContains(v Vertex) Neighbors {
	if !contains(self, v) {
		self = append(self, v)
	}

	return self
}
