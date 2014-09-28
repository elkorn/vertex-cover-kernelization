package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkGraph(t *testing.T) {
	g := MkGraph(0)
	assert.Equal(t, len(g.Edges), 0)
	assert.Equal(t, g.NVertices(), 0)
}

func TestaddVertex(t *testing.T) {
	g := MkGraph(0)
	g.addVertex()
	assert.True(t, g.hasVertex(MkVertex(0)))
}

func TestRemoveVertex(t *testing.T) {
	g := MkGraph(3)

	err := g.RemoveVertex(2)
	assert.Nil(t, err)
	assert.False(t, g.hasVertex(2))
	g.addVertex()
	g.AddEdge(1, 4)
	g.AddEdge(1, 3)
	g.AddEdge(4, 3)

	err = g.RemoveVertex(4)
	assert.Nil(t, err)
	assert.Equal(t, false, g.hasEdge(1, 4))
	assert.Equal(t, false, g.hasEdge(4, 3))
	assert.Equal(t, true, g.hasEdge(1, 3))
}

func TestAddEdge(t *testing.T) {
	g := MkGraph(0)
	g.addVertex()
	g.addVertex()
	g.addVertex()
	err := g.AddEdge(1, 2)
	assert.Nil(t, err)

	edge := g.Edges[0]
	assert.Equal(t, edge.from, 1)
	assert.Equal(t, edge.to, 2)

	err = g.AddEdge(2, 1)
	assert.NotNil(t, err)
	err = g.AddEdge(1, 1)
	assert.NotNil(t, err)

	err = g.AddEdge(3, 1)
	edge = g.Edges[1]
	assert.Equal(t, edge.from, 3)
	assert.Equal(t, edge.to, 1)
}

func TestVertexCoverSimpleGraph(t *testing.T) {
	g := MkGraph(0)
	g.addVertex()
	g.addVertex()
	g.addVertex()

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	assert.Equal(t, g.IsVertexCover(2), true)
	assert.Equal(t, g.IsVertexCover(3), false)
	assert.Equal(t, g.IsVertexCover(1, 3), true)
}

func TestVertexCoverNonTrivialGraph1(t *testing.T) {
	g := mkGraph1()
	assert.Equal(t, g.IsVertexCover(5), false)
	assert.Equal(t, g.IsVertexCover(1, 3, 5), true)
}

func TestVertexCoverNonTrivialGraph2(t *testing.T) {
	g := mkGraph2()
	assert.Equal(t, g.IsVertexCover(2, 3, 4, 5, 7), true)
}

func TestVertexDegree(t *testing.T) {
	g := MkGraph(5)
	assertDegreeIsCorrect := func(v Vertex, expectedDegree int) {
		assert.Equal(t, expectedDegree, g.Degree(v))
	}

	for i := 1; i <= 5; i++ {
		assertDegreeIsCorrect(Vertex(i), 0)
	}

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)
	assertDegreeIsCorrect(Vertex(2), 4)

	g.AddEdge(1, 5)
	assertDegreeIsCorrect(Vertex(5), 2)
	assertDegreeIsCorrect(Vertex(1), 2)
}

func TestGetNeighbors(t *testing.T) {
	g := MkGraph(5)

	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)
	g.AddEdge(5, 3)

	assert.Equal(t, Neighbors{1, 2, 4, 5}, g.getNeighbors(3))
	assert.Equal(t, Neighbors{3}, g.getNeighbors(1))

	g = MkGraph(8)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	g.AddEdge(1, 8)
	g.AddEdge(2, 8)

	assert.Equal(t, Neighbors{2, 3, 8}, g.getNeighbors(1))
}

func TestaddVertexWithAutoId(t *testing.T) {
	g := MkGraph(12)
	assert.Equal(t, Vertex(13), g.generateVertex())
}

func TestRegularity(t *testing.T) {
	g := MkGraph(4)
	assert.True(t, g.IsRegular())
	g.AddEdge(1, 2)
	assert.False(t, g.IsRegular())
	g.AddEdge(3, 4)
	assert.True(t, g.IsRegular())
	g.AddEdge(1, 4)
	assert.False(t, g.IsRegular())
	g.AddEdge(3, 2)
	assert.True(t, g.IsRegular())
}
