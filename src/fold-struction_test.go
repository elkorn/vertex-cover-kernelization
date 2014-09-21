package graph

import (
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

func TestGeneralFold1(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)

	// showGraph(g)
	// inVerboseContext(func() {
	// generalFold(g, nil, MAX_INT)
	// })
	// gPrime := generalFold(g, nil, MAX_INT)
	// showGraph(gPrime)

	// inVerboseContext(func() {
	// 	crown := findCrown(gPrime, nil, MAX_INT)
	// 	Debug("%v", crown)
	// })
	// gv.highlightCrown(crown)
	// gv.Display()

	// assert.Equal(t, 5, gPrime.NVertices())
	// assert.True(t, gPrime.hasVertex(9))
	// assert.True(t, gPrime.hasEdge(9, 5))
	// assert.True(t, gPrime.hasEdge(9, 6))
	// assert.True(t, gPrime.hasEdge(9, 7))
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

	// showGraph(g)
	generalFold(g, nil, 190)
	// showGraph(g)
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
