package graph

func (self *Graph) removeVerticesWithDegreeGreaterThan(k int) Neighbors {
	result := Neighbors{}

	for _, vertex := range self.Vertices {
		if vertex.degree > k {
			result = append(result, vertex)
			self.RemoveVertex(vertex.id)
		}
	}

	return result
}

// ILP forumlation is the second mehtod, but it has been moved to a separate file.
