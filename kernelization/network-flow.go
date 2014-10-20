package kernelization

import (
	"math"

	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type NetworkFlow struct {
	source, sink graph.Vertex
	graph        *graph.Graph
	net          Net
}

type NetArc struct {
	capacity    int
	flow        int
	reverseFlow int
	edge        *graph.Edge
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func (self *NetArc) residuum() int {
	return self.capacity - self.flow
}

func mkNetArc(e *graph.Edge) *NetArc {
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

func (self *Net) Capacity(edge *graph.Edge) int {
	return (*self).arcs[edge.From-1][edge.To-1].capacity
}

func (self *Net) Flow(edge *graph.Edge) int {
	return (*self).arcs[edge.From-1][edge.To-1].flow
}

func (self *Net) Residuum(edge *graph.Edge) int {
	return (*self).arcs[edge.From-1][edge.To-1].residuum()
}

func mkNet(g *graph.Graph) Net {
	result := Net{
		arcs:   make([][]*NetArc, g.NVertices()),
		length: make([]int, g.NVertices()),
	}

	for i := 0; i < g.NVertices(); i++ {
		result.arcs[i] = make([]*NetArc, g.NVertices())
	}

	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		from := edge.From.ToInt()
		to := edge.To.ToInt()
		result.arcs[from][to] = mkNetArc(edge)
		result.arcs[to][from] = mkNetArc(nil)
		result.length[from] += 1
	})

	return result
}

func connectSourceAndSink(bipartite *graph.Graph, source, sink graph.Vertex) {
	verticesAfter := bipartite.NVertices() - 2
	verticesBefore := (verticesAfter) / 2

	for i := 0; i < verticesBefore; i++ {
		bipartite.AddEdge(source, graph.Vertex(i+1))
	}

	for i := verticesBefore; i < verticesAfter; i++ {
		bipartite.AddEdge(graph.Vertex(i+1), sink)
	}
}

func mkNetworkFlow(g *graph.Graph) *NetworkFlow {
	bipartite := makeBipartiteForNetworkFlow(g)
	result := &NetworkFlow{
		graph:  bipartite,
		source: graph.Vertex(bipartite.CurrentVertexIndex - 1),
	}

	result.sink = graph.Vertex(bipartite.CurrentVertexIndex)
	connectSourceAndSink(bipartite, result.source, result.sink)

	result.net = mkNet(bipartite)
	return result
}
