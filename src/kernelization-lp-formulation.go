package graph

import "math/rand"

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

type Selection map[Vertex]int

func (self Selection) Copy() Selection {
	result := Selection{}
	for k, v := range self {
		result[k] = v
	}

	return result
}

type lpNode struct {
	selection  Selection
	level      int
	lowerBound int
}

func mkLpNode(g *Graph, selection Selection, level int) *lpNode {
	result := new(lpNode)
	result.selection = selection
	result.lowerBound = computeLowerBound(g, selection)
	result.level = level

	Debug("New lp node\n\tselection: %v\n\tcost: %v\n\tlevel: %v", result.selection, result.lowerBound, result.level)
	return result
}

func computeLowerBound(g *Graph, preselected Selection) int {
	result := 0
	g.ForAllEdges(func(edge *Edge, _ int, done chan<- bool) {
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
	})

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

	Debug("Resolving conflict v1(%v) vs v2(%v)", d1, d2)

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

func (self *Graph) getEdgeEndpoints() Vertices {
	// TODO: Refactor to not use the containment map.
	result := make(Vertices, 0, self.currentVertexIndex)
	contains := make([]bool, self.currentVertexIndex)
	appendIfNotContains := func(vs ...Vertex) Vertices {
		for _, v := range vs {
			vi := v.toInt()
			if !contains[vi] {
				contains[vi] = true
				result = append(result, v)
			}
		}

		return result
	}

	self.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		result = appendIfNotContains(edge.from, edge.to)
	})

	return result
}

// Similar to Vertex.degree -> this should be push-based while computing the lower bound.
func getNumberOfCoveredEdges(g *Graph, s Selection) int {
	result := 0
	covered := make(map[int]bool)
	Debug("Selection %v", s)
	for val := range s {
		vertex := Vertex(val)
		g.ForAllEdges(func(edge *Edge, i int, done chan<- bool) {
			if !covered[i] && (edge.from == vertex || edge.to == vertex) {
				result++
				covered[i] = true
				Debug("%v covers %v -> %v", vertex, edge, result)
			}
		})
	}
	return result
}

// Takes in all the edges and returns the least-costing combination
// according to the LP formulation.
func branchAndBound(g *Graph) Selection {
	// 1. Initial value for the best combination
	bestSelection := Selection{}
	n := g.NEdges()
	// 2. Initialize a priority queue.
	queue := PriorityQueue{}
	vertices := g.getEdgeEndpoints()
	Debug("Edge endpoints: %v", vertices)
	selection := Selection{}
	// 3. Generate the first node with initial selection and compute its lower bound.
	// 4. Insert the node into the PQ.
	queue.PushVal(mkLpNode(g, selection, 0))
	bestLowerBound := MAX_INT
	// 5. while there is something in the PQ
	for !queue.Empty() {
		// 6. Remove the first element from the PQ and assign it to the parent node.
		node := queue.PopVal().(*lpNode)
		Debug("Working %v", node)
		// 7. If the lower bound is better then the current one...
		if node.lowerBound < bestLowerBound {
			Debug("Has better lower bound (%v < %v)", node.lowerBound, bestLowerBound)
			// 8. Set the new level to a parent's + 1.
			newLevel := node.level + 1
			selection := node.selection
			// 9. If this level equals the number of vertices - 1...
			// if newLevel == n-1 {
			// // This condition is OK for TSP, has to be changed for this formulation.
			// }
			// This is my proposition for the condition. Let's see if it makes sense...
			nCoveredEdges := getNumberOfCoveredEdges(g, selection)
			Debug("Covers %v edges", nCoveredEdges)
			// return bestSelection
			if nCoveredEdges == n {
				Debug("Covers all edges.")
				// 10. Compute the cost of the combo.
				// 11. Set the current lower bound as the best one.
				bestLowerBound = node.lowerBound
				Debug("New lower bound - %v", bestLowerBound)
				// 12. Set the current selection as the best one.
				bestSelection = selection
				Debug("New best selection - %v", bestSelection)
			} else { // 13. If not (9.)...
				// 14. For all vertices v such that v is not in the selection of the parent...
				Debug("Does not cover all edges.")
				for _, v := range vertices {
					if selection[v] != 0 {
						continue
					}

					// 15. Copy the parent selection to new node
					newSelection := bestSelection.Copy()
					// 16. Add v to the selection.
					newSelection[v] = 1
					// 17. Compute the lower bound.
					newNode := mkLpNode(g, newSelection, newLevel)
					// 18. If the new lower bound is better...
					Debug("Checking lower bound for %v", newNode)
					if newNode.lowerBound < bestLowerBound {
						Debug("Looks good, pushing into the queue.")
						// 19. Insert the node into the priority queue.
						queue.PushVal(newNode)
					}
				}
			}
		}

	}

	return bestSelection
}
