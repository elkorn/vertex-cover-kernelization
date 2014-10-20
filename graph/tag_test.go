package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTag(t *testing.T) {
	g := MkGraph1()
	t1 := MkTag(Vertex(2), g)
	for i := 1; i < len(t1.neighbors); i++ {
		d1 := g.Degree(t1.neighbors[i-1])
		d2 := g.Degree(t1.neighbors[i])
		assert.True(t, d1 >= d2, fmt.Sprintf("deg(%v) [%v] must be greater than deg(%v) [%v]", t1.neighbors[i-1], d1, t1.neighbors[i], d2))
	}
}

func TestTagCompare(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 6)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(5, 7)

	t1, t2 := MkTag(Vertex(1), g), MkTag(Vertex(2), g)
	assert.Equal(t, 1, t1.Compare(t2, g))
	assert.Equal(t, -1, t2.Compare(t1, g))

	t1, t2 = MkTag(Vertex(3), g), MkTag(Vertex(6), g)
	assert.Equal(t, 0, t1.Compare(t2, g))

	g.AddEdge(6, 8)
	t1, t2 = MkTag(Vertex(3), g), MkTag(Vertex(6), g)
	assert.Equal(t, -1, t1.Compare(t2, g))
	assert.Equal(t, 1, t2.Compare(t1, g))
}
