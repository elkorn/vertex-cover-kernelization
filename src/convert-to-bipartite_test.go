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
	for i, actual := range actual {
		Debug("expected: %v, actual: %v", *expected[i], *actual)
		assert.Equal(t, *expected[i], *actual)
	}
}

func TestAddBipartiteEdges(t *testing.T) {
	original := mkGraphWithVertices(3)
	original.AddEdge(1, 2)
	original.AddEdge(1, 3)

	expected := make([]*Edge, 4)
	expected[0] = &Edge{1, 2}
	expected[1] = &Edge{1, 3}
	expected[2] = &Edge{4, 5}
	expected[3] = &Edge{4, 6}
	addBipartiteEdges(original, 4)
	assertAllEdgesEqual(t, expected, original.Edges)
}

func TestMakeBipartite(t *testing.T) {
	g := mkGraphWithVertices(4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)

	expectedVertices := []Vertex{1, 2, 3, 4, 5, 6, 7, 8}
	expectedEdges := Edges{MkEdge(4, 1), MkEdge(2, 3), MkEdge(8, 5), MkEdge(6, 7)}

	inVerboseContext(func() {
		actual := makeBipartite(g)
		for _, v := range expectedVertices {
			assert.True(t, actual.Vertices[v])
		}

		assertAllEdgesEqual(t, expectedEdges, actual.Edges)
	})
}
