package graph

import "github.com/deckarep/golang-set"

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

type lpNode struct {
	selection  mapset.Set
	level      int
	lowerBound int
}

var CONFLICT_RESOLVER func(*Graph, int, int) bool = func(g *Graph, d1, d2 int) bool {
	return d1 >= d2
}

func mkLpNode(g *Graph, selection mapset.Set, level int) *lpNode {
	result := new(lpNode)
	result.selection = selection
	result.lowerBound = computeLowerBound(g, selection)
	result.level = level

	// Debug("New lp node on level %v: %v (%v)", result.level, result.selection, result.lowerBound)
	return result
}

func computeLowerBound(g *Graph, preselected mapset.Set) int {
	fullSelection := mapset.NewSet().Union(preselected)
	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		// Maintaining the invariant: {u,v} ∈ E ⇒  Xu + Xv ≥ 1
		if !(fullSelection.Contains(edge.from) || fullSelection.Contains(edge.to)) {
			// Select only one node, preferably with one with the larger degree.
			// Maintaining the invariant: Minimize \GS X_v
			selected := resolveConflict(g, edge.from, edge.to)
			fullSelection.Add(selected)
		}
	})

	return fullSelection.Cardinality()
}

func objectiveFunction(feasibleSolutions []mapset.Set) mapset.Set {
	res := mapset.NewSet()
	minWeight := MAX_INT
	for _, solution := range feasibleSolutions {
		totalWeight := solution.Cardinality()
		if totalWeight < minWeight {
			res = solution
			minWeight = totalWeight
		}
	}

	return res
}

func resolveConflict(g *Graph, v1, v2 Vertex) Vertex {
	if CONFLICT_RESOLVER(g, g.Degree(v1), g.Degree(v2)) {
		return v1
	}

	return v2
}

func (self *Graph) getEdgeEndpoints() mapset.Set {
	// TODO: Refactor to not use the containment map.
	result := mapset.NewSet()
	self.ForAllEdges(func(edge *Edge, done chan<- bool) {
		result.Add(edge.from)
		result.Add(edge.to)
	})

	return result
}

// Similar to Vertex.degree -> this should be push-based while computing the lower bound.
func getNumberOfCoveredEdges(g *Graph, s mapset.Set) int {
	covered := mapset.NewSet()
	for val := range s.Iter() {
		vertex := val.(Vertex)
		// Debug("Vertex %v", vertex)
		g.ForAllNeighbors(vertex, func(edge *Edge, done chan<- bool) {
			covered.Add(edge)
			// Debug("\t covers %v-%v (%v)", edge.from, edge.to, covered.Cardinality())
		})
	}

	result := covered.Cardinality()
	Debug("Edges covered: %v/%v", result, g.NEdges())
	return result
}

// Takes in all the edges and returns the least-costing combination
// according to the LP formulation.
func branchAndBound(g *Graph) mapset.Set {
	return BranchAndBound(g, nil, MAX_INT)
}

func BranchAndBound(g *Graph, halt chan<- bool, k int) mapset.Set {
	// 1. Initial value for the best combination
	bestSelection := mapset.NewSet()
	n := g.NEdges()
	total, worked := 0, 0
	// 2. Initialize a priority queue.
	// The size of the priority queue would be calculated as in ../combinations.go
	// TODO: Benchmark if it is worth it to pre-calculate the queue capacity.
	queue := MkPriorityQueue()
	vertices := g.getEdgeEndpoints()
	Debug("Edge endpoints: %v", vertices)
	// 3. Generate the first node with initial selection and compute its lower bound.
	// 4. Insert the node into the PQ.
	queue.Push(mkLpNode(g, bestSelection, 0))
	total++
	bestLowerBound := MAX_INT
	// 5. while there is something in the PQ
	for !queue.Empty() {
		// 6. Remove the first element from the PQ and assign it to the parent node.
		node := queue.Pop()
		Debug("Working ----------" /*, node.selection*/)
		Debug("Lower bound: %v vs %v vs k %v", node.lowerBound, bestLowerBound, k)
		// 7. If the lower bound is better then the current one...
		if node.lowerBound < bestLowerBound {
			worked++
			// 8. Set the new level to a parent's + 1.
			// 9. If all edges are covered...
			if node.level == n {
				// Debug("Covers all edges.")
				// 10. Compute the cost of the combo.
				// 11. Set the current lower bound as the best one.
				bestLowerBound = node.lowerBound
				// 12. Set the current selection as the best one.
				bestSelection = node.selection
				Debug("New best selection - %v elements", bestSelection.Cardinality())
				Debug("New lower bound - %v", bestLowerBound)
				if k != MAX_INT && bestLowerBound <= k {
					break
				}
			} else { // 13. If not (9.)...
				// 14. For all vertices v such that v is not in the selection of the parent...
				// Debug("Does not cover all edges.")
				for vInter := range vertices.Iter() {
					v := vInter.(Vertex)
					if node.selection.Contains(v) {
						continue
					}

					// 15. Copy the parent selection to new node
					newSelection := node.selection.Clone()
					// 16. Add v to the selection.
					newSelection.Add(v)
					// 17. Compute the lower bound.
					newNode := mkLpNode(g, newSelection, node.level)
					// 18. If the new lower bound is better...
					// Debug("new selection: %v", newNode.selection)
					// Debug("lower bound %v vs %v", newNode.lowerBound, bestLowerBound)
					total++
					if newNode.lowerBound < bestLowerBound {
						// Debug("Looks good, pushing into the queue.")
						newNode.level = getNumberOfCoveredEdges(g, newSelection)
						// 19. Insert the node into the priority queue.
						queue.Push(newNode)
					}
				}
			}
		} else {
			Debug("Omitting.")
		}
	}

	Debug("For %v edges, %v vertices:\n", g.NEdges(), g.NVertices())
	Debug("Worked through %3.2f%% (%v/%v) solutions\n", (float64(worked)/float64(total))*100, worked, total)

	if bestLowerBound > k {
		Debug("Cannot find a vertex cover of size \\leq k.")
		Debug("Best lower bound: %v, cardinality: %v.", bestLowerBound, bestSelection.Cardinality())
		halt <- true
		return nil
	}

	Debug("Best selection (%v elements) satisfying k: %v\n", bestSelection.Cardinality(), bestSelection)
	return bestSelection
}
