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
