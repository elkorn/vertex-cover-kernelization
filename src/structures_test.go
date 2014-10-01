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
	gpi := mkGoodPair(Vertex(1))
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
	gpi := mkGoodPair(Vertex(2))
	assert.Equal(t, 5, gpi.countNeighborhoodEdges(g))
}

func TestIdentifyGoodPairs(t *testing.T) {
	// Testing u.
	g := MkGraph(13)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	g.AddEdge(2, 7)
	g.AddEdge(2, 8)

	g.AddEdge(4, 5)
	g.AddEdge(4, 6)

	g.AddEdge(9, 10)
	g.AddEdge(9, 11)
	g.AddEdge(9, 12)
	g.AddEdge(10, 9)
	g.AddEdge(10, 13)
	g.AddEdge(12, 10)
	pairs := identifyGoodPairs(g)
	for p := range pairs.Iter() {
		pp := p.(*goodPair)
		Debug("Good pair with u: %v, z: %v, domination: %v, edges: %v", pp.U(), pp.Z(), pp.numNeighborhoodAlmostDominatedPairs, pp.numNeighborhoodEdges)
	}

	assert.Equal(t, 2, pairs.Cardinality())
	p := pairs.Iter()
	pp := (<-p).(*goodPair)
	assert.Equal(t, Vertex(9), pp.U())
	assert.Equal(t, Vertex(12), pp.Z())
	assert.Equal(t, 3, pp.numNeighborhoodAlmostDominatedPairs)
	assert.Equal(t, 1, pp.numNeighborhoodEdges)
	pp = (<-p).(*goodPair)
	assert.Equal(t, Vertex(10), pp.U())
	assert.Equal(t, Vertex(12), pp.Z())
	assert.Equal(t, 3, pp.numNeighborhoodAlmostDominatedPairs)
	assert.Equal(t, 1, pp.numNeighborhoodEdges)
}
