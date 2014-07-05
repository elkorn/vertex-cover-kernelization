package graph

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInt(t *testing.T) {
	data := []int{1, 2, 3}

	searchFor := func(v int) int {
		return sort.Search(len(data), func(i int) bool {
			return data[i] == v
		})
	}

	assert.Equal(t, 1, searchFor(2))
}

// func TestContains(t *testing.T) {
// 	neighbors := Neighbors{2, 3}
//
// 	assert.True(t, contains(neighbors, 2))
// }
