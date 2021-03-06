package matching

import (
	"container/list"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestMaximalMatching(t *testing.T) {
	g := graph.MkGraph1()

	m, o := FindMaximalMatching(g)
	assert.Equal(t, 2, m.NEdges(), "The matching of graph1 should contain 2 edges.")
	assert.Equal(t, 1, o.Cardinality(), "The outsiders of graph1 should contain 1 unmatched vertex.")
	assert.True(t, m.HasEdge(1, 2))
	assert.True(t, m.HasEdge(3, 4))

	g = graph.MkGraph5()
	m, o = FindMaximalMatching(g)
	assert.Equal(t, 3, m.NEdges(), "The matching of graph5 should contain 3 edges.")
	assert.Equal(t, 1, o.Cardinality(), "The outsiders of graph5 should contain 1 unmatched vertex.")
	assert.True(t, m.HasEdge(1, 2))
	assert.True(t, m.HasEdge(4, 5))
	assert.True(t, m.HasEdge(6, 7))
}

func TestMaximumMatching1(t *testing.T) {
	g := graph.MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(1, 5)

	max := FindMaximumMatching(g)

	assert.Equal(t, 2, max.NEdges())
	assert.True(t, max.HasEdge(1, 2))
	assert.True(t, max.HasEdge(3, 4))
}

func TestMaximumMatching2(t *testing.T) {
	g := graph.MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 6)
	g.AddEdge(4, 5)
	g.AddEdge(1, 5)

	max := FindMaximumMatching(g)

	assert.Equal(t, 3, max.NEdges())
	assert.True(t, max.HasEdge(1, 2))
	assert.True(t, max.HasEdge(3, 6))
	assert.True(t, max.HasEdge(4, 5))
}

func TestMaximumMatchingSmallBlossom(t *testing.T) {
	g := graph.MkGraph(5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)

	max := FindMaximumMatching(g)

	assert.Equal(t, 2, max.NEdges())
	assert.True(t, max.HasEdge(1, 2))
	assert.True(t, max.HasEdge(3, 4))
}

func TestMaximumMatching3(t *testing.T) {
	g := graph.MkGraph(6)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)

	max := FindMaximumMatching(g)
	assert.Equal(t, 2, max.NEdges())
	assert.True(t, max.HasEdge(1, 3))
	assert.True(t, max.HasEdge(2, 4))
}

func TestIndexOf(t *testing.T) {
	list := list.New()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	assert.Equal(t, 0, indexOf(1, list))
	assert.Equal(t, 1, indexOf(2, list))
	assert.Equal(t, 2, indexOf(3, list))
}

func TestMaximumMatching4(t *testing.T) {
	g := graph.MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(6, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 7)
	g.AddEdge(8, 7)
	g1 := FindMaximumMatching(g)
	assert.Equal(t, 3, g1.NEdges())
}
