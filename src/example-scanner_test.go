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

	g := ScanGraph("../examples/sh2-fake.dim.sh")

	for _, vertex := range expected.vertices {
		assert.True(t, g.hasVertex(vertex))
	}

	for _, edge := range expected.edges {
		assert.True(t, g.hasEdge(edge.from, edge.to), fmt.Sprintf("The resulting graph should contain an edge %v->%v", edge.from, edge.to))
	}

	ScanGraph("../examples/sh2/sh2-3.dim.sh")
}