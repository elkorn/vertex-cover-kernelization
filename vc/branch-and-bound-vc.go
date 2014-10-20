package vc

import "github.com/deckarep/golang-set"
import (
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type BnbNode struct {
	selection  mapset.Set
	level      int
	lowerBound int
}

var CONFLICT_RESOLVER func(*graph.Graph, int, int) bool = func(g *graph.Graph, d1, d2 int) bool {
	return d1 >= d2
}

func mkBnbNode(g *graph.Graph, selection mapset.Set, level int) *BnbNode {
	result := new(BnbNode)
	result.selection = selection
	result.lowerBound = computeLowerBound(g, selection)
	result.level = level

	// utility.Debug("New lp node on level %v: %v (%v)", result.level, result.selection, result.lowerBound)
	return result
}

func computeLowerBound(g *graph.Graph, preselected mapset.Set) int {
	fullSelection := mapset.NewSet().Union(preselected)
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		// Maintaining the invariant: {u,v} ∈ E ⇒  Xu + Xv ≥ 1
		if !(fullSelection.Contains(edge.From) || fullSelection.Contains(edge.To)) {
			// Select only one node, preferably with one with the larger degree.
			// Maintaining the invariant: Minimize \GS X_v
			selected := resolveConflict(g, edge.From, edge.To)
			fullSelection.Add(selected)
		}
	})

	return fullSelection.Cardinality()
}

func objectiveFunction(feasibleSolutions []mapset.Set) mapset.Set {
	res := mapset.NewSet()
	minWeight := utility.MAX_INT
	for _, solution := range feasibleSolutions {
		totalWeight := solution.Cardinality()
		if totalWeight < minWeight {
			res = solution
			minWeight = totalWeight
		}
	}

	return res
}

func resolveConflict(g *graph.Graph, v1, v2 graph.Vertex) graph.Vertex {
	if CONFLICT_RESOLVER(g, g.Degree(v1), g.Degree(v2)) {
		return v1
	}

	return v2
}

func getEdgeEndpoints(g *graph.Graph) mapset.Set {
	// TODO: Refactor to not use the containment map.
	result := mapset.NewSet()
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		result.Add(edge.From)
		result.Add(edge.To)
	})

	return result
}

// Similar to graph.Vertex.degree -> this should be push-based while computing the lower bound.
func getNumberOfCoveredEdges(g *graph.Graph, s mapset.Set) int {
	covered := mapset.NewSet()
	for val := range s.Iter() {
		vertex := val.(graph.Vertex)
		// utility.Debug("Vertex %v", vertex)
		g.ForAllNeighbors(vertex, func(edge *graph.Edge, done chan<- bool) {
			covered.Add(edge)
			// utility.Debug("\t covers %v-%v (%v)", edge.From, edge.To, covered.Cardinality())
		})
	}

	result := covered.Cardinality()
	utility.Debug("Edges covered: %v/%v", result, g.NEdges())
	return result
}

// Takes in all the edges and returns the least-costing combination
// according to the LP formulation.
func branchAndBound(g *graph.Graph) mapset.Set {
	return BranchAndBound(g, nil, utility.MAX_INT)
}

func BranchAndBound(g *graph.Graph, halt chan<- bool, k int) mapset.Set {
	// 1. Initial value for the best combination
	bestSelection := mapset.NewSet()
	n := g.NEdges()
	total, worked := 0, 0
	// 2. Initialize a priority queue.
	// The size of the priority queue would be calculated as in ../combinations.go
	// TODO: Benchmark if it is worth it to pre-calculate the queue capacity.
	queue := MkPriorityQueue()
	vertices := getEdgeEndpoints(g)
	utility.Debug("Edge endpoints: %v", vertices)
	// 3. Generate the first node with initial selection and compute its lower bound.
	// 4. Insert the node into the PQ.
	queue.Push(mkBnbNode(g, bestSelection, 0))
	total++
	bestLowerBound := utility.MAX_INT
	// 5. while there is something in the PQ
	for !queue.Empty() {
		// 6. Remove the first element from the PQ and assign it to the parent node.
		node := queue.Pop()
		utility.Debug("Working ----------" /*, node.selection*/)
		utility.Debug("Lower bound: %v vs %v vs k %v", node.lowerBound, bestLowerBound, k)
		// 7. If the lower bound is better then the current one...
		if node.lowerBound < bestLowerBound {
			worked++
			// 8. Set the new level to a parent's + 1.
			// 9. If all edges are covered...
			if node.level == n {
				// utility.Debug("Covers all edges.")
				// 10. Compute the cost of the combo.
				// 11. Set the current lower bound as the best one.
				bestLowerBound = node.lowerBound
				// 12. Set the current selection as the best one.
				bestSelection = node.selection
				utility.Debug("New best selection - %v elements", bestSelection.Cardinality())
				utility.Debug("New lower bound - %v", bestLowerBound)
				if k != utility.MAX_INT && bestLowerBound <= k {
					break
				}
			} else { // 13. If not (9.)...
				// 14. For all vertices v such that v is not in the selection of the parent...
				// utility.Debug("Does not cover all edges.")
				for vInter := range vertices.Iter() {
					v := vInter.(graph.Vertex)
					if node.selection.Contains(v) {
						continue
					}

					// 15. Copy the parent selection to new node
					newSelection := node.selection.Clone()
					// 16. Add v to the selection.
					newSelection.Add(v)
					// 17. Compute the lower bound.
					newNode := mkBnbNode(g, newSelection, node.level)
					// 18. If the new lower bound is better...
					// utility.Debug("new selection: %v", newNode.selection)
					// utility.Debug("lower bound %v vs %v", newNode.lowerBound, bestLowerBound)
					total++
					if newNode.lowerBound < bestLowerBound {
						// utility.Debug("Looks good, pushing into the queue.")
						newNode.level = getNumberOfCoveredEdges(g, newSelection)
						// 19. Insert the node into the priority queue.
						queue.Push(newNode)
					}
				}
			}
		} else {
			utility.Debug("Omitting.")
		}
	}

	utility.Debug("For %v edges, %v vertices:\n", g.NEdges(), g.NVertices())
	utility.Debug("Worked through %3.2f%% (%v/%v) solutions\n", (float64(worked)/float64(total))*100, worked, total)

	if bestLowerBound > k {
		utility.Debug("Cannot find a vertex cover of size \\leq k.")
		utility.Debug("Best lower bound: %v, cardinality: %v.", bestLowerBound, bestSelection.Cardinality())
		halt <- true
		return nil
	}

	utility.Debug("Best selection (%v elements) satisfying k: %v\n", bestSelection.Cardinality(), bestSelection)
	return bestSelection
}
