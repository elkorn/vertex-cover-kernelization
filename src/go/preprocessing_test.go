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
	assert.Equal(t, true, g.hasVertex(1))
	assert.Equal(t, false, g.hasVertex(2))
	assert.Equal(t, true, g.hasVertex(3))
	assert.Equal(t, true, g.hasVertex(4))
	assert.Equal(t, true, g.hasVertex(5))

	assert.Equal(t, false, g.hasEdge(1, 2))
	assert.Equal(t, false, g.hasEdge(2, 3))
	assert.Equal(t, false, g.hasEdge(4, 2))
	assert.Equal(t, false, g.hasEdge(5, 2))
}

// func TestGetVerticesOfDegreeWithOnlyAdjacentNeighbors(t *testing.T) {
// 	g := mkGraphWithVertices(5)
//
// 	g.AddEdge(2, 5)
// 	g.AddEdge(3, 5)
// 	g.AddEdge(2, 3)
// 	g.AddEdge(1, 4)
//
// 	result := g.getVerticesOfDegreeWithOnlyAdjacentNeighbors(2)
//
// 	assert.Equal(t, Neighbors{2, 3}, result[5])
// }
