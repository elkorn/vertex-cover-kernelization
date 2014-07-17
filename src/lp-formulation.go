package graph

import "math/rand"

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

type Selection map[Vertex]int

type lpNode struct {
	selection   Selection
	lowerBound  int
	left, right *lpNode
}

func computeLowerBound(g *Graph, preselected Selection) int {
	result := 0
	for _, edge := range g.Edges {
		// Maintaining the invariant: {u,v} \SUB0 E \==> Xu + Xv >= 1 (use mathematics.vim to write this correctly)
		if preselected[edge.from] < 1 && preselected[edge.to] < 1 {
			// This is stupid and temporary - `Vertex.degree` has to be implemented.
			// Select only one node, preferably with one with the larger degree.
			// Maintaining the invariant: Minimize \GS X_v
			selected := resolveConflict(g, edge.from, edge.to)
			Debug("%v vs %v -> %v", edge.from, edge.to, selected)
			// Should a copy be made here?
			preselected[selected] = 1
		}
	}

	for _, val := range preselected {
		result += val
	}

	return result
}

func objectiveFunction(feasibleSolutions []Selection) Selection {
	res := Selection{}
	minWeight := MAX_INT
	for _, solution := range feasibleSolutions {
		totalWeight := 0
		for _, weight := range solution {
			totalWeight = totalWeight + weight
		}

		if totalWeight < minWeight {
			res = solution
			minWeight = totalWeight
		}
	}

	return res
}

func resolveConflict(g *Graph, v1, v2 Vertex) Vertex {
	d1, err := g.Degree(v1)
	if nil != err {
		panic(err)
	}

	d2, err := g.Degree(v2)
	if nil != err {
		panic(err)
	}

	switch true {
	case d1 > d2:
		return v1
	case d1 < d2:
		return v2
	default:
		if rand.Intn(2) == 0 {
			return v1
		}

		return v2
	}
}

func getEndpoints(edges Edges) []Vertex {
	contains := make(map[Vertex]bool)
	result := make([]Vertex, 0)
	appendIfNotContains := func(v ...Vertex) []Vertex {
		for _, v := range v {
			if !contains[v] {
				contains[v] = true
				result = append(result, v)
			}
		}

		return result
	}

	for _, edge := range edges {
		result = appendIfNotContains(edge.from, edge.to)
	}
	return result
}

// Takes in all the edges and returns the least-costing combination according to the LP formulation.
func branchAndBound(edges Edges) []int {
	// 1. Initial value for the best combination
	// bestCombination := MAX_INT
	// // 2. Initialize a priority queue.
	// queue := PriorityQueue{}

	return nil
}
