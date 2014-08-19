package graph

func (self *Graph) removeVerticesWithDegreeGreaterThan(k int) Neighbors {
	degrees := make(map[Vertex]int)
	result := Neighbors{}

	self.ForAllEdges(func(edge *Edge, _ int, done chan<- bool) {
		degrees[edge.from]++
		degrees[edge.to]++
	})

	for vertex, degree := range degrees {
		if degree > k {
			result = append(result, vertex)
			self.RemoveVertex(vertex)
		}
	}

	return result
}

// ILP forumlation is the second mehtod, but it has been moved to a separate file.
