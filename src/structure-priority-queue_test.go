package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Analyze and create test cases for each structure type. @start-from-here

func TestStrong2TuplePriority(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	str := MkStructure(1, Vertex(1), Vertex(2))

	// 2 ≤ d ( u ) ≤ 3 and 2 ≤ d (v) ≤ 3
	assert.Equal(t, 1, str.computePriority(g))

	g.AddEdge(1, 6)
	g.AddEdge(1, 7)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	// d ( u ) ≥ 4 and d (v) ≥ 4
	assert.Equal(t, 1, str.computePriority(g))

	g.RemoveEdge(2, 3)
	g.RemoveEdge(2, 5)
	g.RemoveEdge(2, 6)
	// Does not fit the cases.
	assert.NotEqual(t, 1, str.computePriority(g))
}

func Test2TuplePriority(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	str := MkStructure(1, Vertex(1), Vertex(2))

	assert.Equal(t, 2, str.computePriority(g))
}

func Test_GoodPair_Du3_Deg5NeighborsWithoutCommonNeighbors(t *testing.T) {
	// u := 1
	// z := 2
	// The simplest case.
	g := MkGraph(16)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	g.AddEdge(2, 13)
	g.AddEdge(2, 14)
	g.AddEdge(2, 15)
	g.AddEdge(2, 16)

	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 9)
	g.AddEdge(4, 10)
	g.AddEdge(4, 11)
	g.AddEdge(4, 12)
	str := mkGoodPair(Vertex(1), Vertex(2))
	assert.Equal(t, 3, str.computePriority(g))

	// Neighbors share a common neighbor different than u.
	g = MkGraph(15)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	g.AddEdge(2, 13)
	g.AddEdge(2, 14)
	g.AddEdge(2, 15)
	g.AddEdge(2, 8)

	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 9)
	g.AddEdge(4, 10)
	g.AddEdge(4, 11)
	g.AddEdge(4, 12)
	assert.NotEqual(t, 3, str.computePriority(g))
}

func Tes_tGoodPair_Du3_DzGeq5(t *testing.T) {
	g := MkGraph(15)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	g.AddEdge(2, 13)
	g.AddEdge(2, 14)
	g.AddEdge(2, 15)
	g.AddEdge(2, 8)

	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 9)
	g.AddEdge(4, 10)
	g.AddEdge(4, 11)
	g.AddEdge(4, 12)
	str := mkGoodPair(Vertex(1), Vertex(2))
	assert.Equal(t, 4, str.computePriority(g))
}

func Test_GoodPair_Du3_DzGeq4(t *testing.T) {
	g := MkGraph(15)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	g.AddEdge(2, 13)
	g.AddEdge(2, 14)
	g.AddEdge(2, 15)

	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 9)
	g.AddEdge(4, 10)
	g.AddEdge(4, 11)
	g.AddEdge(4, 12)
	str := mkGoodPair(Vertex(1), Vertex(2))
	assert.Equal(t, 5, str.computePriority(g))
}

func Test_GoodPair_Du4_Deg5NeighborsGeq3_NeighborsHaveGeq1Edge(t *testing.T) {
	g := MkGraph(21)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(2, 10)
	g.AddEdge(2, 11)
	g.AddEdge(2, 12)
	g.AddEdge(2, 3)

	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 14)
	g.AddEdge(4, 15)
	g.AddEdge(4, 17)

	g.AddEdge(5, 18)
	g.AddEdge(5, 19)
	g.AddEdge(5, 20)
	g.AddEdge(5, 21)

	str := mkGoodPair(Vertex(1), Vertex(2))
	// TODO: There probably is a bug in neighborsOfUShareCommonVertexOtherThanU. @start-from-here
	// It manifests itself here in the way that the logic decides that the
	// neighbors do not share a neighbor other than u when this is clearly not
	// the case having the edge 2-3.
	// This does not affect the outcome of this test, because the case for
	// priority 6 because for this graph hasOnlyDegree5Neighbors is false and
	// the second condition for priority 7 does not have to be checked.

	// Write a test for that in structures_test.go.

	assert.Equal(t, 6, str.computePriority(g))
}

func Test_GoodPair_Du4_NeighborsWithoutCommonNeighbors(t *testing.T) {
	g := MkGraph(21)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(3, 9)

	g.AddEdge(2, 10)
	g.AddEdge(2, 11)
	g.AddEdge(2, 12)
	g.AddEdge(2, 13)

	g.AddEdge(4, 14)
	g.AddEdge(4, 15)
	g.AddEdge(4, 16)
	g.AddEdge(4, 17)

	g.AddEdge(5, 18)
	g.AddEdge(5, 19)
	g.AddEdge(5, 20)
	g.AddEdge(5, 21)

	str := mkGoodPair(Vertex(1), Vertex(2))
	assert.Equal(t, 7, str.computePriority(g))
}
