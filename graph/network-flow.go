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
		arcs:   make([][]*NetArc, g.NVertices()),
		length: make([]int, g.NVertices()),
	}

	for i := 0; i < g.NVertices(); i++ {
		result.arcs[i] = make([]*NetArc, g.NVertices())
	}

	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		from := edge.from.toInt()
		to := edge.to.toInt()
		result.arcs[from][to] = mkNetArc(edge)
		result.arcs[to][from] = mkNetArc(nil)
		result.length[from] += 1
	})

	return result
}

func connectSourceAndSink(bipartite *Graph, source, sink Vertex) {
	verticesAfter := bipartite.NVertices() - 2
	verticesBefore := (verticesAfter) / 2

	for i := 0; i < verticesBefore; i++ {
		bipartite.AddEdge(source, Vertex(i+1))
	}

	for i := verticesBefore; i < verticesAfter; i++ {
		bipartite.AddEdge(Vertex(i+1), sink)
	}
}

func mkNetworkFlow(g *Graph) *NetworkFlow {
	bipartite := makeBipartiteForNetworkFlow(g)
	result := &NetworkFlow{
		graph:  bipartite,
		source: Vertex(bipartite.currentVertexIndex - 1),
	}

	result.sink = Vertex(bipartite.currentVertexIndex)
	connectSourceAndSink(bipartite, result.source, result.sink)

	result.net = mkNet(bipartite)
	return result
}
