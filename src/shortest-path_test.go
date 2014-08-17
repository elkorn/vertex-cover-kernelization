package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	g := mkGraphWithVertices(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 4)

	nf := &NetworkFlow{
		source: Vertex(1),
		sink:   Vertex(4),
		graph:  g,
		net:    mkNet(g),
	}
	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 3}, path, "The path nodes have to be correct.")
	assert.Equal(t, []int{0, 1, 2, 1}, distance, "The path distance has to be correct.")
}

func TestShortestPath2(t *testing.T) {
	g := mkGraphWithVertices(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	nf := &NetworkFlow{
		source: Vertex(1),
		sink:   Vertex(4),
		graph:  g,
		net:    mkNet(g),
	}
	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 1, 2, 3}, path, "The path nodes have to be correct.")
	assert.Equal(t, []int{0, 1, 2, 3}, distance, "The path distance has to be correct.")
}

func TestShortestPathUndirected(t *testing.T) {
	g := mkGraphWithVertices(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(5, 4)

	nf := &NetworkFlow{
		graph:  g,
		source: Vertex(1),
		sink:   Vertex(6),
		net:    mkNet(g),
	}

	exists, path, distance := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in an undirected graph has to be found.")
	assert.Equal(t, []int{0, 4, 3, 5}, path, "The correct path has to be found in an undirected graph.")
	assert.Equal(t, []int{0, 1, 2, 2, 1, 3}, distance, "The distance has to be correct.")
}

func TestShortestPathArbitraryEndpoints(t *testing.T) {
	g := mkGraphWithVertices(6)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 6)
	g.AddEdge(5, 4)

	n1 := mkNet(g)
	exists, path, distance := shortestPath(n1, 2, 5)
	for i := 1; i < len(path); i++ {
		n1.arcs[path[i-1]][path[i]].flow = 1
	}

	nv := MkNetworkVisualizer()
	nv.Display(&n1)
	assert.True(t, exists, "The path in an undirected graph has to be found.")
	assert.Equal(t, []int{1, 2, 3, 4}, path, "The correct path has to be found in an undirected graph.")
	assert.Equal(t, []int{4, 0, 1, 2, 3, 3}, distance, "The distance has to be correct.")
}
