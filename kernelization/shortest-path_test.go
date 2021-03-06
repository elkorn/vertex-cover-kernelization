package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 4)

	nf := &NetworkFlow{
		source: graph.Vertex(1),
		sink:   graph.Vertex(4),
		graph:  g,
		net:    mkNet(g),
	}

	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 3}, path.Values(), "The path nodes have to be correct.")
	assert.Equal(t, []int{0, 1, 2, 1}, distance, "The path distance has to be correct.")
}

func TestShortestPath2(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	nf := &NetworkFlow{
		source: graph.Vertex(1),
		sink:   graph.Vertex(4),
		graph:  g,
		net:    mkNet(g),
	}
	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 1, 2, 3}, path.Values(), "The path nodes have to be correct.")
	assert.Equal(t, []int{0, 1, 2, 3}, distance, "The path distance has to be correct.")
}

func TestShortestPathUndirected(t *testing.T) {
	g := graph.MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(5, 4)

	nf := &NetworkFlow{
		graph:  g,
		source: graph.Vertex(1),
		sink:   graph.Vertex(6),
		net:    mkNet(g),
	}

	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in an undirected graph has to be found.")
	assert.Equal(t, []int{0, 4, 3, 5}, path.Values(), "The correct path has to be found in an undirected ")
	assert.Equal(t, []int{0, 1, 2, 2, 1, 3}, distance, "The distance has to be correct.")
}

func TestShortestPathArbitraryEndpoints(t *testing.T) {
	g := graph.MkGraph(6)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(5, 4)

	n1 := mkNet(g)
	exists, path, distance := shortestPath(n1, 2, 6)
	iter := path.Iter()
	from := <-iter

	for p := range iter {
		to := p
		n1.arcs[from][to].flow = 1
		from = to
	}

	assert.True(t, exists, "The path in an undirected graph has to be found.")
	assert.Equal(t, []int{1, 2, 3, 5}, path.Values(), "The correct path has to be found in an undirected ")
	assert.Equal(t, []int{1, 0, 1, 2, 2, 3}, distance, "The distance has to be correct.")
}
