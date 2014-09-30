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
	g.AddEdge(10, 13)
	g.AddEdge(12, 10)
	tags := computeTags(g)
	inVerboseContext(func() {
		pairs := identifyGoodPairs(g)
		for p := range pairs.Iter() {
			pp := p.(*goodPair)
			Debug("Good pair with u: %v, z: %v, domination: %v, edges: %v", pp.U(), pp.Z(), pp.numNeighborhoodAlmostDominatedPairs, pp.numNeighborhoodEdges)
			Debug("Tag: %v", tags[pp.U().toInt()])
		}
	})
}
