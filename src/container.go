package graph

type Container interface {
	Push(value int)
	Pop() int
	Values() []int
	Empty() bool
}
