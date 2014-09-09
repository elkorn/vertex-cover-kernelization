package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVertices(t *testing.T) {
	g := MkGraph(3)

	expected := make(Vertices, 6)
	for i := 0; i < 6; i++ {
		expected[i] = Vertex(i + 1)
	}

	actual := getVertices(g)
	assert.Equal(t, expected, actual)
}

func assertAllEdgesEqual(t *testing.T, expected Edges, actual *Graph) {
	assert.Equal(t, len(expected), actual.NEdges(), "The number of edges must be the same.")
	check := func(actual Edge) {
		result := false
		for _, expEdge := range expected {
			if *expEdge == actual {
				result = true
				break
			}
		}

		assert.True(t, result, "Expected the graph to have edge "+fmt.Sprintf("%v", actual))
	}
	actual.ForAllEdges(func(edge *Edge, done chan<- bool) {
		check(*edge)
	})
}

func TestMakeBipartite(t *testing.T) {
	g := MkGraph(4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)

	expectedVertices := Vertices{1, 2, 3, 4, 5, 6, 7, 8}
	expectedEdges := Edges{
		MkEdge(4, 5),
		MkEdge(2, 7),
	}

	actual := makeBipartite(g)
	for _, v := range expectedVertices {
		assert.True(t, actual.hasVertex(v))
	}

	assertAllEdgesEqual(t, expectedEdges, actual)
}

func TestPrintBipartite(t *testing.T) {
	g := mkGraph1()
	// g1 := makeBipartite(g)
	// nf := mkNetworkFlow(g)
	gv1 := MkGraphVisualizer(g)
	// gv2 := MkGraphVisualizer(g1)

	gv1.Display()
	networkFlowKernelization(g, 3)
	// max.ForAllEdges(func(edge *Edge, done chan<- bool) {
	// 	gv2.HighlightEdge(edge, "red")
	// })
	// gv1.HighlightCover(cover, "yellow")
	gv1.Display()

}
