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

func computeLowerBound(g *Graph, preselection map[Vertex]int) int {
	result := 0
	mkNode := func(v Vertex) Node {
		degree, err := g.Degree(v)
		if nil != err {
			panic(err)
		}

		return Node{int(v), degree}
	}

	for _, edge := range g.Edges {
		// Maintaining the invariant: {u,v} \SUB0 E \==> Xu + Xv >= 1 (use mathematics.vim to write this correctly)
		if preselection[edge.from] < 1 && preselection[edge.to] < 1 {
			// This is stupid and temporary - `Vertex.degree` has to be implemented.
			n1 := mkNode(edge.from)
			n2 := mkNode(edge.to)
			// Select only one node, preferably with one with the larger degree.
			// Maintaining the invariant: Minimize \GS X_v
			selected := Vertex(resolveConflict(n1, n2).id)
			Debug("%v vs %v -> %v", n1, n2, selected)
			// Should a copy be made here?
			preselection[selected] = 1
		}
	}

	for _, val := range preselection {
		result += val
	}

	return result
}

// Takes in all the edges and returns the least-costing combination according to the LP formulation.
func branchAndBound(edges Edges) []int {
	return nil
}
