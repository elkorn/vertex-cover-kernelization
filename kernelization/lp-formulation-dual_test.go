package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/graphviz"
	"github.com/stretchr/testify/assert"
)

func TestDualformulation1(t *testing.T) {
	g := graph.MkGraph1()
	formulation := mklpDualFormulation(g, 10)
	matching, err := formulation.solve()
	assert.Nil(t, err)
	assert.Equal(t, 2, matching.Cardinality())
	gv := graphviz.MkGraphVisualizer(g)
	gv.HighlightMatchingSet(matching, "red")
	// gv.Display()
}

func TestDualformulation2(t *testing.T) {
	g := graph.MkPetersenGraph()
	formulation := mklpDualFormulation(g, 10)
	_, err := formulation.solve()
	assert.Nil(t, err)
}
