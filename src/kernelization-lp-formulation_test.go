package graph

import (
	"fmt"
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestObjectiveFunction(t *testing.T) {
	s1 := mapset.NewSet()
	s1.Add(Vertex(5))
	s1.Add(Vertex(6))

	s2 := mapset.NewSet()
	s2.Add(Vertex(5))

	assert.Equal(t, s2, objectiveFunction([]mapset.Set{s1, s2}))
}

func TestResolveConflict(t *testing.T) {
	g := MkGraph(2)
	n1 := Vertex(1)
	n2 := Vertex(2)

	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.degrees[n1.toInt()] = 4
	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.degrees[n2.toInt()] = g.degrees[n1.toInt()]
	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	g.degrees[n2.toInt()] = g.degrees[n1.toInt()] + 1
	assert.Equal(t, n2, resolveConflict(g, n1, n2))

}

func TestCalculateLowerBound(t *testing.T) {
	g := MkGraph(10)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(1, 6)
	g.AddEdge(8, 9)

	selection := mapset.NewSet()

	assert.Equal(t, 3, computeLowerBound(g, selection))

	selection.Add(Vertex(8))
	selection.Add(Vertex(9))
	selection.Add(Vertex(5))

	assert.Equal(t, 5, computeLowerBound(g, selection))
}

func TestgetEdgeEndpoints(t *testing.T) {
	g := MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	expected := make(Vertices, 3)
	expected[0] = 1
	expected[1] = 2
	expected[2] = 4
	assert.Equal(t, expected, g.getEdgeEndpoints())
}

func TestMkLpNode(t *testing.T) {
	g := MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	selection := mapset.NewSet()
	level := 1
	node := mkLpNode(g, selection, level)
	assert.NotNil(t, node)
	assert.Equal(t, selection, node.selection)
	assert.Equal(t, 2, node.lowerBound)
	assert.Equal(t, 1, node.level)
}

func TestGetNumberOfCoveredEdges(t *testing.T) {
	g := mkGraph1()
	s := mapset.NewSet()
	s.Add(Vertex(1))
	s.Add(Vertex(2))
	assert.Equal(t, 5, getNumberOfCoveredEdges(g, s))

	g = mkGraph6()
	s.Add(Vertex(4))
	s.Add(Vertex(5))
	assert.Equal(t, 7, getNumberOfCoveredEdges(g, s))
}

func TestBranchAndBound1(t *testing.T) {
	g := MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	optimalSelection := mapset.NewSet()
	optimalSelection.Add(Vertex(2))
	assert.True(t, optimalSelection.Equal(branchAndBound(g)))

	g = mkGraph6()
	optimalSelection = mapset.NewSet()
	optimalSelection.Add(Vertex(4))
	optimalSelection.Add(Vertex(5))
	assert.True(t, optimalSelection.Equal(branchAndBound(g)))

}

func TestBranchAndBound2(t *testing.T) {
	g := mkPetersenGraph()
	innerVertices, outerVertices := mapset.NewSet(), mapset.NewSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(Vertex(i))
		innerVertices.Add(Vertex(i + 5))
	}

	cover := branchAndBound(g)
	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))
}

func TestBranchAndBound3(t *testing.T) {
	g := mkReversePetersenGraph()
	innerVertices, outerVertices := mapset.NewSet(), mapset.NewSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(Vertex(i))
		innerVertices.Add(Vertex(i + 5))
	}

	cover := branchAndBound(g)

	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))
}
