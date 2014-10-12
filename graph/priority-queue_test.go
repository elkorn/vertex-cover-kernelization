package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerBoundPriority(t *testing.T) {
	n1 := &bnbNode{level: 1, lowerBound: 4}
	n2 := &bnbNode{level: 1, lowerBound: 3}
	n3 := &bnbNode{level: 1, lowerBound: 2}
	n4 := &bnbNode{level: 1, lowerBound: 1}
	items := [4]*bnbNode{n1, n2, n3, n4}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := MkPriorityQueue()
	for _, value := range items {
		pq.Push(value)
	}

	previousBnbNode := &bnbNode{level: 1, lowerBound: 0}

	// Take the items out, they arrive in an increasing order by lower bound.
	for !pq.Empty() {
		item := pq.Pop()
		assert.True(t, item.lowerBound > previousBnbNode.lowerBound, "Smaller lower bound should be prioritized.")
		previousBnbNode = item
	}
}

func TestPopVal(t *testing.T) {
	q1 := priorityQueue{}
	q2 := priorityQueue{}
	item1 := &pqItem{
		value: &bnbNode{level: 1},
	}
	item2 := item1
	q1.Push(item1)
	q2.Push(item2)
	assert.Equal(t, q1.Pop().(*pqItem).value, q2.PopVal())
}

func TestPushVal(t *testing.T) {
	pq := priorityQueue{}
	item := &bnbNode{}

	pq.PushVal(item)
	assert.Equal(t, item, pq.PopVal())
}
