package graph

func removeAt(source Edges, position int) Edges {
	return append(source[:position], source[position+1:]...)
}

func indexOf(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}

	return n + 1
}

func contains(neighbors Neighbors, v Vertex) bool {
	// Debug(fmt.Sprintf("==== SEARCH ====="))
	length := len(neighbors)
	// Debug(fmt.Sprintf("Searching for %v in %v", v, neighbors))
	foundIndex := indexOf(length, func(i int) bool {
		return neighbors[i] == v
	})

	// Debug(fmt.Sprintf("Found index %v", foundIndex))
	// Debug(fmt.Sprintf("==== END SEARCH ===="))
	return foundIndex < length && neighbors[foundIndex] == v
}
