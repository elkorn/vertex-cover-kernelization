package kernelization

import (
	"fmt"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestGetVertices(t *testing.T) {
	g := graph.MkGraph(3)

	expected := make(graph.Vertices, 6)
	for i := 0; i < 6; i++ {
		expected[i] = graph.Vertex(i + 1)
	}

	actual := getVertices(g)
	assert.Equal(t, expected, actual)
}

func assertAllEdgesEqual(t *testing.T, expected graph.Edges, actual *graph.Graph) {
	assert.Equal(t, len(expected), actual.NEdges(), "The number of edges must be the same.")
	check := func(actual graph.Edge) {
		result := false
		for _, expEdge := range expected {
			if *expEdge == actual {
				result = true
				break
			}
		}

		assert.True(t, result, "Expected the graph to have edge "+fmt.Sprintf("%v", actual))
	}
	actual.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		check(*edge)
	})
}

func TestMakeBipartite(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)

	expectedVertices := graph.Vertices{1, 2, 3, 4, 5, 6, 7, 8}
	expectedEdges := graph.Edges{
		graph.MkEdge(4, 5),
		graph.MkEdge(2, 7),
	}

	actual := makeBipartite(g)
	for _, v := range expectedVertices {
		assert.True(t, actual.HasVertex(v))
	}

	assertAllEdgesEqual(t, expectedEdges, actual)
}
