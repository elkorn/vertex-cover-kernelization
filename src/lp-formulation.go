package graph

import "math/rand"

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

func objectiveFunction(feasibleSolutions []map[Vertex]int) map[Vertex]int {
	res := make(map[Vertex]int)
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

func resolveConflict(n1, n2 Node) Node {
	switch true {
	case n1.degree > n2.degree:
		return n1
	case n1.degree < n2.degree:
		return n2
	default:
		if rand.Intn(2) == 0 {
			return n1
		}

		return n2
	}
}

// Takes in all the edges and returns the least-costing combination according to the LP formulation.
func branchAndBound(edges Edges) []int {
	return nil
}
