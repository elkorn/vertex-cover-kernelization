package graph

type Stack struct {
	values []interface{}
	count  int
}

func (self *Stack) Push(value interface{}) {
	if self.count >= len(self.values) {
		values := make([]interface{}, len(self.values)*2)
		copy(values, self.values)
		self.values = values
	}

	self.values[self.count] = value
	self.count++
}

func (self *Stack) Pop() interface{} {
	if self.Empty() {
		panic("Trying to pop from an empty stack.")
	}

	value := self.values[self.count-1]
	self.count--
	return value
}

// TODO add an Iter method.
// It should return a buffered channel of length self.count
func (self *Stack) Values() []interface{} {
	tmp := MkStack(self.count)
	tmp.values = self.values
	tmp.count = self.count
	result := make([]interface{}, self.count)
	i := 0
	for !tmp.Empty() {
		result[i] = tmp.Pop()
		i++
	}

	return result
}

func (self *Stack) Empty() bool {
	return self.count == 0
}

func MkStack(capacity int) *Stack {
	return &Stack{
		values: make([]interface{}, capacity),
		count:  0,
	}
}
