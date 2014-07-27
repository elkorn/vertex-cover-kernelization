package graph

type NetworkFlow struct {
	source, sink Vertex
	graph        *Graph
	arcs         []*NetArc
}

type NetArc struct {
	capacity int
}

func mkNetArc() *NetArc {
	return &NetArc{1}
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

	result.arcs = make([]*NetArc, len(bipartite.Edges))
	for i := range bipartite.Edges {
		result.arcs[i] = mkNetArc()
	}

	return result
}
