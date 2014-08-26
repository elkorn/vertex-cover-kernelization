package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContractBlossom(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	contractionMap := make(NeighborMap, 4)
	// Starting from the root
	// keep contracting the vertices belonging to the blossom
	// to the root.
	contractionMap[0] = g.getNeighbors(1)
	g.contractEdges(contractionMap)
	contractionMap[0] = g.getNeighbors(1)
	g.contractEdges(contractionMap)
	assert.Equal(t, 1, g.NVertices())
	assert.Equal(t, 0, g.NEdges())
}
