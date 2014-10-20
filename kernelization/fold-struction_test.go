package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/stretchr/testify/assert"
)

func TestStruction(t *testing.T) {
	g := graph.MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5) // 5 is t on the diagram
	g.AddEdge(2, 6) // 6 is w on the diagram
	g.AddEdge(3, 7) // 7 is x on the diagram
	g.AddEdge(4, 8) // 8 is y on the diagram
	g.AddEdge(4, 9) // 9 is z on the diagram

	g1 := struction(g, graph.Vertex(1))
	assert.Equal(t, 7, g1.NVertices())
	assert.Equal(t, 8, g1.NEdges())
	assert.True(t, g1.HasVertex(graph.Vertex(10)))
	assert.True(t, g1.HasVertex(graph.Vertex(11)))
	assert.False(t, g1.HasVertex(graph.Vertex(1)))
	assert.False(t, g1.HasVertex(graph.Vertex(2)))
	assert.False(t, g1.HasVertex(graph.Vertex(3)))
	assert.False(t, g1.HasVertex(graph.Vertex(4)))
	assert.True(t, g1.HasEdge(11, 10))
	assert.True(t, g1.HasEdge(11, 7))
	assert.True(t, g1.HasEdge(11, 8))
	assert.True(t, g1.HasEdge(11, 9))
	assert.True(t, g1.HasEdge(10, 5))
	assert.True(t, g1.HasEdge(10, 6))
	assert.True(t, g1.HasEdge(10, 8))
	assert.True(t, g1.HasEdge(10, 9))
}

func TestStruction2(t *testing.T) {
	g := graph.MkGraph(12)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(4, 5)

	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(3, 8)
	g.AddEdge(4, 9)
	g.AddEdge(5, 10)
	g.AddEdge(5, 11)
	g.AddEdge(5, 12)

	g1 := struction(g, graph.Vertex(1))
	assert.Equal(t, 10, g1.NVertices())
	assert.Equal(t, 15, g1.NEdges())
	assert.Equal(t, 6, g1.Degree(13))
	assert.Equal(t, 5, g1.Degree(14))
	assert.Equal(t, 7, g1.Degree(15))
	assert.True(t, g1.HasEdge(13, 14))
	assert.True(t, g1.HasEdge(13, 15))
	assert.True(t, g1.HasEdge(13, 6))
	assert.True(t, g1.HasEdge(13, 10))
	assert.True(t, g1.HasEdge(13, 11))
	assert.True(t, g1.HasEdge(13, 12))
	assert.True(t, g1.HasEdge(13, 12))

	assert.True(t, g1.HasEdge(14, 15))
	assert.True(t, g1.HasEdge(14, 7))
	assert.True(t, g1.HasEdge(14, 8))
	assert.True(t, g1.HasEdge(14, 9))

	assert.True(t, g1.HasEdge(15, 7))
	assert.True(t, g1.HasEdge(15, 8))
	assert.True(t, g1.HasEdge(15, 10))
	assert.True(t, g1.HasEdge(15, 11))
	assert.True(t, g1.HasEdge(15, 12))
}

func TestReduceAlmostCrown(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	findCrown(g, nil, utility.MAX_INT)
	utility.Debug("\n")
	reduceAlmostCrown(g, nil, utility.MAX_INT)
}

func TestGeneralFold2(t *testing.T) {
	// TODO: There probably is a bug when connecting the almost-crown vertices to
	// the fold-root. Investigate.
	g := graph.MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 8)
	g.AddEdge(8, 9)
	g.AddEdge(8, 2)
	g.AddEdge(8, 3)
	g.AddEdge(9, 2)
	g.AddEdge(9, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	generalFold(g, nil, 190)
	// g1 :=  generalFold(g, 0)
}

func TestGeneralFold3(t *testing.T) {
	g := graph.MkGraph(7)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	// g1 :=  generalFold(g)
}

func TestGeneralFold4(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	gPrime, kPrime := generalFold(g, nil, 1)
	assert.Equal(t, 1, gPrime.NVertices())
	assert.Equal(t, -1, kPrime)
	assert.True(t, g.HasVertex(graph.Vertex(4)))
}

func TestGeneralFold5(t *testing.T) {
	g := graph.MkGraph(9)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(5, 6)
	g.AddEdge(6, 8)
	g.AddEdge(7, 8)
	g.AddEdge(7, 9)
	g.AddEdge(8, 9)
	g.AddEdge(9, 2)
	generalFold(g, nil, utility.MAX_INT)
	// TODO: There are bugs in generalFold or findCrown.
	// 1) nothing gets folded in this graph. According to Lemma 5.2, there
	// should be a parameter reduction of at least 2.
	// 2) Doing a consecutive fold on the graph causes a crash.
	// generalFold(gPrime, nil, utility.MAX_INT)
}

// func TestBussKernelization(t *testing.T) {
// 	g := MkPetersenGraph()
// 	InVerboseContext(func() {
// 		reduction := kernelizeIfHasCoverOfSize(g, 3)
// 		utility.Debug("Reduction by Buss: %v", reduction)
// 	})
// }
