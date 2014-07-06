package graph

import (
	"fmt"
	"log"
)

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
	// Debug("==== SEARCH =====")
	length := len(neighbors)
	// Debug("Searching for %v in %v", v, neighbors)
	foundIndex := indexOf(length, func(i int) bool {
		return neighbors[i] == v
	})

	// Debug("Found index %v", foundIndex)
	// Debug("==== END SEARCH ====")
	return foundIndex < length && neighbors[foundIndex] == v
}

func Debug(format string, args ...interface{}) {
	if options.Verbose {
		log.Print(fmt.Sprintf(format, args...))
	}
}
