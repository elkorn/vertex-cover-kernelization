package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntQueuePush(t *testing.T) {
	q := MkIntQueue(5)

	q.Push(10)
	q.Push(9)
	q.Push(8)
	q.Push(7)
	q.Push(6)
	assert.Equal(t, 10, q.Pop())
	assert.Equal(t, 9, q.Pop())
	assert.Equal(t, 8, q.Pop())
	assert.Equal(t, 7, q.Pop())
	assert.Equal(t, 6, q.Pop())
	assert.True(t, q.Empty())
}
