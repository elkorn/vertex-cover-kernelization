package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
