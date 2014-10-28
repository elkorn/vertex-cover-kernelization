package graph

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanGraph(t *testing.T) {
	expected := struct {
		vertices []Vertex
		edges    Edges
	}{
		vertices: []Vertex{1, 2, 3, 4, 5, 6, 7},
		edges:    Edges{MkEdgeFromInts(0, 1), MkEdgeFromInts(0, 2), MkEdgeFromInts(0, 3), MkEdgeFromInts(1, 2), MkEdgeFromInts(1, 4), MkEdgeFromInts(1, 5), MkEdgeFromInts(2, 3), MkEdgeFromInts(2, 6), MkEdgeFromInts(3, 6), MkEdgeFromInts(3, 5), MkEdgeFromInts(4, 5), MkEdgeFromInts(4, 0)},
	}

	g := ScanGraph("../examples/sh2-fake.dim")

	for _, vertex := range expected.vertices {
		assert.True(t, g.HasVertex(vertex))
	}

	for _, edge := range expected.edges {
		assert.True(t, g.HasEdge(edge.From, edge.To), fmt.Sprintf("The resulting graph should contain an edge %v->%v", edge.From, edge.To))
	}
}

func TestScanDot(t *testing.T) {
	g := ScanDot("./test.dot")
	assert.Equal(t, 7, g.NVertices())
	assert.Equal(t, 3, g.NEdges())
	assert.True(t, g.HasEdge(2, 3))
	assert.True(t, g.HasEdge(1, 6))
	assert.True(t, g.HasEdge(5, 7))
	g = ScanDot("./test-big.dot")
	assert.Equal(t, 168, g.NEdges())
}
