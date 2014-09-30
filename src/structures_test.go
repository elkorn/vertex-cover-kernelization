package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAlmostDominatedPairs(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(2, 7)
	g.AddEdge(2, 8)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	gpi := mkGoodPairInfo(MkGoodPair(Vertex(1)))
	assert.Equal(t, 2, gpi.countAlmostDominatedPairs(g))
}

func TestCountNeighborhoodEdges(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 3)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(2, 7)
	g.AddEdge(2, 8)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	gpi := mkGoodPairInfo(MkGoodPair(Vertex(2)))
	assert.Equal(t, 5, gpi.countNeighborhoodEdges(g))
}
