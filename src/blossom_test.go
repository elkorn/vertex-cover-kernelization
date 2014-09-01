package graph

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestContractBlossomSketch(t *testing.T) {
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

func TestContractBlossom1(t *testing.T) {
	g := MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)

	b := blossom{
		Root:     MkVertex(1),
		edges:    mapset.NewSet(),
		vertices: mapset.NewSet(),
	}

	g.AddEdge(2, 6)
	g.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if edge.from == 1 {
			return
		}

		b.edges.Add(edge)
		b.vertices.Add(edge.from)
		b.vertices.Add(edge.to)
	})

	b.Contract(g)
	assert.Equal(t, 1, g.NEdges())
	assert.Equal(t, 2, g.NVertices())
	assert.True(t, g.hasEdge(1, 2))
}
func TestContractBlossom2(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)
	g.AddEdge(2, 6)

	b := blossom{
		Root:     MkVertex(1),
		edges:    mapset.NewSet(),
		vertices: mapset.NewSet(),
	}

	g.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if edge.from == 1 || edge.to == 7 {
			return
		}

		b.edges.Add(edge)
		b.vertices.Add(edge.from)
		b.vertices.Add(edge.to)
	})

	b.Contract(g)
	assert.Equal(t, 2, g.NEdges())
	assert.Equal(t, 3, g.NVertices())
	assert.True(t, g.hasEdge(1, 2))
	assert.True(t, g.hasEdge(2, 7))

}

func TestContractBlossom3(t *testing.T) {
	g := MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(8, 2)
	g.AddEdge(9, 5)

	b := blossom{
		Root:     MkVertex(1),
		edges:    mapset.NewSet(),
		vertices: mapset.NewSet(),
	}

	g.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if edge.from == 1 || edge.from == 9 {
			return
		}

		b.edges.Add(edge)
		b.vertices.Add(edge.from)
		b.vertices.Add(edge.to)
	})

	gv := MkGraphVisualizer()
	gv.Display(g)
	inVerboseContext(func() {
		b.Contract(g)
	})

	assert.Equal(t, 2, g.NEdges())
	assert.Equal(t, 3, g.NVertices())
	assert.True(t, g.hasEdge(1, 2))
	assert.True(t, g.hasEdge(2, 9))

	gv.Display(g)
}
