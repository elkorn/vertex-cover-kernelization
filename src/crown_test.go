package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const k = MAX_INT

func TestFindCrown1(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(8, 7)
	halt := make(chan bool, 1)

	crown := findCrown(g, halt, k)
	assert.Equal(t, 1, crown.Width())
	assert.True(t, crown.H.Contains(Vertex(4)))
	assert.True(t, crown.I.Contains(Vertex(1)))
	assert.True(t, crown.I.Contains(Vertex(2)))
	assert.True(t, crown.I.Contains(Vertex(3)))
	// gv := MkGraphVisualizer(g)
	// gv.highlightCrown(crown)
	// gv.Display()
}

func TestReduceCrown1(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(8, 7)
	halt := make(chan bool, 1)
	ReduceCrown(g, halt, k)
	assert.Equal(t, 4, g.NVertices())
	assert.True(t, g.hasVertex(5))
	assert.True(t, g.hasVertex(6))
	assert.True(t, g.hasVertex(7))
	assert.True(t, g.hasVertex(8))
	assert.Equal(t, 2, g.NEdges())
	assert.True(t, g.hasEdge(5, 6))
	assert.True(t, g.hasEdge(7, 8))
}

func TestReduceCrown2(t *testing.T) {
	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
	halt := make(chan bool, 1)
	verticesBefore := g.NVertices()
	// crown := findCrown(g, halt, k)
	// crownWidth, independentSetCardinality := crown.Width(), crown.I.Cardinality()
	crownWidth, independentSetCardinality := 228, 542

	ReduceCrown(g, halt, k)
	assert.Equal(t,
		verticesBefore-crownWidth-independentSetCardinality,
		g.NVertices())
}

func TestStopIfSizeBoundaryReached(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(8, 7)

	halt := make(chan bool, 1)

	assert.Nil(t, findCrown(g, halt, 1))
}

// func TestFindCrownBigGraph(t *testing.T) {
// 	// This test should not crash. If it does, the algorithm is broken.
// 	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
// 	halt := make(chan bool, 1)
// 	findCrown(g, halt, MAX_INT)
// }
