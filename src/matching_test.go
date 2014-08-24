package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalMatching(t *testing.T) {
	g := mkGraph1()

	m, o := maximalMatching(g)
	assert.Equal(t, 2, m.Cardinality(), "The matching of graph1 should contain only 2 edges.")
	assert.Equal(t, 6, o.Cardinality(), "The outsiders of graph1 should contain all unmatched edges.")
	assert.True(t, m.Contains(g.getEdgeByCoordinates(0, 1)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(2, 3)))
}
