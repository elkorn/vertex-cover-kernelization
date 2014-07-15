package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertexEquality(t *testing.T) {
	a := Vertex{1, 0}
	b := Vertex{1, 0}

	assert.Equal(t, a, b)
}
