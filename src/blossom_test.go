package graph

import "testing"

func TestContractBlossom(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	gv := MkGraphVisualizer()
	gv.Display(g)
	contractionMap := make(NeighborMap, 4)
	contractionMap[0] = g.getNeighbors(1)
	g.contractEdges(contractionMap)
	contractionMap[0] = g.getNeighbors(1)
	g.contractEdges(contractionMap)
	gv.Display(g)
}
