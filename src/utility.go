package graph

import (
	"fmt"
	"log"
	"math"

	"github.com/deckarep/golang-set"
)

func removeAt(source Edges, position int) Edges {
	return append(source[:position], source[position+1:]...)
}

func sliceIndexOf(n int, f func(int) bool) int {
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
	foundIndex := sliceIndexOf(length, func(i int) bool {
		return neighbors[i] == v
	})

	// Debug("Found index %v", foundIndex)
	// Debug("==== END SEARCH ====")
	return foundIndex < length && neighbors[foundIndex] == v
}

func mkBoolMatrix(n, cap int) [][]bool {
	result := make([][]bool, n, cap)
	for i := range result {
		result[i] = make([]bool, n, cap)
	}

	return result
}

func PrintSet(set mapset.Set) {
	for s := range set.Iter() {
		Debug("%v", s)
	}
}

func Debug(format string, args ...interface{}) {
	if options.Verbose {
		log.Print(fmt.Sprintf(format, args...))
	}
}

func inVerboseContext(fn func()) {
	SetOptions(Options{Verbose: true})
	fn()
	SetOptions(Options{Verbose: false})

}

func rng(args ...int) []int {
	c := make([]int, len(args))
	copy(c, args)
	return c
}

func IntAbs(val int) int {
	return int(math.Abs(float64(val)))
}
