package graph

type IntStack struct {
	s *Stack
}

func (self *IntStack) Push(value int) {
	self.s.Push(value)
}

func (self *IntStack) Pop() int {
	if self.Empty() {
		panic("Trying to pop from an empty stack.")
	}

	return self.s.Pop().(int)
}

// // TODO add an Iter method.
// It should return a buffered channel of length self.count
func (self *IntStack) Values() []int {
	result := make([]int, self.s.count)
	for i, original := range self.s.Values() {
		result[i] = original.(int)
	}

	return result
}

func (self *IntStack) Empty() bool {
	return self.s.Empty()
}

func MkIntStack(capacity int) *IntStack {
	return &IntStack{
		s: MkStack(capacity),
	}
}
