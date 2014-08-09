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
		x := int(edge.from) - 1
		result.arcs[x][(edge.to)-1] = mkNetArc(edge)
		result.length[x] += 1
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

func (self *NetworkFlow) ComputeMaxFlow() int {
	// I gotta grasp this conceptually first.
	return 0
}

func (self *NetworkFlow) isTraversable(from, to int, dist []int) bool {
	arc := self.net.arcs[from][to]
	if nil == arc {
		return false
	}

	return dist[to] < 0 && arc.residuum() > 0
}

// R. Sedgewick
func (self *NetworkFlow) dfs(ptr, dist []int, from, to Vertex, leftoverFlow int) int {
	u := int(from) - 1
	// dest := int(to) - 1
	if from == to {
		return leftoverFlow
	}

	Debug("%v -> %v (%v)", from, to, leftoverFlow)
	Debug("ptr: %v", ptr)
	for ; ptr[u] < self.net.length[u]; ptr[u]++ {
		arc := self.net.arcs[u][ptr[u]]
		// this will be a non-issue if I decide to convert the Net to a dense 2D arrray.
		if nil == arc {
			continue
		}

		if dist[arc.edge.to-1] == dist[u]+1 && arc.residuum() > 0 {
			df := self.dfs(ptr, dist, arc.edge.to, to, min(leftoverFlow, arc.residuum()))
			Debug("Finished DFS")
			if df > 0 {
				arc.flow += df
				return df
			}
		}
	}

	return 0
}
