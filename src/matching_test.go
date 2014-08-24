package graph

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestMaximalMatching(t *testing.T) {
	g := mkGraph1()

	m, o := maximalMatching(g)
	assert.Equal(t, 2, m.Cardinality(), "The matching of graph1 should contain 2 edges.")
	assert.Equal(t, 6, o.Cardinality(), "The outsiders of graph1 should contain 6 unmatched edges.")
	assert.True(t, m.Contains(g.getEdgeByCoordinates(0, 1)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(2, 3)))

	g = mkGraph5()
	m, o = maximalMatching(g)
	assert.Equal(t, 3, m.Cardinality(), "The matching of graph5 should contain 3 edges.")
	assert.Equal(t, 4, o.Cardinality(), "The outsiders of graph5 should contain 4 unmatched edges.")
	assert.True(t, m.Contains(g.getEdgeByCoordinates(0, 1)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(3, 4)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(5, 6)))
}

func TestIsExposed(t *testing.T) {
	matching := mapset.NewSet()
	matching.Add(MkEdge(1, 2))
	matching.Add(MkEdge(2, 3))
	matching.Add(MkEdge(4, 5))
	matching.Add(MkEdge(7, 8))

	assert.False(t, Vertex(1).isExposed(matching))
	assert.False(t, Vertex(2).isExposed(matching))
	assert.False(t, Vertex(3).isExposed(matching))
	assert.False(t, Vertex(4).isExposed(matching))
	assert.False(t, Vertex(5).isExposed(matching))
	assert.False(t, Vertex(7).isExposed(matching))
	assert.False(t, Vertex(8).isExposed(matching))

	assert.True(t, Vertex(6).isExposed(matching))
	assert.True(t, Vertex(9).isExposed(matching))
}
