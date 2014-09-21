package graph

import "container/heap"

type tuplePqItem struct {
	value    *tuple
	index    int
	priority int
}

// A tuplePriorityQueue implements heap.Interface and holds tuplePqItems.
type tuplePriorityQueue []*tuplePqItem

type TuplePriorityQueueProxy struct {
	pq *tuplePriorityQueue
}

func MkTuplePriorityQueue() *TuplePriorityQueueProxy {
	result := &TuplePriorityQueueProxy{
		pq: new(tuplePriorityQueue),
	}

	heap.Init(result.pq)
	return result
}

func (self *TuplePriorityQueueProxy) Push(node *tuple) {
	heap.Push(self.pq, &tuplePqItem{
		value: node,
	})
}

func (self *TuplePriorityQueueProxy) Pop() *tuple {
	result := heap.Pop(self.pq).(*tuplePqItem)
	return result.value
}

func (self *TuplePriorityQueueProxy) Empty() bool {
	return self.pq.Empty()
}

func (pq tuplePriorityQueue) Len() int { return len(pq) }

func (pq tuplePriorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq tuplePriorityQueue) Less(i, j int) bool {
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

func (pq tuplePriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *tuplePriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*tuplePqItem)
	Debug("Push (%v) %v", item.value, item.priority)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *tuplePriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	Debug("Pop (%v) %v", item.value, item.priority)
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
