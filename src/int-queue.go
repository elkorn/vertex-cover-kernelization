package graph

type IntQueue struct {
	q *Queue
}

func MkIntQueue(size int) *IntQueue {
	return &IntQueue{
		q: MkQueue(size),
	}
}

func (self *IntQueue) Push(n int) {
	self.q.Push(n)
}

func (self *IntQueue) Pop() int {
	return self.q.Pop().(int)
}

func (self *IntQueue) Empty() bool {
	return self.q.Empty()
}
