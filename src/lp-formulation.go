package graph

import "math/rand"

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

type Selection map[Vertex]int

type lpNode struct {
	selection  Selection
	level      int
	lowerBound int
}

func mkLpNode(g *Graph, selection Selection) *lpNode {
	result := new(lpNode)
	result.selection = selection
	result.lowerBound = computeLowerBound(g, selection)

	return result
}

func computeLowerBound(g *Graph, preselected Selection) int {
	result := 0
	for _, edge := range g.Edges {
		// Maintaining the invariant: {u,v} \SUB0 E \==> Xu + Xv >= 1 (use mathematics.vim to write this correctly)
		if preselected[edge.from] < 1 && preselected[edge.to] < 1 {
			// Select only one node, preferably with one with the larger degree.
			// Maintaining the invariant: Minimize \GS X_v
			selected := resolveConflict(g, edge.from, edge.to)
			Debug("%v vs %v -> %v", edge.from, edge.to, selected)
			// Should a copy be made here?
			preselected[selected] = 1
		}
		// else -> numberOfCoveredEdges += 1
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

// Similar to Vertex.degree -> this should be push-based while computing the lower bound.
func getNumberOfCoveredEdges(g *Graph, s Selection) int {
	result := 0
	for val := range s {
		vertex := Vertex(val)
		Debug("Vertex: %v", vertex)
		for _, edge := range g.Edges {
			if edge.from == vertex || edge.to == vertex {
				result++
			}
		}
	}
	return result
}

// Takes in all the edges and returns the least-costing combination
// according to the LP formulation.
func branchAndBound(g *Graph) []int {
	// 1. Initial value for the best combination
	bestLowerBound := MAX_INT
	bestSelection := Selection{}
	n := len(g.Vertices)
	// 2. Initialize a priority queue.
	queue := PriorityQueue{}
	vertices := getEndpoints(g.Edges)
	selection := Selection{vertices[0]: 1}
	// 3. Generate the first node with vertex [1] and compute its lower bound.
	// 4. Insert the node into the PQ.
	queue.Push(mkLpNode(g, selection))
	// 5. while there is something in the PQ
	for !queue.Empty() {
		// 6. Remove the first element from the PQ and assign it to the parent node.
		node := queue.PopVal().(*lpNode)

		// 7. If the lower bound is better then the current one...
		if node.lowerBound < bestLowerBound {
			// 8. Set the new level to a parent's + 1.
			newLevel := node.level + 1
			selection := node.selection
			// 9. If this level equals the number of vertices - 1...
			// if newLevel == n-1 {
			// // This condition is OK for TSP, has to be changed for this formulation.
			// }
			// This is my proposition for the condition. Let's see if it makes sense...
			// if len(getCoveredEdges(g, selection)) == n {
			// 	// 10. Compute the cost of the combo.
			// 	// ...
			// }
			for _, vertex := range vertices {
				if selection[vertex] != 0 {
					continue
				}
			}
		}

	}

	return nil
}
