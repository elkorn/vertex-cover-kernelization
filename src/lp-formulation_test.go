package graph

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectiveFunction(t *testing.T) {
	s1 := Selection{}
	for i := 0; i < 10; i++ {
		s1[Vertex(i)] = 0
	}

	s1[Vertex(5)] = 1
	s1[Vertex(6)] = 1

	s2 := Selection{}
	for i := 0; i < 10; i++ {
		s2[Vertex(i)] = 0
	}

	s2[Vertex(5)] = 1

	assert.Equal(t, s2, objectiveFunction([]Selection{s1, s2}))
}

func TestResolveConflict(t *testing.T) {
	g := mkGraphWithVertices(2)
	n1 := Vertex(1)
	n2 := Vertex(2)

	assert.Equal(t, n2, resolveConflict(g, n1, n2))

	g.degrees[1] = 4
	assert.Equal(t, n1, resolveConflict(g, n1, n2))

	// Seeding the rand differently can break this test.
	rand.Seed(1)
	g.degrees[n2] = g.degrees[n1]
	assert.Equal(t, n2, resolveConflict(g, n1, n2))
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

	selection := Selection{}

	assert.Equal(t, 3, computeLowerBound(g, selection))

	selection[8] = 1
	selection[9] = 1
	selection[5] = 1

	assert.Equal(t, 5, computeLowerBound(g, selection))
}

func TestGetEndpoints(t *testing.T) {
	edges := Edges{Edge{1, 2}, Edge{2, 4}}
	expected := make([]Vertex, 3)
	expected[0] = 1
	expected[1] = 2
	expected[2] = 4
	assert.Equal(t, expected, getEndpoints(edges))
}
