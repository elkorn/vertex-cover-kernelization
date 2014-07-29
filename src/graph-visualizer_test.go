package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDot(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	expected := `graph test {
		1 -- 2;
		1 -- 3;
		2 -- 3;
	}`

	assert.Equal(t, expected, g.ToDot("test"))
}
