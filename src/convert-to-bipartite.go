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

func addBipartiteEdges(originalEdges []*Edge) []*Edge {
	before := len(originalEdges)
	after := 2 * before
	border := before + 1
	result := make([]*Edge, after)
	// Invariant: F = {(A_v,B_v)|(v,u) \in E or (u,v) \in E}
	for i, edge := range originalEdges {
		result[i] = edge
	}

	for i := before; i < after; i++ {
		orig := originalEdges[i-before]
		result[i] = &Edge{
			Vertex(int(orig.from) + border),
			Vertex(int(orig.to) + border),
		}
	}

	return result
}
