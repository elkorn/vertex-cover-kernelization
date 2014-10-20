package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
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
	v, u := graph.Vertex(1), graph.Vertex(2)
	assert.True(t, dominates(v, u, g))
	g.RemoveEdge(1, 2)
	// There is no edge.
	assert.False(t, dominates(v, u, g))
	g.AddEdge(1, 2)
	g.RemoveEdge(1, 3)
	// 3 is not in common neighborhood.
	assert.False(t, dominates(graph.Vertex(1), graph.Vertex(2), g))

}

func TestAlmostDominates(t *testing.T) {
	g := graph.MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(5, 2)
	g.AddEdge(5, 3)
	g.AddEdge(5, 4)

	assert.True(t, almostDominates(graph.Vertex(1), graph.Vertex(5), g))
	assert.True(t, almostDominates(graph.Vertex(5), graph.Vertex(1), g))
	g.RemoveEdge(2, 5)
	assert.True(t, almostDominates(graph.Vertex(1), graph.Vertex(5), g))
	assert.True(t, almostDominates(graph.Vertex(5), graph.Vertex(1), g))
	g.RemoveEdge(3, 5)
	assert.True(t, almostDominates(graph.Vertex(1), graph.Vertex(5), g))
	assert.False(t, almostDominates(graph.Vertex(5), graph.Vertex(1), g))
}
