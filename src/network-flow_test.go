package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkToNetworkFlow(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	source := Vertex(7)
	sink := Vertex(8)

	result := mkNetworkFlow(g)

	// 1. we need source and sink nodes.
	assert.True(t, result.graph.hasVertex(7))
	assert.True(t, result.graph.hasVertex(8))
	assert.Equal(t, source, result.source)
	assert.Equal(t, sink, result.sink)
	// 2. Source has to be connected to every vertex in set A.
	for i := 1; i <= 3; i++ {
		assert.True(t, result.graph.hasEdge(source, Vertex(i)), "Source has to be connected to all edges in set A.")
	}

	// 3. Sink has to be connected to every vertex in set B.
	for i := 4; i <= 6; i++ {
		assert.True(t, result.graph.hasEdge(Vertex(i), sink), "Sink has to be connected to all edges in set B.")
	}
}
