package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFordFulkerson(t *testing.T) {
	g := mkGraphWithVertices(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	nf := &NetworkFlow{
		graph:  g,
		source: Vertex(1),
		sink:   Vertex(4),
		net:    mkNet(g),
	}

	flowPath, flowValue := fordFulkerson(nf)

	assert.Equal(t, Edges{&Edge{1, 2}, &Edge{2, 3}, &Edge{3, 4}}, flowPath)
	assert.Equal(t, flowValue, 3)
}

func TestFordFulkerson2(t *testing.T) {
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

	flowPath, flowValue := fordFulkerson(nf)

	assert.Equal(t, Edges{&Edge{1, 5}, &Edge{5, 4}, &Edge{4, 6}}, flowPath)
	assert.Equal(t, flowValue, 3)
}
