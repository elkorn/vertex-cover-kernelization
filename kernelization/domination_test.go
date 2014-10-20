package kernelization

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDominates(t *testing.T) {
	g := graph.MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	v, u := Vertex(1), Vertex(2)
	assert.True(t, v.dominates(u, g))
	g.RemoveEdge(1, 2)
	// There is no edge.
	assert.False(t, v.dominates(u, g))
	g.AddEdge(1, 2)
	g.RemoveEdge(1, 3)
	// 3 is not in common neighborhood.
	assert.False(t, Vertex(1).dominates(Vertex(2), g))

}

func TestAlmostDominates(t *testing.T) {
	g := graph.MkGraph(5)
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
