package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVertices(t *testing.T) {
	g := mkGraphWithVertices(3)

	expected := make([]Vertex, 6)
	for i := 0; i < 6; i++ {
		expected[i] = Vertex(i + 1)
	}

	actual := getVertices(g)
	assert.Equal(t, expected, actual)
}

func assertAllEdgesEqual(t *testing.T, expected, actual []*Edge) {
	assert.Equal(t, len(expected), len(actual), "The number of edges must be the same.")
	for i, actual := range actual {
		Debug("expected: %v, actual: %v", *expected[i], *actual)
		assert.Equal(t, *expected[i], *actual)
	}
}

func TestMakeBipartite(t *testing.T) {
	g := mkGraphWithVertices(4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)

	expectedVertices := []Vertex{1, 2, 3, 4, 5, 6, 7, 8}
	expectedEdges := Edges{
		MkEdge(4, 5),
		MkEdge(2, 7),
	}

	actual := makeBipartite(g)
	for _, v := range expectedVertices {
		assert.True(t, actual.Vertices[v])
	}

	assertAllEdgesEqual(t, expectedEdges, actual.Edges)
}
