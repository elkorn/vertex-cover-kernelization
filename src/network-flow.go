package graph

import "math"

type NetworkFlow struct {
	source, sink Vertex
	graph        *Graph
	net          Net
}

type NetArc struct {
	capacity    int
	flow        int
	reverseFlow int
	edge        *Edge
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func (self *NetArc) residuum() int {
	return self.capacity - self.flow
}

func mkNetArc(e *Edge) *NetArc {
	return &NetArc{
		capacity: 1,
		flow:     0,
		edge:     e,
	}
}

type Net struct {
	arcs   [][]*NetArc
	length []int
}

func (self *Net) Capacity(edge *Edge) int {
	return (*self).arcs[edge.from-1][edge.to-1].capacity
}

func (self *Net) Flow(edge *Edge) int {
	return (*self).arcs[edge.from-1][edge.to-1].flow
}

func (self *Net) Residuum(edge *Edge) int {
	return (*self).arcs[edge.from-1][edge.to-1].residuum()
}

func mkNet(g *Graph) Net {
	result := Net{
		arcs:   make([][]*NetArc, len(g.Vertices)),
		length: make([]int, len(g.Vertices)),
	}

	for i := 0; i < len(g.Vertices); i++ {
		result.arcs[i] = make([]*NetArc, len(g.Vertices))
	}

	for _, edge := range g.Edges {
		from := edge.from.toInt()
		to := edge.to.toInt()
		result.arcs[from][to] = mkNetArc(edge)
		result.arcs[to][from] = mkNetArc(nil)
		result.length[from] += 1
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

func (self *NetworkFlow) isTraversable(from, to int, dist []int) bool {
	arc := self.net.arcs[from][to]
	if nil == arc {
		return false
	}

	return dist[to] < 0 && arc.residuum() > 0
}
