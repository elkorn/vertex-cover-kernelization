package graph

import "container/heap"

// An pqItem is something we manage in a priority queue.
type pqItem struct {
	value    *lpNode // The value of the item; arbitrary.
	priority int     // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds pqItems.
type PriorityQueue []*pqItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PopVal() interface{} {
	return pq.Pop().(*pqItem).value
}

// update modifies the priority and value of an pqItem in the queue.
func (pq *PriorityQueue) update(item *pqItem, value *lpNode, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
