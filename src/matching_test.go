package graph

import (
	"testing"

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
