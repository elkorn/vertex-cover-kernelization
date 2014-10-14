package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
