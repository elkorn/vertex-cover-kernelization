package graph

type Neighbors Vertices

func sliceIndexOf(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}

	return n + 1
}

func (self Neighbors) appendIfNotContains(v Vertex) Neighbors {
	if !Contains(self, v) {
		self = append(self, v)
	}

	return self
}

func Contains(neighbors Neighbors, v Vertex) bool {
	// Debug("==== SEARCH =====")
	length := len(neighbors)
	// Debug("Searching for %v in %v", v, neighbors)
	foundIndex := sliceIndexOf(length, func(i int) bool {
		return neighbors[i] == v
	})

	// Debug("Found index %v", foundIndex)
	// Debug("==== END SEARCH ====")
	return foundIndex < length && neighbors[foundIndex] == v
}
