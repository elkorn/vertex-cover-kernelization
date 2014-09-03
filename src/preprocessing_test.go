package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveOfDegree(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)

	g.removeVerticesOfDegree(4)
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
	g := MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	// TODO: here be errors
	inVerboseContext(func() {
		g.getVerticesOfDegreeWithOnlyAdjacentNeighbors(2)
	})
	result := g.getVerticesOfDegreeWithOnlyAdjacentNeighbors(2)

	assert.Equal(t, Neighbors{2, 3}, result[4])
	assert.Equal(t, Neighbors{3, 5}, result[1])
	assert.Equal(t, Neighbors{2, 5}, result[2])
}

func TestRemoveAllVerticesAccordingToMap(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	theMap := make(NeighborMap, 5)
	theMap[4] = Neighbors{2, 3}
	theMap[1] = Neighbors{5, 3}
	theMap[2] = Neighbors{2, 5}

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
	g := MkGraph(5)

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

func TestGetVerticesOfDegreeWithOnlyDisjointNeighbors(t *testing.T) {
	g := mkGraph3()

	result := g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2)
	assert.Equal(t, Neighbors{2, 3}, result[0])
	assert.Nil(t, result[1])
	assert.Nil(t, result[2])
	assert.Nil(t, result[3])
	assert.Nil(t, result[4])
	assert.Nil(t, result[5])
	assert.Nil(t, result[6])

	g = mkGraph4()

	g.addVertex()

	g.AddEdge(1, 8)
	g.AddEdge(2, 8)
	/*
	           1-----8
	          / \    |
	     3---+   +---2
	    / \         / \
	   7---6       5---4

	*/

	result = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(3)
	assert.Nil(t, result[0])
	assert.Nil(t, result[1])
	assert.Nil(t, result[2])
	assert.Nil(t, result[3])
	assert.Nil(t, result[4])
	assert.Nil(t, result[5])
	assert.Nil(t, result[6])
	assert.Nil(t, result[7])

	g = mkGraph4()

	g.addVertex()
	g.AddEdge(1, 8)
	/*
	           1-----8
	          / \
	     3---+   +---2
	    / \         / \
	   7---6       5---4

	*/

	result = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(3)
	assert.Equal(t, Neighbors{2, 3, 8}, result[0])
	assert.Nil(t, result[1])
	assert.Nil(t, result[2])
	assert.Nil(t, result[3])
	assert.Nil(t, result[4])
	assert.Nil(t, result[5])
	assert.Nil(t, result[6])
	assert.Nil(t, result[7])

	// Edge case: neighbors of a vertex with degree of 1.
	result = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(1)
	assert.Nil(t, result[0])
	assert.Nil(t, result[1])
	assert.Nil(t, result[2])
	assert.Nil(t, result[3])
	assert.Nil(t, result[4])
	assert.Nil(t, result[5])
	assert.Nil(t, result[6])
	assert.Equal(t, Neighbors{1}, result[7])

	g = mkGraph5()
	result = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2)
	assert.Equal(t, Neighbors{2, 3}, result[0])
	assert.Nil(t, result[1])
	assert.Nil(t, result[2])
	assert.Nil(t, result[3])
	assert.Nil(t, result[4])
	assert.Equal(t, Neighbors{2, 7}, result[5])
	assert.Nil(t, result[6])
}

func TestContractEdges(t *testing.T) {
	g := mkGraph4()
	contractionMap := make(NeighborMap, 1)
	contractionMap[0] = Neighbors{2, 3}
	g.contractEdges(contractionMap)

	assert.False(t, g.hasVertex(2))
	assert.False(t, g.hasVertex(3))

	assert.True(t, g.hasEdge(1, 4))
	assert.True(t, g.hasEdge(1, 5))
	assert.True(t, g.hasEdge(1, 6))
	assert.True(t, g.hasEdge(1, 7))

	g = mkGraph5()
	contractionMap = make(NeighborMap, 6)
	contractionMap[0] = Neighbors{2, 3}
	contractionMap[5] = Neighbors{2, 7}
	g.contractEdges(contractionMap)

	assert.False(t, g.hasVertex(2))
	assert.False(t, g.hasVertex(3))
	assert.False(t, g.hasVertex(7))

	assert.True(t, g.hasEdge(1, 4))
	assert.True(t, g.hasEdge(1, 5))
	assert.True(t, g.hasEdge(6, 4))
	assert.True(t, g.hasEdge(6, 5))
}
