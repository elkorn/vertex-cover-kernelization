package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOfDegree(t *testing.T) {
	g := mkGraphWithVertices(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)

	g.removeVerticesOfDegree(4)
	assert.Equal(t, 4, len(g.Vertices))
	assert.True(t, g.hasVertex(1))
	assert.False(t, g.hasVertex(2))
	assert.True(t, g.hasVertex(3))
	assert.True(t, g.hasVertex(4))
	assert.True(t, g.hasVertex(5))

	assert.False(t, g.hasEdge(1, 2))
	assert.False(t, g.hasEdge(2, 3))
	assert.False(t, g.hasEdge(4, 2))
	assert.False(t, g.hasEdge(5, 2))
}

func TestGetVerticesOfDegreeWithOnlyAdjacentNeighbors(t *testing.T) {
	g := mkGraphWithVertices(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	result := g.getVerticesOfDegreeWithOnlyAdjacentNeighbors(2)

	assert.Equal(t, Neighbors{2, 3}, result[5])
	assert.Equal(t, Neighbors{5, 3}, result[2])
	assert.Equal(t, Neighbors{2, 5}, result[3])
}

func TestRemoveAllVerticesAccordingToMap(t *testing.T) {
	g := mkGraphWithVertices(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	theMap := make(NeighborMap)
	theMap[5] = Neighbors{2, 3}
	theMap[2] = Neighbors{5, 3}
	theMap[3] = Neighbors{2, 5}

	g.removeAllVerticesAccordingToMap(theMap)
	assert.False(t, g.hasVertex(2))
	assert.False(t, g.hasVertex(3))
	assert.False(t, g.hasVertex(5))
	assert.True(t, g.hasVertex(4))
	assert.True(t, g.hasVertex(1))

	assert.False(t, g.hasEdge(2, 3))
	assert.False(t, g.hasEdge(3, 5))
	assert.True(t, g.hasEdge(1, 4))
}

func TestRemoveVertivesOfDegreeWithOnlyAdjacentNeighbors(t *testing.T) {
	g := mkGraphWithVertices(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	g.removeVertivesOfDegreeWithOnlyAdjacentNeighbors(2)
	assert.False(t, g.hasVertex(2))
	assert.False(t, g.hasVertex(3))
	assert.False(t, g.hasVertex(5))
	assert.True(t, g.hasVertex(4))
	assert.True(t, g.hasVertex(1))

	assert.False(t, g.hasEdge(2, 3))
	assert.False(t, g.hasEdge(3, 5))
	assert.True(t, g.hasEdge(1, 4))
}
