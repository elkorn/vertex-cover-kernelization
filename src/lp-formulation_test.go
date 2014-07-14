package graph

import (
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
