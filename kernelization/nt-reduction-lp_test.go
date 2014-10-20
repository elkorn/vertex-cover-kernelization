package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/stretchr/testify/assert"
)

func TestNtReductionLP(t *testing.T) {
	g := graph.MkGraph(6)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(5, 6)
	formulation := mkNtReductionLP(g, 10)
	p, q, r, err := formulation.solve()
	assert.True(t, p.Contains(Vertex(2)))
	assert.True(t, p.Contains(Vertex(5)))
	assert.Equal(t, 0, q.Cardinality(), "{2,5} is the vertex cover, no kernel should remain")
	assert.Equal(t, g.NVertices()-p.Cardinality(), r.Cardinality(), "{2,5} is the vertex cover, all other vertices can be excluded.")
	assert.Nil(t, err)
	crown := findCrown(g, nil, 10)
	// InVerboseContext(func() {
	utility.Debug("Crown: %v", crown)
	// })

}
