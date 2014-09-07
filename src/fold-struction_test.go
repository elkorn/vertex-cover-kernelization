package graph

import "testing"

func TestStruction(t *testing.T) {
	g := MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5) // 5 is t on the diagram
	g.AddEdge(2, 6) // 6 is w on the diagram
	g.AddEdge(3, 7) // 7 is x on the diagram
	g.AddEdge(4, 8) // 8 is y on the diagram
	g.AddEdge(4, 9) // 9 is z on the diagram

	showGraph(g)

	g1 := struction(g, Vertex(1))

	showGraph(g1)
}
