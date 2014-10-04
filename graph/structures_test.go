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

func TestNeighborsOfUShareCommonVertexOtherThanU(t *testing.T) {
	str := MkStructure(-1, Vertex(1), Vertex(2))

	g := MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	share, _ := str.neighborsOfUShareCommonVertexOtherThanU(Vertex(1), Vertex(2), g)
	assert.False(t, share)

	g.AddEdge(3, 5)

	share, _ = str.neighborsOfUShareCommonVertexOtherThanU(Vertex(1), Vertex(2), g)
	assert.True(t, share)

	g.RemoveEdge(3, 5)
	g.AddEdge(3, 6)
	g.AddEdge(4, 6)

	share, _ = str.neighborsOfUShareCommonVertexOtherThanU(Vertex(1), Vertex(2), g)
	assert.True(t, share)

	g = MkGraph(21)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(2, 10)
	g.AddEdge(2, 11)
	g.AddEdge(2, 12)

	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 14)
	g.AddEdge(4, 15)
	g.AddEdge(4, 17)

	g.AddEdge(5, 18)
	g.AddEdge(5, 19)
	g.AddEdge(5, 20)
	g.AddEdge(5, 21)

	share, _ = str.neighborsOfUShareCommonVertexOtherThanU(Vertex(1), Vertex(2), g)
	assert.False(t, share)

	// TODO: There was a bug in neighborsOfUShareCommonVertexOtherThanU.
	// It manifested itself here in the way that the logic decides that the
	// neighbors do not share a neighbor other than u when this is clearly not
	// the case having the edge 2-3.
	g = MkGraph(21)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)

	g.AddEdge(2, 10)
	g.AddEdge(2, 11)
	g.AddEdge(2, 12)
	g.AddEdge(2, 3)

	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)

	g.AddEdge(4, 14)
	g.AddEdge(4, 15)
	g.AddEdge(4, 17)

	g.AddEdge(5, 18)
	g.AddEdge(5, 19)
	g.AddEdge(5, 20)
	g.AddEdge(5, 21)

	share, _ = str.neighborsOfUShareCommonVertexOtherThanU(Vertex(1), Vertex(2), g)
	assert.True(t, share)
}

func TestIdentifyStructuresInProteins(t *testing.T) {
	// TODO: Devise better test cases.
	// For now, this should not fail.
	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
	identifyStructures(g, MAX_INT)
}
