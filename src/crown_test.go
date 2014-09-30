package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const k = MAX_INT

func TestFindCrown1(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	halt := make(chan bool, 1)

	gv := MkGraphVisualizer(g)
	m, o := FindMaximalMatching(g)
	gv.HighlightMatching(m, "red")
	gv.HighlightCover(o, "green")
	// inVerboseContext(func() {
	// 	findCrown(g, halt, k)
	// })
	crown := findCrown(g, halt, k)
	assert.Equal(t, 1, crown.Width())
	assert.True(t, crown.H.Contains(Vertex(4)))
	assert.True(t, crown.I.Contains(Vertex(2)))
	assert.True(t, crown.I.Contains(Vertex(3)))
	gv.highlightCrown(crown)
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
	assert.Equal(t, 5, g.NVertices())
	assert.True(t, g.hasVertex(1))
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
	crown := findCrown(g, halt, 246)
	crownWidth, independentSetCardinality := crown.Width(), crown.I.Cardinality()
	ReduceCrown(g, halt, 246)
	assert.Equal(t,
		verticesBefore-crownWidth-independentSetCardinality,
		g.NVertices())
}

func TestReduceProteins(t *testing.T) {
	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
	halt := make(chan bool, 1)

	// Test according to F.N.Abu-Khzam et al. paper. (Table 1)
	kPrimePrev, kPrime := 0, 246

	for {
		kPrime, _ = ReduceCrown(g, halt, kPrime)
		if kPrimePrev == kPrime {
			break
		}

		kPrimePrev = kPrime
	}

	assert.Equal(t, 99, kPrime)
	assert.Equal(t, 243, g.NVertices())
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

func TestFindCrownBug(t *testing.T) {
	g := MkGraph(10)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)
	g.AddEdge(6, 7)
	g.AddEdge(4, 5)
	g.AddEdge(4, 9)
	g.AddEdge(4, 10)

	h := make(chan bool, 1)
	crown := findCrown(g, h, MAX_INT)
	reduceCrown(g, crown)
	matching := FindMaximumMatching(g)
	Debug("MAX MATCH: %v", matching.Edges)
}
