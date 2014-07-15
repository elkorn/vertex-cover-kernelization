package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test_eq struct {
	a int
	b int
}

func TestStructEquality(t *testing.T) {
	st1 := test_eq{1, 1}
	st2 := test_eq{1, 1}

	assert.Equal(t, st1, st2)
}
