package graph

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &pqItem{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &pqItem{
		value:    "orange",
		priority: 1,
	}

	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	previouspqItem := &pqItem{
		value:    "dummy",
		priority: 999,
		index:    100,
	}
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		assert.True(t, item.priority < previouspqItem.priority)
		previouspqItem = item
	}
}

func TestPopVal(t *testing.T) {
	q1 := PriorityQueue{}
	q2 := PriorityQueue{}
	item1 := &pqItem{
		value:    "banana",
		priority: 1,
	}
	item2 := item1
	q1.Push(item1)
	q2.Push(item2)
	assert.Equal(t, q1.Pop().(*pqItem).value, q2.PopVal())
}
