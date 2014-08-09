package graph

type Stack struct {
	values []int
	count  int
}

func (self *Stack) Push(value int) {
	if self.count >= len(self.values) {
		values := make([]int, len(self.values)*2)
		copy(values, self.values)
		self.values = values
	}

	self.values[self.count] = value
	self.count++
}

func (self *Stack) Pop() int {
	if self.Empty() {
		panic("Trying to pop from an empty stack.")
	}

	value := self.values[self.count-1]
	self.count--
	return value
}

func (self *Stack) Empty() bool {
	return self.count == 0
}

func MkStack() *Stack {
	return &Stack{
		values: []int{},
		count:  0,
	}
}
