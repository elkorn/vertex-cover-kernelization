package graph

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelPriority(t *testing.T) {
	// Some items and their priorities.
	n1 := new(lpNode)
	n2 := new(lpNode)
	n3 := new(lpNode)
	n4 := new(lpNode)
	n1.level = 1
	n2.level = 2
	n3.level = 3
	n4.level = 4
	items := [3]*lpNode{n1, n2, n3}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for _, value := range items {
		pq[i] = &pqItem{
			value: value,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &pqItem{
		value: n4,
	}

	heap.Push(&pq, item)
	pq.update(item, item.value)
	previouspqItem := &pqItem{
		value: &lpNode{level: 1000},
		index: 100,
	}

	// Take the items out; they arrive in decreasing priority order by level.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*pqItem)
		assert.True(t, item.value.level < previouspqItem.value.level, "Deeper level should be prioritized.")
		previouspqItem = item
	}
}

func TestLowerBoundPriority(t *testing.T) {
	n1 := &lpNode{level: 1, lowerBound: 4}
	n2 := &lpNode{level: 1, lowerBound: 3}
	n3 := &lpNode{level: 1, lowerBound: 2}
	n4 := &lpNode{level: 1, lowerBound: 1}
	items := [3]*lpNode{n1, n2, n3}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for _, value := range items {
		pq[i] = &pqItem{
			value: value,
			index: i,
		}
		i++
	}
	heap.Init(&pq)

	item := &pqItem{
		value: n4,
	}

	heap.Push(&pq, item)
	previouspqItem := &pqItem{
		value: &lpNode{level: 1, lowerBound: 0},
		index: 100,
	}

	// Take the items out, they arrive in an increasing order by lower bound.
	inVerboseContext(func() {
		for pq.Len() > 0 {
			item := heap.Pop(&pq).(*pqItem)
			Debug("Popped %v (previous %v) -> %v", item.value.lowerBound, previouspqItem.value.lowerBound, item.value.lowerBound > previouspqItem.value.lowerBound)
			assert.True(t, item.value.lowerBound > previouspqItem.value.lowerBound, "Smaller lower bound should be prioritized.")
			previouspqItem = item
		}
	})
}

func TestPopVal(t *testing.T) {
	q1 := PriorityQueue{}
	q2 := PriorityQueue{}
	item1 := &pqItem{
		value: &lpNode{level: 1},
	}
	item2 := item1
	q1.Push(item1)
	q2.Push(item2)
	assert.Equal(t, q1.Pop().(*pqItem).value, q2.PopVal())
}
