package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDominates(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	v, u := Vertex(1), Vertex(2)

	assert.True(t, v.dominates(u, g))
}
