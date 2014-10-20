package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestNtReductionAbuKhzam1(t *testing.T) {
	g := graph.MkGraph1()
	formulation := mkNtReductionAbuKhzam(g, 10)
	_, _, _, err := formulation.solve()
	assert.Nil(t, err)
}

func TestNtReductionAbuKhzam2(t *testing.T) {
	g := graph.MkPetersenGraph()
	formulation := mkNtReductionAbuKhzam(g, 10)
	_, _, _, err := formulation.solve()
	assert.Nil(t, err)
}

func TestNtReductionAbuKhzam3(t *testing.T) {
	g := graph.MkGraph(7)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(5, 6)
	formulation := mkNtReductionAbuKhzam(g, 10)
	p, q, r, err := formulation.solve()
	assert.True(t, p.Contains(graph.Vertex(2)))
	assert.True(t, p.Contains(graph.Vertex(5)))
	assert.Equal(t, 0, q.Cardinality(), "{2,5} is the vertex cover, no kernel should remain")
	assert.Equal(t, g.NVertices()-p.Cardinality(), r.Cardinality(), "{2,5} is the vertex cover, all other vertices can be excluded.")
	assert.Nil(t, err)
}

// func TestNtReductionAbuKhzam3(t *testing.T) {
// 	g := ScanGraph("../examples/sh2/sh2-3.dim")
// 	formulation := mkNtReductionAbuKhzam(g, 10)
// 	p, q, r, err := formulation.solve()
// 	assert.Nil(t, err)
// 	assert.NotEqual(t, 0, q.Cardinality(), "Why is q always empty?")
// 	assert.NotEqual(t, g.NVertices(), p.Union(r).Cardinality())
// }
