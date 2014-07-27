package graph

type Queue struct {
	nodes []int
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Push(n int) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]int, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Empty() bool {
	return q.count == 0
}

func (q *Queue) Pop() int {
	if q.Empty() {
		panic("Trying to pop from an empty queue.")
	}

	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}
