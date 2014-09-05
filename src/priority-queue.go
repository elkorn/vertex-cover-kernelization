package graph

import "container/heap"

// An pqItem is something we manage in a priority queue.
type pqItem struct {
	value *lpNode // The value of the item; arbitrary.
	// priority int     // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A priorityQueue implements heap.Interface and holds pqItems.
type priorityQueue []*pqItem

type PriorityQueueProxy struct {
	pq *priorityQueue
}

func MkPriorityQueue() *PriorityQueueProxy {
	result := &PriorityQueueProxy{
		pq: new(priorityQueue),
	}

	heap.Init(result.pq)
	return result
}

func (self *PriorityQueueProxy) Push(node *lpNode) {
	heap.Push(self.pq, &pqItem{
		value: node,
	})
}

func (self *PriorityQueueProxy) Pop() *lpNode {
	result := heap.Pop(self.pq).(*pqItem)
	return result.value
}

func (self *PriorityQueueProxy) Empty() bool {
	return self.pq.Empty()
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq priorityQueue) Less(i, j int) bool {
	// If the nodes have the same lower bound (cost),
	// take the one at the deeper level.

	if pq[i].value.lowerBound == pq[j].value.lowerBound {
		return pq[i].value.level >= pq[j].value.level
	}

	// Nodes with the lower cost have priority.
	return pq[i].value.lowerBound < pq[j].value.lowerBound
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	Debug("Push (%v:%v) %v", item.value.level, item.value.lowerBound, item.value.selection)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) PushVal(x interface{}) {
	pq.Push(&pqItem{
		value: x.(*lpNode),
	})
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	Debug("Pop (%v:%v) %v", item.value.level, item.value.lowerBound, item.value.selection)
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) PopVal() interface{} {
	return pq.Pop().(*pqItem).value
}
