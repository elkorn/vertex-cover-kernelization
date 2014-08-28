package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistanceInTree(t *testing.T) {
	root := MkVertex(0)
	tr := MkTree(root, 5)
	for i := 1; i < 5; i++ {
		tr.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	assert.Equal(t, 4, tr.Distance(MkVertex(4), root))
}

func TestPathInTree(t *testing.T) {
	root := MkVertex(0)
	tr := MkTree(root, 5)
	expectedEdges := make([]*Edge, 0, 4)
	for i := 4; i > 0; i-- {
		expectedEdges = append(expectedEdges, MkEdgeFromInts(i-1, i))
		tr.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	actual := tr.Path(MkVertex(4), root)

	for i, expected := range expectedEdges {
		assert.Equal(t, *expected, *actual[i])
	}
}

func TestPathEndpointsOrdering(t *testing.T) {
	root := MkVertex(0)
	tr := MkTree(root, 5)
	for i := 1; i < 5; i++ {
		tr.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	expected := tr.Path(MkVertex(4), root)
	n := len(expected)
	actual := tr.Path(root, MkVertex(4))

	for i, expectedEdge := range expected {
		assert.Equal(t, *expectedEdge, *actual[n-i-1])
	}
}
