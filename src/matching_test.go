package graph

import (
	"fmt"
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestMaximalMatching(t *testing.T) {
	g := mkGraph1()

	m, o := maximalMatching(g)
	assert.Equal(t, 2, m.Cardinality(), "The matching of graph1 should contain 2 edges.")
	assert.Equal(t, 6, o.Cardinality(), "The outsiders of graph1 should contain 6 unmatched edges.")
	assert.True(t, m.Contains(g.getEdgeByCoordinates(0, 1)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(2, 3)))

	g = mkGraph5()
	m, o = maximalMatching(g)
	assert.Equal(t, 3, m.Cardinality(), "The matching of graph5 should contain 3 edges.")
	assert.Equal(t, 4, o.Cardinality(), "The outsiders of graph5 should contain 4 unmatched edges.")
	assert.True(t, m.Contains(g.getEdgeByCoordinates(0, 1)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(3, 4)))
	assert.True(t, m.Contains(g.getEdgeByCoordinates(5, 6)))
}

func TestIsExposed(t *testing.T) {
	matching := mapset.NewSet()
	matching.Add(Edge{1, 2, false})
	matching.Add(Edge{2, 3, false})
	matching.Add(Edge{4, 5, false})
	matching.Add(Edge{7, 8, false})

	assert.False(t, Vertex(1).isExposed(matching))
	assert.False(t, Vertex(2).isExposed(matching))
	assert.False(t, Vertex(3).isExposed(matching))
	assert.False(t, Vertex(4).isExposed(matching))
	assert.False(t, Vertex(5).isExposed(matching))
	assert.False(t, Vertex(7).isExposed(matching))
	assert.False(t, Vertex(8).isExposed(matching))

	assert.True(t, Vertex(6).isExposed(matching))
	assert.True(t, Vertex(9).isExposed(matching))
}

func TestIsAugmentingPath(t *testing.T) {
	path := []int{0, 1, 2, 3, 4, 5, 6, 7}
	matching := mapset.NewSet()
	matching.Add(MkEdgeValFromInts(1, 2))
	matching.Add(MkEdgeValFromInts(3, 4))
	matching.Add(MkEdgeValFromInts(5, 6))

	assert.True(t, isAlternatingPathWithMatching(path, matching), "The test path should be alternating.")
	assert.True(t, MkVertex(0).isExposed(matching), "The start point should be exposed.")
	assert.True(t, MkVertex(7).isExposed(matching), "The end point should be exposed.")
	assert.True(t, isAugmentingPath(path, matching), "The test path should be augmenting.")
}

func TestMatchingAugmentation(t *testing.T) {
	path := []int{0, 1, 2, 3, 4, 5}
	matching := mapset.NewSet()
	matching.Add(MkEdgeValFromInts(1, 2))
	matching.Add(MkEdgeValFromInts(3, 4))
	expected := []Edge{
		MkEdgeValFromInts(0, 1),
		MkEdgeValFromInts(2, 3),
		MkEdgeValFromInts(4, 5),
	}

	inVerboseContext(func() {
		matchingAugmentation(path, matching)
	})
	augmentation := matchingAugmentation(path, matching)
	assert.Equal(t, 3, augmentation.Cardinality(), "Given augmentation should contain 3 edges.")
	for _, edge := range expected {
		assert.True(t, augmentation.Contains(edge), fmt.Sprintf("Given augmentation should contain edge %v-%v", edge.from, edge.to))
	}
}
