package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStruction(t *testing.T) {
	g := MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5) // 5 is t on the diagram
	g.AddEdge(2, 6) // 6 is w on the diagram
	g.AddEdge(3, 7) // 7 is x on the diagram
	g.AddEdge(4, 8) // 8 is y on the diagram
	g.AddEdge(4, 9) // 9 is z on the diagram

	g1 := struction(g, Vertex(1))
	assert.Equal(t, 7, g1.NVertices())
	assert.Equal(t, 8, g1.NEdges())
	assert.True(t, g1.hasVertex(Vertex(10)))
	assert.True(t, g1.hasVertex(Vertex(11)))
	assert.False(t, g1.hasVertex(Vertex(1)))
	assert.False(t, g1.hasVertex(Vertex(2)))
	assert.False(t, g1.hasVertex(Vertex(3)))
	assert.False(t, g1.hasVertex(Vertex(4)))
	assert.True(t, g1.hasEdge(11, 10))
	assert.True(t, g1.hasEdge(11, 7))
	assert.True(t, g1.hasEdge(11, 8))
	assert.True(t, g1.hasEdge(11, 9))
	assert.True(t, g1.hasEdge(10, 5))
	assert.True(t, g1.hasEdge(10, 6))
	assert.True(t, g1.hasEdge(10, 8))
	assert.True(t, g1.hasEdge(10, 9))
}

func TestStruction2(t *testing.T) {
	g := MkGraph(12)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)

	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(4, 9)
	g.AddEdge(5, 10)
	g.AddEdge(5, 11)
	g.AddEdge(5, 12)

	g1 := struction(g, Vertex(1))
	assert.Equal(t, 10, g1.NVertices())
	assert.Equal(t, 15, g1.NEdges())
	assert.Equal(t, 6, g1.Degree(13))
	assert.Equal(t, 5, g1.Degree(14))
	assert.Equal(t, 7, g1.Degree(15))
	assert.True(t, g1.hasEdge(13, 14))
	assert.True(t, g1.hasEdge(13, 15))
	assert.True(t, g1.hasEdge(13, 6))
	assert.True(t, g1.hasEdge(13, 10))
	assert.True(t, g1.hasEdge(13, 11))
	assert.True(t, g1.hasEdge(13, 12))
	assert.True(t, g1.hasEdge(13, 12))

	assert.True(t, g1.hasEdge(14, 15))
	assert.True(t, g1.hasEdge(14, 7))
	assert.True(t, g1.hasEdge(14, 8))
	assert.True(t, g1.hasEdge(14, 9))

	assert.True(t, g1.hasEdge(15, 7))
	assert.True(t, g1.hasEdge(15, 8))
	assert.True(t, g1.hasEdge(15, 10))
	assert.True(t, g1.hasEdge(15, 11))
	assert.True(t, g1.hasEdge(15, 12))
}

func TestReduceAlmostCrown(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	g.AddEdge(4, 5)
	g.AddEdge(6, 7)

	reduceAlmostCrown(g, nil, MAX_INT)
}

func TestGeneralFold2(t *testing.T) {
	g := MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 8)
	g.AddEdge(8, 9)
	g.AddEdge(8, 2)
	g.AddEdge(8, 3)
	g.AddEdge(9, 2)
	g.AddEdge(9, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	generalFold(g, nil, 190)
	// g1 :=  generalFold(g, 0)

}

func TestGeneralFold3(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	// g1 :=  generalFold(g)
}

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

func TestTag(t *testing.T) {
	g := mkGraph1()
	t1 := MkTag(Vertex(2), g)
	for i := 1; i < len(t1.neighbors); i++ {
		d1 := g.Degree(t1.neighbors[i-1])
		d2 := g.Degree(t1.neighbors[i])
		assert.True(t, d1 >= d2, fmt.Sprintf("deg(%v) [%v] must be greater than deg(%v) [%v]", t1.neighbors[i-1], d1, t1.neighbors[i], d2))
	}
}

func TestTagCompare(t *testing.T) {
	g := MkGraph(7)
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
}
