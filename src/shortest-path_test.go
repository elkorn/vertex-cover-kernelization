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
	exists, path := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 3}, path.Values())
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
	exists, path := shortestPathFromSourceToSink(nf)
	assert.True(t, exists, "The path in the graph should be found by BFS.")
	assert.Equal(t, []int{0, 1, 2, 3}, path.Values())
}
