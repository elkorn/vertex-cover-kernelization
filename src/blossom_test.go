package graph

import (
	"container/list"
	"testing"

	"github.com/deckarep/golang-set"
)

func TestContractBlossom(t *testing.T) {
	g := MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	cycle := list.New()
	vertices := mapset.NewSet()

	g.ForAllVertices(func(v Vertex, index int, done chan<- bool) {
		cycle.PushBack(v)
		vertices.Add(v)
	})

	blossom := MkBlossom(1, cycle, vertices)
	inVerboseContext(func() {
		contractGraph(g, blossom)
	})
}
