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
	for i := 1; i < 5; i++ {
		expectedEdges = append(expectedEdges, MkEdgeFromInts(i-1, i))
		tr.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	path := tr.Path(MkVertex(4), root)
	actual := path.Values()

	for i, expected := range expectedEdges {
		expectedEdge := expected
		actualEdge := actual[i].(*Edge)
		assert.Equal(t, *expectedEdge, *actualEdge)
	}
}

func TestPathEndpointsOrdering(t *testing.T) {
	root := MkVertex(0)
	tr := MkTree(root, 5)
	for i := 1; i < 5; i++ {
		tr.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	expected := tr.Path(MkVertex(4), root).Values()
	n := len(expected)
	actual := tr.Path(root, MkVertex(4)).Values()

	for i, e := range expected {
		expectedEdge := e.(*Edge)
		actualEdge := actual[n-i-1].(*Edge)
		assert.Equal(t, *expectedEdge, *actualEdge)
	}
}
