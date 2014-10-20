package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDualformulation1(t *testing.T) {
	g := MkGraph1()
	formulation := mklpDualFormulation(g, 10)
	matching, err := formulation.solve()
	assert.Nil(t, err)
	assert.Equal(t, 2, matching.Cardinality())
	gv := MkGraphVisualizer(g)
	gv.HighlightMatchingSet(matching, "red")
	// gv.Display()
}

func TestDualformulation2(t *testing.T) {
	g := mkPetersenGraph()
	formulation := mklpDualFormulation(g, 10)
	_, err := formulation.solve()
	assert.Nil(t, err)
}
