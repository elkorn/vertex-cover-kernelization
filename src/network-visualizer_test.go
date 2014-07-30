package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToGraph(t *testing.T) {
	expected := mkGraphWithVertices(4)
	expected.AddEdge(1, 2)
	expected.AddEdge(1, 3)
	expected.AddEdge(2, 4)

	net := mkNet(expected)
	actual := convertToGraph(net)
	for vertex := range expected.Vertices {
		assert.True(t, actual.hasVertex(vertex))
	}

	for _, edge := range expected.Edges {
		assert.True(t, actual.hasEdge(edge.from, edge.to))
	}
}
