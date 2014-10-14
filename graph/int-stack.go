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

func (self *IntStack) Iter() <-chan int {
	iter := make(chan int, self.Size())
	count := self.Size() - 1
	go func() {
		for ; count >= 0; count-- {
			iter <- self.s.values[count].(int)
		}
		close(iter)
	}()

	return iter
}

func (self *IntStack) Size() int {
	return self.s.Size()
}

func (self *IntStack) Values() []int {
	result := make([]int, 0, self.Size())
	for val := range self.Iter() {
		result = append(result, val)
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
