package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestFordFulkerson(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	nf := &NetworkFlow{
		graph:  g,
		source: graph.Vertex(1),
		sink:   graph.Vertex(4),
		net:    mkNet(g),
	}

	flowPath, flowValue := fordFulkerson(nf)
	assert.Equal(t, graph.Edges{graph.MkEdgeFromInts(0, 1), graph.MkEdgeFromInts(1, 2), graph.MkEdgeFromInts(2, 3)}, flowPath)
	assert.Equal(t, flowValue, 3)
}

func TestFordFulkerson2(t *testing.T) {
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

	flowPath, flowValue := fordFulkerson(nf)

	assert.Equal(t, graph.Edges{graph.MkEdgeFromInts(0, 4), graph.MkEdgeFromInts(4, 3), graph.MkEdgeFromInts(3, 5)}, flowPath)
	assert.Equal(t, flowValue, 3)
}
