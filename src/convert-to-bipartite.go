package graph

func getVertices(g *Graph) []Vertex {
	result := make([]Vertex, 0)
	n := len(g.Vertices)
	for v := range g.Vertices {
		result = append(result, v)
	}

	for _, v := range result {
		result = append(result, Vertex(int(v)+n))
	}

	return result
}
