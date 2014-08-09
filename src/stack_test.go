package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := MkStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())

	defer func() {
		// Should recover always run deferred?
		assert.NotNil(t, recover(), "Should throw an error while trying to pop from an empty stack.")
	}()
	s.Pop()
}
