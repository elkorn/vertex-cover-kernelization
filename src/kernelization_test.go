package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveHighDegree(t *testing.T) {
	g1 := mkGraphWithVertices(10)

	g1.AddEdge(1, 2)
	g1.AddEdge(1, 3)
	g1.AddEdge(1, 4)
	g1.AddEdge(1, 5)

	g1.AddEdge(2, 3)
	g1.AddEdge(2, 4)
	g1.AddEdge(2, 5)
	g1.AddEdge(2, 6)
	g1.AddEdge(2, 7)

	g1.AddEdge(9, 8)

	res := g1.removeVerticesWithDegreeGreaterThan(2)
	assert.Equal(t, Neighbors{1, 2}, res)
	assert.False(t, g1.hasVertex(1))
	assert.False(t, g1.hasVertex(2))

	assert.False(t, g1.hasEdge(1, 2))
	assert.False(t, g1.hasEdge(1, 3))
	assert.False(t, g1.hasEdge(1, 4))
	assert.False(t, g1.hasEdge(1, 5))

	assert.False(t, g1.hasEdge(2, 3))
	assert.False(t, g1.hasEdge(2, 4))
	assert.False(t, g1.hasEdge(2, 5))
	assert.False(t, g1.hasEdge(2, 6))
	assert.False(t, g1.hasEdge(2, 7))

	assert.True(t, g1.hasEdge(9, 8))
}
