package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutation(t *testing.T) {
	input := []int{1, 2, 3}
	result := permutations(input)
	assert.Equal(t, 6, len(result))
	assert.Equal(t, rng(1, 2, 3), result[0])
	assert.Equal(t, rng(1, 3, 2), result[1])
	assert.Equal(t, rng(2, 1, 3), result[2])
	assert.Equal(t, rng(2, 3, 1), result[3])
	assert.Equal(t, rng(3, 1, 2), result[4])
	assert.Equal(t, rng(3, 2, 1), result[5])
}
