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
