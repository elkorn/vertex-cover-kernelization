package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Binary search impl. does not seem to find the first element. Why?
// Uncomment the following tests to see them fail.

// func TestSearchInt(t *testing.T) {
// 	data := []int{1, 2, 3}
//
// 	searchFor := func(v int) int {
// 		return sort.Search(len(data), func(i int) bool {
// 			Debug("%v", i)
// 			return data[i] == v
// 		})
// 	}
//
// 	assert.Equal(t, 0, searchFor(1))
// 	assert.Equal(t, 1, searchFor(2))
// }
//
// func TestSearchNeighbor(t *testing.T) {
// 	data := Neighbors{1, 2, 3}
//
// 	searchFor := func(v Vertex) int {
// 		return sort.Search(len(data), func(i int) bool {
// 			return data[i] == v
// 		})
// 	}
//
// 	assert.Equal(t, 0, searchFor(1))
// 	assert.Equal(t, 1, searchFor(2))
// 	assert.Equal(t, 2, searchFor(3))
// }
//
func TestContains(t *testing.T) {
	neighbors := Neighbors{2, 3}

	assert.True(t, contains(neighbors, 2))
}
