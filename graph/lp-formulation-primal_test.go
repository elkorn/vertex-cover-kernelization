package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimalFormulation1(t *testing.T) {
	g := mkGraph1()
	formulation := mklpPrimalFormulation(g, 10)
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
	InVerboseContext(func() {
		formulation := mklpPrimalFormulation(g, 10)
		p, q, r, err := formulation.solve()
		assert.Nil(t, err)
		indep, _ := isIndependentSet(p, g)
		assert.True(t, indep)
		Debug("P: %v", p)
		Debug("Q: %v", q)
		Debug("R: %v", r)
	})
}

// func TestPrimalFormulation3(t *testing.T) {
// 	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
// 	formulation := mklpPrimalFormulation(g, 10)
// 	p, q, r, err := formulation.solve()
// 	assert.Nil(t, err)
// 	assert.NotEqual(t, 0, q.Cardinality(), "Why is q always empty?")
// 	assert.NotEqual(t, g.NVertices(), p.Union(r).Cardinality())
// }
