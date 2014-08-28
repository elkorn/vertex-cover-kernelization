package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootInForest(t *testing.T) {
	f := MkForest(5)
	root := MkVertex(0)
	tr := MkTree(root, 5)
	v := MkVertex(1)
	tr.AddEdge(root, v)
	f.AddTree(tr)

	assert.Equal(t, root, f.Root(v))
}

func TestDistanceInForest(t *testing.T) {
	root1 := MkVertex(0)
	root2 := MkVertex(5)
	tr1 := MkTree(root1, 6)
	tr2 := MkTree(root2, 6)
	for i := 1; i < 5; i++ {
		tr1.AddEdge(MkVertex(i-1), MkVertex(i))
	}

	f := MkForest(6)
	f.AddTree(tr2)
	f.AddTree(tr1)

	assert.Equal(t, 4, f.Distance(MkVertex(4), f.Root(MkVertex(4))))
}

func TestPathInForest1(t *testing.T) {
	root1 := MkVertex(0)
	root2 := MkVertex(5)
	f := MkForest(10)
	f.AddTree(MkTree(root1, 10))
	f.AddTree(MkTree(root2, 10))

	for i := 1; i < 5; i++ {
		f.AddEdge(root1, MkEdgeFromInts(i-1, i))
		f.AddEdge(root2, MkEdgeFromInts(i+4, i+5))
	}

	ep1 := MkTreePath(root1, MkVertex(4))
	ep2 := MkTreePath(root2, MkVertex(9))

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := f.Path(ep1, ep2)

	assert.Equal(t, 10, len(actual))
	assert.Equal(t, expected, actual)

	expected = []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
	actual = f.Path(ep2, ep1)
	assert.Equal(t, 10, len(actual))
	assert.Equal(t, expected, actual)

	// TODO understand this case better, look at the diagrams.
	// Come up with a correct `expected` value and make this pass. @start-from-here
	ep3 := MkTreePath(MkVertex(9), root2)
	expected = []int{0, 1, 2, 3, 4, 9, 8, 7, 6, 5}
	actual = f.Path(ep1, ep3)
	assert.Equal(t, 10, len(actual))
	assert.Equal(t, expected, actual)
}
