package graph

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

var empty_neighborhood = Neighbors{}

func TestRemoveOfDegree(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)

	g.removeVerticesOfDegree(4)
	assert.True(t, g.HasVertex(1))
	assert.False(t, g.HasVertex(2))
	assert.True(t, g.HasVertex(3))
	assert.True(t, g.HasVertex(4))
	assert.True(t, g.HasVertex(5))

	assert.False(t, g.HasEdge(1, 2))
	assert.False(t, g.HasEdge(2, 3))
	assert.False(t, g.HasEdge(4, 2))
	assert.False(t, g.HasEdge(5, 2))
}

func TestGetVerticesOfDegreeWithOnlyAdjacentNeighbors(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	result, _ := g.getVerticesOfDegreeWithOnlyAdjacentNeighbors(2)
	Debug("%v", result)
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
	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))
	assert.False(t, g.HasVertex(5))
	assert.True(t, g.HasVertex(4))
	assert.True(t, g.HasVertex(1))

	assert.False(t, g.HasEdge(2, 3))
	assert.False(t, g.HasEdge(3, 5))
	assert.True(t, g.HasEdge(1, 4))
}

func TestRemoveVertivesOfDegreeWithOnlyAdjacentNeighbors(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	g.removeVertivesOfDegreeWithOnlyAdjacentNeighbors(2)

	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))
	assert.False(t, g.HasVertex(5))
	assert.True(t, g.HasVertex(4))
	assert.True(t, g.HasVertex(1))

	assert.False(t, g.HasEdge(2, 3))
	assert.False(t, g.HasEdge(3, 5))
	assert.True(t, g.HasEdge(1, 4))
}

func TestGetVerticesOfDegreeWithOnlyDisjointNeighbors(t *testing.T) {
	g := MkGraph3()

	result, _ := g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2)
	assert.Equal(t, Neighbors{2, 3}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])

	g = MkGraph4()

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

	result, _ = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(3)
	assert.Equal(t, empty_neighborhood, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, empty_neighborhood, result[7])

	g = MkGraph4()

	g.addVertex()
	g.AddEdge(1, 8)
	/*
	           1-----8
	          / \
	     3---+   +---2
	    / \         / \
	   7---6       5---4

	*/

	result, _ = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(3)
	assert.Equal(t, Neighbors{2, 3, 8}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, empty_neighborhood, result[7])

	// Edge case: neighbors of a vertex with degree of 1.
	result, _ = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(1)
	assert.Equal(t, empty_neighborhood, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, Neighbors{1}, result[7])

	g = MkGraph5()
	result, _ = g.getVerticesOfDegreeWithOnlyDisjointNeighbors(2)
	assert.Equal(t, Neighbors{2, 3}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, Neighbors{2, 7}, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
}

func TestContractEdges(t *testing.T) {
	g := MkGraph4()
	// ShowGraph(g)
	contractionMap := make(NeighborMap, 1)
	contractionMap[0] = Neighbors{2, 3}
	g.contractEdges(contractionMap)

	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))

	assert.True(t, g.HasEdge(1, 4))
	assert.True(t, g.HasEdge(1, 5))
	assert.True(t, g.HasEdge(1, 6))
	assert.True(t, g.HasEdge(1, 7))

	g = MkGraph5()
	contractionMap = make(NeighborMap, 6)
	contractionMap[0] = Neighbors{2, 3}
	contractionMap[5] = Neighbors{2, 7}
	g.contractEdges(contractionMap)

	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))
	assert.False(t, g.HasVertex(7))

	assert.True(t, g.HasEdge(1, 4))
	assert.True(t, g.HasEdge(1, 5))
	assert.True(t, g.HasEdge(6, 4))
	assert.True(t, g.HasEdge(6, 5))

	g = MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)

	vc := branchAndBound(g)
	var folds mapset.Set
	folds = preprocessing4(g)

	vc2 := branchAndBound(g)
	size := vc2.Cardinality()
	for foldInter := range folds.Iter() {
		fold := foldInter.(*fold)
		if vc2.Contains(fold.replacement) {
			size += 2
		}
	}

	assert.Equal(t, vc.Cardinality(), size)

	g = MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	g.fold(Vertex(1))
	assert.Equal(t, 4, g.NEdges())
	assert.Equal(t, 4, g.Degree(Vertex(8)))

	g = MkGraph(13)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(1, 6)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(2, 8)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(4, 9)
	g.AddEdge(5, 10)
	g.AddEdge(6, 11)
	g.AddEdge(7, 12)
	g.AddEdge(8, 13)

	g.fold(Vertex(1))
	g.fold(Vertex(2))
	g.fold(Vertex(3))

	assert.Equal(t, 5, g.NEdges())
	assert.Equal(t, 5, g.Degree(Vertex(16)))
}

func TestPreprocessingMainRoutine(t *testing.T) {
	g := MkGraph(11)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)
	g.AddEdge(5, 8)
	g.AddEdge(8, 10)
	g.AddEdge(8, 11)
	g.AddEdge(10, 11)

	vc := branchAndBound(g)
	parameterReduction, folds := Preprocessing(g)
	vc2 := branchAndBound(g)
	size := computeUnfoldedVertexCoverSize(folds, vc2) + parameterReduction
	assert.Equal(t, vc.Cardinality(), size)
}
