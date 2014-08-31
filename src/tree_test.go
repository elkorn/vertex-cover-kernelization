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

	var actual []*Edge

	actual = tr.Path(MkVertex(4), root)

	assert.Equal(t, len(expectedEdges), len(actual))

	for i, expected := range expectedEdges {
		assert.Equal(t, *expected, *actual[i])
	}

	tr = MkTree(root, 5)
	tr.AddEdge(root, 2)
	tr.AddEdge(2, 3)
	tr.AddEdge(2, 4)
	tr.AddEdge(4, 5)
	actual = tr.Path(3, 5)
	assert.Equal(t, 3, len(actual))
	assert.Equal(t, tr.g.getEdgeByCoordinates(1, 2), actual[0])
	assert.Equal(t, tr.g.getEdgeByCoordinates(1, 3), actual[1])
	assert.Equal(t, tr.g.getEdgeByCoordinates(3, 4), actual[2])
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

func TestCommonAncestor(t *testing.T) {
	root := MkVertex(0)
	tr := MkTree(root, 10)
	tr.AddEdge(1, 2)
	tr.AddEdge(1, 10)
	tr.AddEdge(2, 3)
	tr.AddEdge(2, 4)
	tr.AddEdge(2, 9)
	tr.AddEdge(3, 6)
	tr.AddEdge(3, 7)
	tr.AddEdge(4, 5)
	tr.AddEdge(4, 8)

	// gv := MkGraphVisualizer()
	// gv.Display(tr.g)

	assert.Equal(t, 3, tr.CommonAncestor(6, 7))
	assert.Equal(t, 3, tr.CommonAncestor(7, 6))
	assert.Equal(t, 3, tr.CommonAncestor(6, 3))
	assert.Equal(t, 2, tr.CommonAncestor(7, 5))
	assert.Equal(t, 2, tr.CommonAncestor(7, 8))
	assert.Equal(t, 2, tr.CommonAncestor(3, 4))
	assert.Equal(t, 2, tr.CommonAncestor(3, 9))
	assert.Equal(t, 1, tr.CommonAncestor(8, 10))
}
