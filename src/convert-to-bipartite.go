package graph

func getVertices(g *Graph) []Vertex {
	result := make([]Vertex, 0)
	n := len(g.Vertices)
	for _, v := range g.Vertices {
		result = append(result, v)
	}

	for _, v := range result {
		result = append(result, MkVertex(v.toInt()+n))
	}

	return result
}

func addBipartiteEdges(g *Graph, originalEdges Edges, border int) {
	before := len(originalEdges)
	for i := 0; i < before; i++ {
		orig := originalEdges[i]
		// Invariant: F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
		g.AddEdge(orig.from, Vertex(orig.to.toInt()+border+1))
	}
}

func makeBipartite(g *Graph) *Graph {
	/*
		Convert G(V,E) to a bipartite graph H=(U,F) with the following properties:
		A = V
		B = V
		U = A \sum B
		F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
	*/

	// TODO this is going to cause problems if there are discontinuities in the vertex collection.
	border := len(g.Vertices)
	result := mkGraphWithVertices(border * 2)
	addBipartiteEdges(result, g.Edges, border)
	return result
}
