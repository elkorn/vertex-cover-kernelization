package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCrown1(t *testing.T) {
	g := MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(8, 7)

	crown := findCrown(g)
	assert.Equal(t, 1, crown.Width())
	assert.True(t, crown.H.Contains(Vertex(4)))
	assert.True(t, crown.I.Contains(Vertex(1)))
	assert.True(t, crown.I.Contains(Vertex(2)))
	assert.True(t, crown.I.Contains(Vertex(3)))
	// gv := MkGraphVisualizer(g)
	// for vInter := range crown.I.Iter() {
	// 	gv.HighlightVertex(vInter.(Vertex), "lightgray")
	// }

	// for vInter := range crown.H.Iter() {
	// 	gv.HighlightVertex(vInter.(Vertex), "yellow")
	// }

	// gv.Display()
}
