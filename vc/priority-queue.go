package vc

import (
	"container/heap"

	"github.com/elkorn/vertex-cover-kernelization/utility"
)

// An pqItem is something we manage in a priority queue.
type pqItem struct {
	value *BnbNode // The value of the item; arbitrary.
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

func (self *PriorityQueueProxy) Push(node *BnbNode) {
	heap.Push(self.pq, &pqItem{
		value: node,
	})
}

func (self *PriorityQueueProxy) Pop() *BnbNode {
	result := heap.Pop(self.pq).(*pqItem)
	return result.value
}

func (self *PriorityQueueProxy) Empty() bool {
	return self.pq.Empty()
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq priorityQueue) Less(i, j int) bool {
	// If the nodes are at the same level, take the one with lower cost.
	if pq[i].value.level == pq[j].value.level {
		return pq[i].value.lowerBound < pq[j].value.lowerBound
	}

	// Nodes on a deeper level have priority.
	return pq[i].value.level > pq[j].value.level
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	if item.value.selection == nil {
		utility.Debug("Pop (%v:%v) nil selection", item.value.level, item.value.lowerBound)
	} else {
		utility.Debug("Push (%v:%v) %v elements", item.value.level, item.value.lowerBound, item.value.selection.Cardinality())

	}
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) PushVal(x interface{}) {
	pq.Push(&pqItem{
		value: x.(*BnbNode),
	})
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	if item.value.selection == nil {
		utility.Debug("Pop (%v:%v) nil selection", item.value.level, item.value.lowerBound)
	} else {
		utility.Debug("Pop (%v:%v) %v elements", item.value.level, item.value.lowerBound, item.value.selection.Cardinality())
	}
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) PopVal() interface{} {
	return pq.Pop().(*pqItem).value
}
