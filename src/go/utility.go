package graph

import "sort"

func removeAt(source Edges, position int) Edges {
	return append(source[:position], source[position+1:]...)
}

func contains(neighbors Neighbors, v Vertex) bool {
	length := len(neighbors)
	foundIndex := sort.Search(length, func(i int) bool { return neighbors[i] == v })

	return foundIndex < length && neighbors[foundIndex] == v
}
