package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mkGraphWithVertices(howMany int) *Graph {
	g := MkGraph()
	for i := 1; i <= howMany; i++ {
		g.AddVertex(Vertex(i))
	}

	return g
}

func TestMkGraph(t *testing.T) {
	g := MkGraph()
	assert.Equal(t, len(g.Edges), 0)
	assert.Equal(t, len(g.Vertices), 0)
}

func TestAddVertex(t *testing.T) {
	g := MkGraph()
	err := g.AddVertex(Vertex(1))
	assert.Nil(t, err)
	assert.Equal(t, g.Vertices[1], true)

	err = g.AddVertex(1)
	assert.NotNil(t, err)
}

func TestRemoveVertex(t *testing.T) {
	g := mkGraphWithVertices(3)

	err := g.RemoveVertex(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(g.Vertices))

	g.AddVertex(2)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	err = g.RemoveVertex(2)
	assert.Nil(t, err)
	assert.Equal(t, false, g.hasEdge(1, 2))
	assert.Equal(t, false, g.hasEdge(2, 3))
	assert.Equal(t, true, g.hasEdge(1, 3))
}

func TestAddEdge(t *testing.T) {
	g := MkGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
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
	g := MkGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	assert.Equal(t, g.IsVertexCover(2), true)
	assert.Equal(t, g.IsVertexCover(3), false)
	assert.Equal(t, g.IsVertexCover(1, 3), true)
}

func TestVertexCoverNonTrivialGraph1(t *testing.T) {
	/*
		   1o---o2
			|\ /|
			| o5|
			|/ \|
		   4o---o3
	*/
	g := MkGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	assert.Equal(t, g.IsVertexCover(5), false)
	assert.Equal(t, g.IsVertexCover(1, 3, 5), true)
}

func TestVertexCoverNonTrivialGraph2(t *testing.T) {
	/*

		   1o--------o2
			|\      /|
			|5o----o6|
			| |    | |
			|8o----o7|
			|/      \|
		   4o--------o3
	*/
	g := MkGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 8)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(8, 5)

	assert.Equal(t, g.IsVertexCover(2, 3, 4, 5, 7), true)
}

func TestVertexDegree(t *testing.T) {
	g := MkGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 2)
	g.AddEdge(5, 2)
	degree, err := g.Degree(2)
	assert.Nil(t, err)
	assert.Equal(t, degree, 4)

	degree, err = g.Degree(1)
	assert.Nil(t, err)
	assert.Equal(t, degree, 1)
}

func TestGetNeighbors(t *testing.T) {
	g := mkGraphWithVertices(5)

	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(4, 3)
	g.AddEdge(5, 3)

	assert.Equal(t, Neighbors{1, 2, 4, 5}, g.getNeighbors(3))
	assert.Equal(t, Neighbors{3}, g.getNeighbors(1))

	g = mkGraphWithVertices(8)

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
