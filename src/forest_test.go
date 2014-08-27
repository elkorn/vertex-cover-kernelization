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

func TestPathInForest(t *testing.T) {
	// TODO use only forest API, make this test pass. @start-from-here
	root1 := MkVertex(0)
	root2 := MkVertex(5)
	tr1 := MkTree(root1, 10)
	tr2 := MkTree(root2, 10)
	for i := 1; i < 5; i++ {
		tr1.AddEdge(MkVertex(i-1), MkVertex(i))
		tr2.AddEdge(MkVertex(i+4), MkVertex(i+5))
	}

	f := MkForest(10)
	f.AddTree(tr1)
	f.AddTree(tr2)
	ep1 := MkTreePath(root1, MkVertex(4))
	ep2 := MkTreePath(root2, MkVertex(9))

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inVerboseContext(func() {
		actual := f.path(ep1, ep2)
		assert.Equal(t, 8, len(actual))
		assert.Equal(t, expected, actual)
	})
}
