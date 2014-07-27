package graph

type NetworkFlow struct {
	source, sink Vertex
	graph        *Graph
	net          Net
}

type NetArc struct {
	capacity int
	flow     int
}

func (self *NetArc) residuum() int {
	return self.capacity - self.flow
}

func mkNetArc() *NetArc {
	return &NetArc{1, 0}
}

type Net [][]*NetArc

func mkNet(g *Graph) Net {
	result := make([][]*NetArc, len(g.Edges))
	for i := 0; i < len(g.Vertices); i++ {
		result[i] = make([]*NetArc, len(g.Vertices))
	}

	for _, edge := range g.Edges {
		result[int(edge.from)-1][(edge.to)-1] = mkNetArc()
	}

	return result
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

	result.net = mkNet(bipartite)
	return result
}
