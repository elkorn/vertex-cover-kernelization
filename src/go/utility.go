package graph

import (
	"fmt"
	"sort"
)

func removeAt(source Edges, position int) Edges {
	return append(source[:position], source[position+1:]...)
}

func contains(neighbors Neighbors, v Vertex) bool {
	length := len(neighbors)
	Debug(fmt.Sprintf("Searching for %v in %v", v, neighbors))
	foundIndex := sort.Search(length, func(i int) bool {
		Debug(fmt.Sprintf("[%v] %v == %v ? %v", i, neighbors[i], v, neighbors[i] == v))
		return neighbors[i] == v
	})

	Debug(fmt.Sprintf("Found index %v", foundIndex))
	return foundIndex < length && neighbors[foundIndex] == v
}
