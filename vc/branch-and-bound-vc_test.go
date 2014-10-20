package vc

import (
	"fmt"
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestObjectiveFunction(t *testing.T) {
	s1 := mapset.NewSet()
	s1.Add(graph.Vertex(5))
	s1.Add(graph.Vertex(6))

	s2 := mapset.NewSet()
	s2.Add(graph.Vertex(5))

	assert.Equal(t, s2, objectiveFunction([]mapset.Set{s1, s2}))
}

func TestResolveConflict(t *testing.T) {
	g := graph.MkGraph(4)

	n1 := graph.Vertex(1)
	n2 := graph.Vertex(2)

	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.RemoveEdge(1, 3)
	assert.Equal(t, n2, resolveConflict(g, n1, n2))

}

func TestCalculateLowerBound(t *testing.T) {
	g := graph.MkGraph(10)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(1, 6)
	g.AddEdge(8, 9)

	selection := mapset.NewSet()

	assert.Equal(t, 3, computeLowerBound(g, selection))

	selection.Add(graph.Vertex(8))
	selection.Add(graph.Vertex(9))
	selection.Add(graph.Vertex(5))

	assert.Equal(t, 5, computeLowerBound(g, selection))
}

func TestgetEdgeEndpoints(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	expected := make(graph.Vertices, 3)
	expected[0] = 1
	expected[1] = 2
	expected[2] = 4
	assert.Equal(t, expected, getEdgeEndpoints(g))
}

func TestMkBnbNode(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	selection := mapset.NewSet()
	level := 1
	node := mkBnbNode(g, selection, level)
	assert.NotNil(t, node)
	assert.Equal(t, selection, node.selection)
	assert.Equal(t, 2, node.lowerBound)
	assert.Equal(t, 1, node.level)
}

func TestGetNumberOfCoveredEdges(t *testing.T) {
	g := graph.MkGraph1()
	s := mapset.NewSet()
	s.Add(graph.Vertex(1))
	s.Add(graph.Vertex(2))
	assert.Equal(t, 5, getNumberOfCoveredEdges(g, s))

	g = graph.MkGraph6()
	s.Add(graph.Vertex(4))
	s.Add(graph.Vertex(5))
	assert.Equal(t, 7, getNumberOfCoveredEdges(g, s))
}

func TestBranchAndBound1(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	optimalSelection := mapset.NewSet()
	optimalSelection.Add(graph.Vertex(2))
	cover := branchAndBound(g)
	assert.True(t, optimalSelection.Equal(cover))

	g = graph.MkGraph6()
	optimalSelection = mapset.NewSet()
	optimalSelection.Add(graph.Vertex(4))
	optimalSelection.Add(graph.Vertex(5))
	assert.True(t, optimalSelection.Equal(branchAndBound(g)))

}

func TestBranchAndBound2(t *testing.T) {
	g := graph.MkPetersenGraph()
	innerVertices, outerVertices := mapset.NewSet(), mapset.NewSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(graph.Vertex(i))
		innerVertices.Add(graph.Vertex(i + 5))
	}

	cover := branchAndBound(g)
	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))

}

func TestBranchAndBound3(t *testing.T) {
	g := graph.MkReversePetersenGraph()
	innerVertices, outerVertices := mapset.NewSet(), mapset.NewSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(graph.Vertex(i))
		innerVertices.Add(graph.Vertex(i + 5))
	}

	cover := branchAndBound(g)

	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))
}

func TestBnBBipartite(t *testing.T) {
	size := 7
	g := graph.MkGraph(2 * size)
	for i := 1; i <= size; i++ {
		for j := size + 1; j <= 2*size; j++ {
			g.AddEdge(graph.Vertex(i), graph.Vertex(j))
		}
	}

	cover := branchAndBound(g)
	assert.Equal(t, size, cover.Cardinality())

}
