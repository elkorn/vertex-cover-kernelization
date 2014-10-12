package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimalFormulation1(t *testing.T) {
	g := mkGraph1()
	formulation := mkNtReductionAbuKhzam(g, 10)
	p, q, r, err := formulation.solve()
	InVerboseContext(func() {
		Debug("P: %v", p)
		Debug("Q: %v", q)
		Debug("R: %v", r)
	})
	assert.Nil(t, err)
}

func TestPrimalFormulation2(t *testing.T) {
	g := mkPetersenGraph()
	formulation := mkNtReductionAbuKhzam(g, 10)
	p, q, r, err := formulation.solve()
	assert.Nil(t, err)
	indep, _ := isIndependentSet(p, g)
	assert.True(t, indep)
	InVerboseContext(func() {
		Debug("P: %v", p)
		Debug("Q: %v", q)
		Debug("R: %v", r)
	})
}

func TestPrimalFormulation3(t *testing.T) {
	g := MkGraph(7)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(5, 6)
	formulation := mkNtReductionAbuKhzam(g, 10)
	p, q, r, err := formulation.solve()
	assert.True(t, p.Contains(Vertex(2)))
	assert.True(t, p.Contains(Vertex(5)))
	assert.Equal(t, 0, q.Cardinality(), "{2,5} is the vertex cover, no kernel should remain")
	assert.Equal(t, g.NVertices()-p.Cardinality(), r.Cardinality(), "{2,5} is the vertex cover, all other vertices can be excluded.")
	assert.Nil(t, err)
}

// func TestPrimalFormulation3(t *testing.T) {
// 	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
// 	formulation := mkNtReductionAbuKhzam(g, 10)
// 	p, q, r, err := formulation.solve()
// 	assert.Nil(t, err)
// 	assert.NotEqual(t, 0, q.Cardinality(), "Why is q always empty?")
// 	assert.NotEqual(t, g.NVertices(), p.Union(r).Cardinality())
// }
