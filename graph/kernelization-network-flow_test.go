package graph_test

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestMaterializeVertexDiscontinuityHandlingError(t *testing.T) {
	g := graph.MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(5, 7)
	g.AddEdge(5, 6)

	kPrime := graph.NetworkFlowKernelization(g, 3)

	assert.Equal(t, 1, kPrime)
	assert.True(t, g.HasVertex(2))
	assert.True(t, g.HasVertex(4))
	assert.False(t, g.HasVertex(5))
	assert.False(t, g.HasVertex(6))
	assert.False(t, g.HasVertex(1))
	assert.False(t, g.HasVertex(3))

	assert.True(t, g.HasEdge(2, 4))
	assert.False(t, g.HasEdge(2, 5))
	assert.False(t, g.HasEdge(2, 6))
}
