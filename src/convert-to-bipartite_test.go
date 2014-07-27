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
