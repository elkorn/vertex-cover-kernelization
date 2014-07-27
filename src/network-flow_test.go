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

func TestNet(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	netFlow := mkNetworkFlow(g)
	assert.Equal(t, len(netFlow.graph.Edges), len(netFlow.net))
	for _, edge := range netFlow.graph.Edges {
		arc := netFlow.net[edge.from][edge.to]
		assert.NotNil(t, arc, "Each edge in the network flow must be represented by an Arc.")
		assert.Equal(t, 1, arc.capacity, "Every arc must have an initial capacity of 1.")
		assert.Equal(t, 0, arc.flow, "Every arc must have an initial flow of 0.")
	}
}

func TestBFS(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	netFlow := mkNetworkFlow(g)
	inVerboseContext(func() {
		result, dist := netFlow.bfs()
		Debug("%v, %v", result, dist)
	})
}
