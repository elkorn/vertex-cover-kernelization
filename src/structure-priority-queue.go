package graph

import "container/heap"

type structurePriority int

func (s structure) computePriority(g *Graph) structurePriority {
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
		u, v := elements[0], elements[1]
		du := g.Degree(u)
		dv := g.Degree(v)
		// The tuple case will most likely not have to be checked.
		/*
			A tuple ( S , q ) , where S = { u , v} , is called a 2-tuple if it
			satisfies the following conditions:
			(1) q = 1,
			(2) d ( u ) ≥ d (v) ≥ 1,
			(3) u and v are non-adjacent.
		*/
		if s.q == 1 &&
			dv >= 1 && du >= dv &&
			!g.hasEdge(u, v) {
			// It's a 2-tuple.
			/*
				A 2-tuple ({ u , v}, 1 ) is a strong-2-tuple if it satisfies the
				additional condition:
				d ( u ) ≥ 4 and d (v) ≥ 4, or 2 ≤ d ( u ) ≤ 3 and 2 ≤ d (v) ≤ 3.
			*/
			if du >= 4 && dv >= 4 ||
				du >= 2 && du <= 3 &&
					dv >= 2 && dv <= 3 {
				// It's a strong 2-tuple.
				return 1
			}

			return 2
		}

		hasOnlyDegree5Neighbors := true
		degree5NeighborsCount := 0
		g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
			if g.Degree(getOtherVertex(u, e)) == 5 {
				degree5NeighborsCount++
			} else {
				hasOnlyDegree5Neighbors = false
			}
		})

		neighborsAreDisjoint := true
		g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
			v1 := getOtherVertex(u, e)
			g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
				v2 := getOtherVertex(u, e)
				if v1 == v2 {
					return
				}

				if g.hasEdge(v1, v2) {
					neighborsAreDisjoint = false
					done <- true
				}
			})
		})
		if du == 3 {
			if hasOnlyDegree5Neighbors {
				if neighborsAreDisjoint {
					// TODO: Possible bug. Does neighbors not sharing common
					// neighbors mean that they are disjoint?

					// 3 Γ is a good pair ( u , z ) where d ( u ) = 3 and the
					// neighbors of u are degree-5 vertices such that no two of
					// them share any common neighbors besides u.
					return 3
				}

				neighborsShareCommonVertexOtherThanU := false
				g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
					if neighborsShareCommonVertexOtherThanU {
						done <- true
						return
					}

					v1 := getOtherVertex(u, e)

					g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
						if neighborsShareCommonVertexOtherThanU {
							done <- true
							return
						}

						v2 := getOtherVertex(u, e)
						if v1 == v2 {
							return
						}

						g.ForAllNeighbors(v1, func(e *Edge, done chan<- bool) {
							if neighborsShareCommonVertexOtherThanU {
								done <- true
								return
							}

							n1 := getOtherVertex(v1, e)

							g.ForAllNeighbors(v2, func(e *Edge, done chan<- bool) {
								n2 := getOtherVertex(v2, e)
								if n1 == n2 && n1 != u {
									neighborsShareCommonVertexOtherThanU = true
									done <- true
								}
							})
						})
					})
				})
				if !neighborsShareCommonVertexOtherThanU {
					// 7 Γ is a good pair ( u , z ) where d ( u ) = 4 and all
					// the neighbors of u are degree-5 vertices such that no
					// two of them share a neighbor other than u.
					return 7
				}
			}

			if dv >= 5 {
				// 4 Γ is a good pair (u, z) where d(u) = 3 and d(z) ≥ 5.
				return 4
			}

			if dv >= 4 {
				// 5 Γ is a good pair (u, z) where d(u) = 3 and d(z) ≥ 4.
				return 5
			}

		}

		if du == 4 {
			if degree5NeighborsCount >= 3 {
				if !neighborsAreDisjoint {
					// 6 Γ is a good pair (u , z) where d(u) = 4, u has at least
					// 3 degree-5 neighbors, and there is at least one edge
					// among the neighbors of u.
					return 6
				}
			}

			if dv >= 5 {
				// 9 Γ is a good pair (u, z) where d(u) = 4 and d(z) ≥ 5.
				return 9
			}

			if dv >= 6 {
				// 10 Γ is a good pair (u,  z) where d(u) = 5 and d(z) ≥ 6.
				return 10
			}
		}

		// 12 Γ is any good pair other than the ones appearing in 1–11 above.
		return 12
	}

	panic("Unrecognized structure!")
}

type structurePqItem struct {
	value    *structure
	index    int
	priority structurePriority
}

// A StructurePriorityQueue implements heap.Interface and holds structurePqItems.
type StructurePriorityQueue []*structurePqItem

type StructurePriorityQueueProxy struct {
	pq *StructurePriorityQueue
}

func MkStructurePriorityQueue() *StructurePriorityQueueProxy {
	result := &StructurePriorityQueueProxy{
		pq: new(StructurePriorityQueue),
	}

	heap.Init(result.pq)
	return result
}

func (self *StructurePriorityQueueProxy) Push(node *structure) {
	heap.Push(self.pq, &structurePqItem{
		value: node,
	})
}

func (self *StructurePriorityQueueProxy) Pop() *structure {
	result := heap.Pop(self.pq).(*structurePqItem)
	return result.value
}

func (self *StructurePriorityQueueProxy) Empty() bool {
	return self.pq.Empty()
}

func (pq StructurePriorityQueue) Len() int { return len(pq) }

func (pq StructurePriorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq StructurePriorityQueue) Less(i, j int) bool {
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
