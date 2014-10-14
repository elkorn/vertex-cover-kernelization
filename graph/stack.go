package graph

type Stack struct {
	values []interface{}
	count  int
}

func (self *Stack) Push(value interface{}) {
	if self.count >= len(self.values) {
		self.values = append(self.values, value)
	} else {
		self.values[self.count] = value
	}

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

func (self *Stack) Size() int {
	return self.count
}

func (self *Stack) Iter() <-chan interface{} {
	iter := make(chan interface{}, self.count)
	count := self.count - 1

	go func() {
		for ; count >= 0; count-- {
			iter <- self.values[count]
		}
		close(iter)
	}()

	return iter
}

func (self *Stack) Values() []interface{} {
	result := make([]interface{}, 0, self.Size())
	for val := range self.Iter() {
		result = append(result, val)
	}

	return result
}

func (self *Stack) Empty() bool {
	return self.count == 0
}

func MkStack(capacity int) *Stack {
	return &Stack{
		values: make([]interface{}, 0, capacity),
		count:  0,
	}
}
