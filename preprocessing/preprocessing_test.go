package preprocessing

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/elkorn/vertex-cover-kernelization/vc"
	"github.com/stretchr/testify/assert"
)

var empty_neighborhood = graph.Neighbors{}

func TestRemoveOfDegree(t *testing.T) {
	g := graph.MkGraph(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)

	removeVerticesOfDegree(g, 4)
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
	g := graph.MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	result, _ := getVerticesOfDegreeWithOnlyAdjacentNeighbors(g, 2)
	utility.Debug("%v", result)
	assert.Equal(t, graph.Neighbors{2, 3}, result[4])
	assert.Equal(t, graph.Neighbors{3, 5}, result[1])
	assert.Equal(t, graph.Neighbors{2, 5}, result[2])

}

func TestRemoveAllVerticesAccordingToMap(t *testing.T) {
	g := graph.MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	theMap := make(graph.NeighborMap, 5)
	theMap[4] = graph.Neighbors{2, 3}
	theMap[1] = graph.Neighbors{5, 3}
	theMap[2] = graph.Neighbors{2, 5}

	removeAllVerticesAccordingToMap(g, theMap)
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
	g := graph.MkGraph(5)

	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)

	removeVertivesOfDegreeWithOnlyAdjacentNeighbors(g, 2)

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
	g := graph.MkGraph3()

	result, _ := getVerticesOfDegreeWithOnlyDisjointNeighbors(g, 2)
	assert.Equal(t, graph.Neighbors{2, 3}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])

	g = graph.MkGraph4()

	g.AddVertex()

	g.AddEdge(1, 8)
	g.AddEdge(2, 8)
	/*
	           1-----8
	          / \    |
	     3---+   +---2
	    / \         / \
	   7---6       5---4

	*/

	result, _ = getVerticesOfDegreeWithOnlyDisjointNeighbors(g, 3)
	assert.Equal(t, empty_neighborhood, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, empty_neighborhood, result[7])

	g = graph.MkGraph4()

	g.AddVertex()
	g.AddEdge(1, 8)
	/*
	           1-----8
	          / \
	     3---+   +---2
	    / \         / \
	   7---6       5---4

	*/

	result, _ = getVerticesOfDegreeWithOnlyDisjointNeighbors(g, 3)
	assert.Equal(t, graph.Neighbors{2, 3, 8}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, empty_neighborhood, result[7])

	// Edge case: neighbors of a vertex with degree of 1.
	result, _ = getVerticesOfDegreeWithOnlyDisjointNeighbors(g, 1)
	assert.Equal(t, empty_neighborhood, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, empty_neighborhood, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
	assert.Equal(t, graph.Neighbors{1}, result[7])

	g = graph.MkGraph5()
	result, _ = getVerticesOfDegreeWithOnlyDisjointNeighbors(g, 2)
	assert.Equal(t, graph.Neighbors{2, 3}, result[0])
	assert.Equal(t, empty_neighborhood, result[1])
	assert.Equal(t, empty_neighborhood, result[2])
	assert.Equal(t, empty_neighborhood, result[3])
	assert.Equal(t, empty_neighborhood, result[4])
	assert.Equal(t, graph.Neighbors{2, 7}, result[5])
	assert.Equal(t, empty_neighborhood, result[6])
}

func TestContractEdges(t *testing.T) {
	g := graph.MkGraph4()
	// ShowGraph(g)
	contractionMap := make(graph.NeighborMap, 1)
	contractionMap[0] = graph.Neighbors{2, 3}
	contractEdges(g, contractionMap)

	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))

	assert.True(t, g.HasEdge(1, 4))
	assert.True(t, g.HasEdge(1, 5))
	assert.True(t, g.HasEdge(1, 6))
	assert.True(t, g.HasEdge(1, 7))

	g = graph.MkGraph5()
	contractionMap = make(graph.NeighborMap, 6)
	contractionMap[0] = graph.Neighbors{2, 3}
	contractionMap[5] = graph.Neighbors{2, 7}
	contractEdges(g, contractionMap)

	assert.False(t, g.HasVertex(2))
	assert.False(t, g.HasVertex(3))
	assert.False(t, g.HasVertex(7))

	assert.True(t, g.HasEdge(1, 4))
	assert.True(t, g.HasEdge(1, 5))
	assert.True(t, g.HasEdge(6, 4))
	assert.True(t, g.HasEdge(6, 5))

	g = graph.MkGraph(7)
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

	cov := vc.BranchAndBound(g, nil, utility.MAX_INT)
	var folds mapset.Set
	folds = preprocessing4(g)

	cov2 := vc.BranchAndBound(g, nil, utility.MAX_INT)
	size := cov2.Cardinality()
	for foldInter := range folds.Iter() {
		fold := foldInter.(*fold)
		if cov2.Contains(fold.replacement) {
			size += 2
		}
	}

	assert.Equal(t, cov.Cardinality(), size)

	g = graph.MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	foldVertex(g, graph.Vertex(1))
	assert.Equal(t, 4, g.NEdges())
	assert.Equal(t, 4, g.Degree(graph.Vertex(8)))

	g = graph.MkGraph(13)
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

	foldVertex(g, graph.Vertex(1))
	foldVertex(g, graph.Vertex(2))
	foldVertex(g, graph.Vertex(3))

	assert.Equal(t, 5, g.NEdges())
	assert.Equal(t, 5, g.Degree(graph.Vertex(16)))
}

func TestPreprocessingMainRoutine(t *testing.T) {
	g := graph.MkGraph(11)
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

	cov := vc.BranchAndBound(g, nil, utility.MAX_INT)
	parameterReduction, folds := Preprocessing(g)
	cov2 := vc.BranchAndBound(g, nil, utility.MAX_INT)
	size := ComputeUnfoldedVertexCoverSize(folds, cov2) + parameterReduction
	assert.Equal(t, cov.Cardinality(), size)
}
