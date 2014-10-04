package graph

import (
	"container/list"
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestContractBlossom(t *testing.T) {
	g := MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(1, 6)

	cycle := list.New()
	vertices := mapset.NewSet()

	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		if v == 6 {
			return
		}

		cycle.PushBack(v)
		vertices.Add(v)
	})

	blossom := MkBlossom(1, cycle, vertices)
	g1 := contractGraph(g, blossom)
	assert.Equal(t, 2, g1.NVertices())
	assert.True(t, g1.HasVertex(1))
	assert.True(t, g1.HasVertex(6))
	assert.Equal(t, 1, g1.NEdges())
	assert.True(t, g1.HasEdge(1, 6))
}
