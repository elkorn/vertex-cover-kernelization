package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaterializeVertexDiscontinuityHandlingError(t *testing.T) {
	g := mkGraphWithVertices(7)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(5, 7)
	g.AddEdge(5, 6)

	Preprocessing(g)
	kPrime := networkFlowKernelization(g, 3)

	assert.Equal(t, 3, kPrime)
	assert.True(t, g.hasVertex(2))
	assert.True(t, g.hasVertex(4))
	assert.True(t, g.hasVertex(5))
	assert.True(t, g.hasVertex(6))
	assert.False(t, g.hasVertex(1))
	assert.False(t, g.hasVertex(3))

	assert.True(t, g.hasEdge(2, 4))
	assert.True(t, g.hasEdge(2, 5))
	assert.True(t, g.hasEdge(2, 6))
}
