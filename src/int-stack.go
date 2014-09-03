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

// TODO: add an Iter method.
// It should return a buffered channel of length self.count
func (self *IntStack) Iter() <-chan int {
	iter := make(chan int, self.s.count)
	count := self.s.count - 1

	// This is cheating a bit- the algorithms are supposed to be single-threaded.
	// TODO: Implement a proper iterator?
	go func() {
		for ; count >= 0; count-- {
			iter <- self.s.values[count].(int)
		}
		close(iter)
	}()

	return iter
}

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
