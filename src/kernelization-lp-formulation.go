package graph

import (
	"math/rand"

	"github.com/deckarep/golang-set"
)

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

type lpNode struct {
	selection  mapset.Set
	level      int
	lowerBound int
}

func mkLpNode(g *Graph, selection mapset.Set, level int) *lpNode {
	result := new(lpNode)
	result.selection = selection
	result.lowerBound = computeLowerBound(g, selection)
	result.level = level

	Debug("New lp node on level %v", result.level)
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

func (self *Graph) getEdgeEndpoints() mapset.Set {
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
		g.ForAllNeighbors(vertex, func(edge *Edge, done chan<- bool) {
			covered.Add(edge)
		})
	}

	result := covered.Cardinality()
	Debug("Edges covered: %v/%v", result, g.NEdges())
	return result
}

// Takes in all the edges and returns the least-costing combination
// according to the LP formulation.
func branchAndBound(g *Graph) mapset.Set {
	// 1. Initial value for the best combination
	bestSelection := mapset.NewSet()
	n := g.NEdges()
	// 2. Initialize a priority queue.
	queue := PriorityQueue{}
	vertices := g.getEdgeEndpoints()
	Debug("Edge endpoints: %v", vertices)
	selection := bestSelection
	// 3. Generate the first node with initial selection and compute its lower bound.
	// 4. Insert the node into the PQ.
	queue.PushVal(mkLpNode(g, selection, 0))
	bestLowerBound := MAX_INT
	// 5. while there is something in the PQ
	for !queue.Empty() {
		// time.Sleep(500 * time.Millisecond)
		// 6. Remove the first element from the PQ and assign it to the parent node.
		node := queue.PopVal().(*lpNode)
		Debug("Working %v ----------", node.selection)
		Debug("Lower bound: %v vs %v", node.lowerBound, bestLowerBound)
		// 7. If the lower bound is better then the current one...
		if node.lowerBound < bestLowerBound {
			// 8. Set the new level to a parent's + 1.
			newLevel := node.level + 1
			selection = node.selection
			// 9. If this level equals the number of vertices - 1...
			nCoveredEdges := getNumberOfCoveredEdges(g, selection)
			if nCoveredEdges == n {
				Debug("Covers all edges.")
				// 10. Compute the cost of the combo.
				// 11. Set the current lower bound as the best one.
				bestLowerBound = node.lowerBound
				// 12. Set the current selection as the best one.
				bestSelection = selection
				Debug("New best selection - %v", bestSelection)
				Debug("New lower bound - %v", bestLowerBound)
			} else { // 13. If not (9.)...
				// 14. For all vertices v such that v is not in the selection of the parent...
				Debug("Does not cover all edges.")
				for vInter := range vertices.Iter() {
					v := vInter.(Vertex)
					if selection.Contains(v) {
						continue
					}

					// 15. Copy the parent selection to new node
					newSelection := selection.Clone()
					// 16. Add v to the selection.
					newSelection.Add(v)
					// 17. Compute the lower bound.
					newNode := mkLpNode(g, newSelection, newLevel)
					// 18. If the new lower bound is better...
					Debug("new selection: %v", newNode.selection)
					Debug("lower bound %v vs %v", newNode.lowerBound, bestLowerBound)
					if newNode.lowerBound < bestLowerBound {
						Debug("Looks good, pushing into the queue.")
						// 19. Insert the node into the priority queue.
						queue.PushVal(newNode)
					}
				}
			}
		}
	}

	Debug("Best selection: %v", bestSelection)
	return bestSelection
}
