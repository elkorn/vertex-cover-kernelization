package graph

import (
	"container/heap"

	"github.com/deckarep/golang-set"
)

type structurePriority int

func (s *structure) computePriority(g *Graph) structurePriority {
	/*
		This will reflect the ideas from the paper:
		1 Γ is a strong 2-tuple.
		2 Γ is a 2-tuple.
		3 Γ is a good pair ( u , z ) where d ( u ) = 3 and the neighbors of u are degree-5 vertices such that no two of them share any
		common neighbors besides u.
		4 Γ is a good pair ( u , z ) where d ( u ) = 3 and d ( z ) ≥ 5.
		5 Γ is a good pair ( u , z ) where d ( u ) = 3 and d ( z ) ≥ 4.
		6 Γ is a good pair ( u , z ) where d ( u ) = 4, u has at least three degree-5 neighbors, and the graph induced by N ( u ) contains
		at least one edge (i.e., there is at least one edge among the neighbors of u).
		7 Γ is a good pair ( u , z ) where d ( u ) = 4 and all the neighbors of u are degree-5 vertices such that no two of them share a
		neighbor other than u.
		8 Γ is a vertex z with d ( z ) ≥ 8.
		9 Γ is a good pair ( u , z ) where d ( u ) = 4 and d ( z ) ≥ 5.
		10 Γ is a good pair ( u , z ) where d ( u ) = 5 and d ( z ) ≥ 6.
		11 Γ is a vertex z such that d ( z ) ≥ 7.
		12 Γ is any good pair other than the ones appearing in 1–11 above.
	*/
	cardinality := s.S.Cardinality()
	elements := make([]Vertex, cardinality)
	i := 0
	for elem := range s.S.Iter() {
		elements[i] = elem.(Vertex)
		i++
	}

	switch cardinality {
	case 1:
		deg := g.Degree(elements[0])
		if deg >= 8 {
			// 8 Γ is a vertex z with d ( z ) ≥ 8.
			return 8
		} else if deg >= 7 {
			// 11 Γ is a vertex z such that d ( z ) ≥ 7.
			return 11
		}

	case 2:
		// It's a good pair.
		u, z := elements[0], elements[1]
		du := g.Degree(u)
		dz := g.Degree(z)
		// The tuple case will most likely not have to be checked.
		/*
			A tuple ( S , q ) , where S = { u , v} , is called a 2-tuple if it
			satisfies the following conditions:
			(1) q = 1,
			(2) d (u) ≥ d (z) ≥ 1,
			(3) u and z are non-adjacent.
		*/
		if s.q == 1 &&
			dz >= 1 && du >= dz &&
			!g.HasEdge(u, z) {
			// It's a 2-tuple.
			/*
				A 2-tuple ({ u , v}, 1 ) is a strong-2-tuple if it satisfies the
				additional condition:
				d ( u ) ≥ 4 and d (v) ≥ 4, or 2 ≤ d ( u ) ≤ 3 and 2 ≤ d (v) ≤ 3.
			*/
			if du >= 4 && dz >= 4 ||
				du >= 2 && du <= 3 &&
					dz >= 2 && dz <= 3 {
				// It's a strong 2-tuple.
				return 1
			}

			return 2
		}

		degree5NeighborsCount, hasOnlyDegree5Neighbors := s.countDegree5Neighbors(u, g)
		if du == 3 || du == 4 {
			neighborsShareCommonNeighborOtherThanU, neighborsAreDisjoint := s.neighborsOfUShareCommonVertexOtherThanU(u, z, g)
			if du == 3 {
				if hasOnlyDegree5Neighbors &&
					!neighborsShareCommonNeighborOtherThanU {
					// 3 Γ is a good pair ( u , z ) where d ( u ) = 3 and the
					// neighbors of u are degree-5 vertices such that no two of
					// them share any common neighbors besides u.
					return 3
				}

				if dz >= 5 {
					// 4 Γ is a good pair (u, z) where d(u) = 3 and d(z) ≥ 5.
					return 4
				}

				if dz >= 4 {
					// 5 Γ is a good pair (u, z) where d(u) = 3 and d(z) ≥ 4.
					return 5
				}
			}

			if du == 4 {
				if hasOnlyDegree5Neighbors &&
					!neighborsShareCommonNeighborOtherThanU {
					// 7 Γ is a good pair ( u , z ) where d ( u ) = 4 and all
					// the neighbors of u are degree-5 vertices such that no
					// two of them share a neighbor other than u.
					return 7
				}

				if degree5NeighborsCount >= 3 {
					if !neighborsAreDisjoint {
						// 6 Γ is a good pair (u , z) where d(u) = 4, u has at least
						// 3 degree-5 neighbors, and there is at least one edge
						// among the neighbors of u.
						return 6
					}
				}

				if dz >= 5 {
					// 9 Γ is a good pair (u, z) where d(u) = 4 and d(z) ≥ 5.
					return 9
				}
			}
		}

		if du == 5 && dz >= 6 {
			// 10 Γ is a good pair (u,  z) where d(u) = 5 and d(z) ≥ 6.
			return 10
		}

		// 12 Γ is any good pair other than the ones appearing in 1–11 above.
		return 12
	}

	panic("Unrecognized structure!")
}

func (self *goodPair) computePriority(g *Graph) structurePriority {
	return self.pair.computePriority(g)
}

type StructurePriorityQueueProxy struct {
	pq                    *StructurePriorityQueue
	numberOfStrong2Tuples int
}

func MkStructurePriorityQueue() *StructurePriorityQueueProxy {
	result := &StructurePriorityQueueProxy{
		pq: new(StructurePriorityQueue),
	}

	heap.Init(result.pq)
	return result
}

func (self *StructurePriorityQueueProxy) Push(node *structure, g *Graph) {
	priority := node.computePriority(g)
	if 1 == priority {
		self.numberOfStrong2Tuples++
		node.q = 1
	}

	heap.Push(self.pq, &structurePqItem{
		value:    node,
		priority: node.computePriority(g),
	})

}

func (self *StructurePriorityQueueProxy) Pop() (*structure, structurePriority) {
	result := heap.Pop(self.pq).(*structurePqItem)
	if 1 == result.priority {
		self.numberOfStrong2Tuples--
	}

	return result.value, result.priority
}

func (self *StructurePriorityQueueProxy) Empty() bool {
	return self.pq.Empty()
}

func (self *StructurePriorityQueueProxy) ContainsStrong2Tuple() bool {
	return self.numberOfStrong2Tuples > 0
}

func (self *StructurePriorityQueueProxy) PopAllStrong2Tuples() mapset.Set {
	result := mapset.NewSet()
	for self.ContainsStrong2Tuple() {
		s2t, _ := self.Pop()
		result.Add(s2t)
	}

	return result
}

type structurePqItem struct {
	value    *structure
	index    int
	priority structurePriority
}

// A StructurePriorityQueue implements heap.Interface and holds structurePqItems.
type StructurePriorityQueue []*structurePqItem

func (pq StructurePriorityQueue) Len() int { return len(pq) }

func (pq StructurePriorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq StructurePriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq StructurePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *StructurePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*structurePqItem)
	Debug("Push (%v) %v", item.value, item.priority)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *StructurePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	Debug("Pop (%v) %v", item.value, item.priority)
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
