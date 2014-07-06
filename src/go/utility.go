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

func mkGraph1() *Graph {
	/*
		   1o---o2
			|\ /|
			| o5|
			|/ \|
		   4o---o3
	*/
	g := mkGraphWithVertices(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	return g
}

func mkGraph2() *Graph {
	/*
		   1o--------o2
			|\      /|
			|5o----o6|
			| |    | |
			|8o----o7|
			|/      \|
		   4o--------o3
	*/
	g := mkGraphWithVertices(8)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 8)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(8, 5)
	return g
}

func mkGraph3() *Graph {
	/*
	           1
	          / \
	     3---+   +---2
	    / \         / \
	   7   6       5   4
	*/
	g := mkGraphWithVertices(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	return g
}

func mkGraph4() *Graph {
	/*
	           1
	          / \
	     3---+   +---2
	    / \         / \
	   7---6       5---4
	*/

	g := mkGraph3()

	g.AddEdge(6, 7)
	g.AddEdge(4, 5)

	return g
}

func mkGraph5() *Graph {
	/*
		  1   6
		 / \ / \
		3   2   7
		   / \
		  5---4
	*/

	g := mkGraphWithVertices(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(4, 5)
	g.AddEdge(6, 7)

	return g
}
