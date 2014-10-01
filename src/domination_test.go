package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDominates(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	v, u := Vertex(1), Vertex(2)

	assert.True(t, v.dominates(u, g))
}

func TestAlmostDominates(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(5, 2)
	g.AddEdge(5, 3)
	g.AddEdge(5, 4)

	assert.True(t, Vertex(1).almostDominates(Vertex(5), g))
	assert.True(t, Vertex(5).almostDominates(Vertex(1), g))
	g.RemoveEdge(2, 5)
	assert.True(t, Vertex(1).almostDominates(Vertex(5), g))
	assert.True(t, Vertex(5).almostDominates(Vertex(1), g))
	g.RemoveEdge(3, 5)
	assert.True(t, Vertex(1).almostDominates(Vertex(5), g))
	assert.False(t, Vertex(5).almostDominates(Vertex(1), g))
}
