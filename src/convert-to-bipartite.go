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

func addBipartiteEdges(g *Graph, border int) {
	before := len(g.Edges)
	after := 2 * before
	// Invariant: F = {(A_v,B_v)|(v,u) \in E or (u,v) \in E}
	for i := before; i < after; i++ {
		orig := g.Edges[i-before]
		Debug("Adding edge %v-%v", int(orig.from)+border, int(orig.to)+border)
		g.AddEdge(
			Vertex(int(orig.from)+border),
			Vertex(int(orig.to)+border))
	}

}

func makeBipartite(g *Graph) *Graph {
	/*
		Convert G(V,E) to a bipartite graph H=(U,F) with the following properties:
		A = V
		B = V
		U = A \sum B
		F = {(A_v,B_v)|(v,u) \in E or (u,v) \in E}
	*/

	before := len(g.Edges)
	// TODO this is going to cause problems if there are discontinuities in the vertex collection.
	border := len(g.Vertices)
	result := mkGraphWithVertices(len(g.Vertices) * 2)
	for i := 0; i < before; i++ {
		orig := g.Edges[i]
		result.AddEdge(orig.from, orig.to)
	}

	addBipartiteEdges(result, border)
	return result
}
