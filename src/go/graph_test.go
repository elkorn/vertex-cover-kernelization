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

func TestVertexCover(t *testing.T) {
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
