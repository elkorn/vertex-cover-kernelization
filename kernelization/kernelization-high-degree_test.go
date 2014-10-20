package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/vc"
	"github.com/stretchr/testify/assert"
)

func TestRemoveHighDegree(t *testing.T) {
	g1 := graph.MkGraph(10)

	g1.AddEdge(1, 2)
	g1.AddEdge(1, 3)
	g1.AddEdge(1, 4)
	g1.AddEdge(1, 5)

	g1.AddEdge(2, 3)
	g1.AddEdge(2, 4)
	g1.AddEdge(2, 5)
	g1.AddEdge(2, 6)
	g1.AddEdge(2, 7)

	g1.AddEdge(9, 8)

	vc := vc.BranchAndBound(g1)
	removed, remCount := g1.removeVerticesWithDegreeGreaterThan(2)
	vc2 := vc.BranchAndBound(g1)
	assert.True(t, Contains(removed, 1))
	assert.True(t, Contains(removed, 2))

	assert.False(t, g1.HasVertex(1))
	assert.False(t, g1.HasVertex(2))

	assert.False(t, g1.HasEdge(1, 2))
	assert.False(t, g1.HasEdge(1, 3))
	assert.False(t, g1.HasEdge(1, 4))
	assert.False(t, g1.HasEdge(1, 5))

	assert.False(t, g1.HasEdge(2, 3))
	assert.False(t, g1.HasEdge(2, 4))
	assert.False(t, g1.HasEdge(2, 5))
	assert.False(t, g1.HasEdge(2, 6))
	assert.False(t, g1.HasEdge(2, 7))

	assert.True(t, g1.HasEdge(9, 8))

	assert.Equal(t, vc.Cardinality(), vc2.Cardinality()+remCount)
}
