package graph

func getVertices(g *Graph) Vertices {
	result := make(Vertices, 0)
	n := len(g.Vertices)
	for _, v := range g.Vertices {
		result = append(result, v)
	}

	for _, v := range result {
		result = append(result, MkVertex(v.toInt()+n))
	}

	return result
}

func addBipartiteEdges(g *Graph, original *Graph, border int) {
	original.ForAllEdges(func(edge *Edge, idx int, done chan<- bool) {
		// Invariant: F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
		g.AddEdge(edge.from, MkVertex(edge.to.toInt()+border))
	})
}

func makeBipartite(g *Graph) *Graph {
	/*
		Convert G(V,E) to a bipartite graph H=(U,F) with the following properties:
		A = V
		B = V
		U = A \sum B
		F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
	*/

	// TODO: this is going to cause problems if there are discontinuities in the vertex collection.
	border := len(g.Vertices)
	result := MkGraphRememberingDeletedVertices(border*2, g.isVertexDeleted) // remember deleted vertices
	addBipartiteEdges(result, g, border)
	return result
}
