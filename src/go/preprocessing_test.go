package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveHighDegree(t *testing.T) {
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

	RemoveVerticesOfDegree(g, 4)
	assert.Equal(t, 4, len(g.Vertices))
	assert.Equal(t, true, g.hasVertex(1))
	assert.Equal(t, false, g.hasVertex(2))
	assert.Equal(t, true, g.hasVertex(3))
	assert.Equal(t, true, g.hasVertex(4))
	assert.Equal(t, true, g.hasVertex(5))

	assert.Equal(t, false, g.hasEdge(1, 2))
	assert.Equal(t, false, g.hasEdge(2, 3))
	assert.Equal(t, false, g.hasEdge(4, 2))
	assert.Equal(t, false, g.hasEdge(5, 2))
}
