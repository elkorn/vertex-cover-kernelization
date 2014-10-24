package vc

import (
	"fmt"
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestNaiveVC2(t *testing.T) {
	g := graph.MkPetersenGraph()
	innerVertices, outerVertices := mapset.NewThreadUnsafeSet(), mapset.NewThreadUnsafeSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(graph.Vertex(i))
		innerVertices.Add(graph.Vertex(i + 5))
	}

	found, cover := NaiveVC(g, 6)
	assert.True(t, found)
	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))

}

func TestNaiveVC3(t *testing.T) {
	g := graph.MkReversePetersenGraph()
	innerVertices, outerVertices := mapset.NewThreadUnsafeSet(), mapset.NewThreadUnsafeSet()
	for i := 1; i < 6; i++ {
		outerVertices.Add(graph.Vertex(i))
		innerVertices.Add(graph.Vertex(i + 5))
	}

	found, cover := NaiveVC(g, 6)

	assert.True(t, found)
	assert.Equal(t, 6, cover.Cardinality())
	assert.Equal(t, 3, cover.Intersect(outerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  outer vertices (from %v)", cover, outerVertices))
	assert.Equal(t, 3, cover.Intersect(innerVertices).Cardinality(), fmt.Sprintf("The cover of the Petersen graph (%v) should contain 3  inner vertices (from %v)", cover, innerVertices))
}
