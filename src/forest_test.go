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
	edges := make([]*Edge, 8)

	for i := 1; i < 5; i++ {
		edges[i-1] = MkEdgeFromInts(i-1, i)
		edges[i+3] = MkEdgeFromInts(i+4, i+5)
		f.AddEdge(root1, edges[i-1])
		f.AddEdge(root2, edges[i+3])
	}

	ep1 := MkTreePath(root1, MkVertex(4))
	ep2 := MkTreePath(root2, MkVertex(9))

	var expected []*Edge
	var actual []*Edge

	expected = []*Edge{edges[0], edges[1], edges[2], edges[3], edges[4], edges[5], edges[6], edges[7]}
	actual = f.Path(ep1, ep2)

	assert.Equal(t, 8, len(actual))
	assert.Equal(t, expected, actual)

	expected = []*Edge{edges[4], edges[5], edges[6], edges[7], edges[0], edges[1], edges[2], edges[3]}
	actual = f.Path(ep2, ep1)
	assert.Equal(t, 8, len(actual))
	assert.Equal(t, expected, actual)

	ep3 := MkTreePath(MkVertex(9), root2)
	expected = []*Edge{edges[7], edges[6], edges[5], edges[4]}
	actual = f.Path(ep3)

	assert.Equal(t, 4, len(actual))
	assert.Equal(t, expected, actual)
	// TODO: what about not-in-tree edges?
	expected = []*Edge{edges[0], edges[1], edges[2], edges[3], edges[7], edges[6], edges[5], edges[4]}
	actual = f.Path(ep1, ep3)
	assert.Equal(t, 8, len(actual))
	assert.Equal(t, expected, actual)
}
