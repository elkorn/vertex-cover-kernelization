package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntStack(t *testing.T) {
	s := MkIntStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, []int{3, 2, 1}, s.Values())

	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())

	defer func() {
		// Should recover always run deferred?
		assert.NotNil(t, recover(), "Should throw an error while trying to pop from an empty stack.")
	}()
	s.Pop()
}

func TestIntStackIter(t *testing.T) {
	s := MkIntStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	expected := []int{3, 2, 1}
	i := 0
	for actual := range s.Iter() {
		assert.Equal(t, expected[i], actual)
		i++
	}
}
