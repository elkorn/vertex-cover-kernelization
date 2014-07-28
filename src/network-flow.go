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
	result := make([][]*NetArc, len(g.Vertices))
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

func (self *NetworkFlow) ComputeMaxFlow() Edges {
	// It has to return a set of edges constituting the max. flow
	result := Edges{}

	return result
}

// https://sites.google.com/site/indy256/algo/dinic_flow
func (self *NetworkFlow) bfs() (bool, []int) {
	// Define `dist[v]` to be the length of the shortest
	// path from source to v in the current instance.
	dist := make([]int, len(self.net))
	for i := range dist {
		dist[i] = -1
	}

	dist[self.source] = 0
	queue := MkQueue(len(self.net))
	queue.Push(int(self.source - 1))
	for i := 0; i < queue.count; i++ {
		from := queue.Pop()
		Debug("From: %v", from)
		for to, arc := range self.net[from] {
			Debug("To: %v (%v)", to, arc)
			if nil == arc {
				continue
			}

			if dist[to] < 0 && arc.residuum() > 0 {
				dist[to] = dist[from] + 1
				queue.Push(to)
			}
		}
	}

	return dist[int(self.sink-1)] >= 0, dist
}
