package graph

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectiveFunction(t *testing.T) {
	s1 := make(map[Vertex]int)
	for i := 0; i < 10; i++ {
		s1[Vertex(i)] = 0
	}

	s1[Vertex(5)] = 1
	s1[Vertex(6)] = 1

	s2 := make(map[Vertex]int)
	for i := 0; i < 10; i++ {
		s2[Vertex(i)] = 0
	}

	s2[Vertex(5)] = 1

	assert.Equal(t, s2, objectiveFunction([]map[Vertex]int{s1, s2}))
}

func TestResolveConflict(t *testing.T) {
	n1 := Node{1, 1}
	n2 := Node{2, 3}

	assert.Equal(t, n2, resolveConflict(n1, n2))

	n1.degree = 4
	assert.Equal(t, n1, resolveConflict(n1, n2))

	// Seeding the rand differently can break this test.
	rand.Seed(1)
	n2.degree = n1.degree
	assert.Equal(t, n2, resolveConflict(n1, n2))
}

func TestCalculateLowerBound(t *testing.T) {
	g := mkGraphWithVertices(10)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(1, 6)
	g.AddEdge(8, 9)

	selection := make(map[Vertex]int)
	for i := range g.Vertices {
		selection[i] = 0
	}

	// It's very sensible to reimplement as a struct with a degree field.
	// It's gonna be looked up a lot during this computation.
	assert.Equal(t, 3, computeLowerBound(g, selection))

	selection[8] = 1
	selection[9] = 1
	selection[5] = 1

	assert.Equal(t, 5, computeLowerBound(g, selection))
}
