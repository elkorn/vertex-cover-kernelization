package graph

type NetworkFlow struct {
	source, sink Vertex
	graph        *Graph
}

func mkNetworkFlow(g *Graph) *NetworkFlow {
	verticesBefore := len(g.Vertices)
	verticesAfter := verticesBefore * 2
	bipartite := makeBipartite(g)
	bipartite.AddVertex()
	result := &NetworkFlow{
		graph:  bipartite,
		source: Vertex(bipartite.currentVertexIndex),
	}

	bipartite.AddVertex()

	result.sink = Vertex(bipartite.currentVertexIndex)
	for i := 0; i < verticesBefore; i++ {
		bipartite.AddEdge(result.source, Vertex(i+1))
	}

	for i := verticesBefore; i < verticesAfter; i++ {
		bipartite.AddEdge(Vertex(i+1), result.sink)
	}

	return result
}
