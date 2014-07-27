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

func TestAddBipartiteEdges(t *testing.T) {
	originalEdges := []*Edge{
		0: &Edge{1, 2},
		1: &Edge{1, 3},
	}
	expected := make([]*Edge, 4)
	expected[0] = &Edge{1, 2}
	expected[1] = &Edge{1, 3}
	expected[2] = &Edge{4, 5}
	expected[3] = &Edge{4, 6}
	// inVerboseContext(func() {
	for i, actual := range addBipartiteEdges(originalEdges) {
		Debug("expected: %v, actual: %v", *expected[i], *actual)
		assert.Equal(t, *expected[i], *actual)
	}
	// })
}
