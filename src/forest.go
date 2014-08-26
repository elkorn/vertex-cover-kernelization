package graph

// TODO Add a Path method, which will be used in Distance.
import (
	"errors"
	"fmt"
)

// The capacity of the forest should be the same as the capacity of each of its
// trees, to accomodate the graph performance optimizations.
type forest struct {
	trees            []*tree
	vertexTreeLookup []*tree
}

func MkForest(capacity int) *forest {
	return &forest{
		trees:            make([]*tree, capacity),
		vertexTreeLookup: make([]*tree, capacity),
	}
}

func (self *forest) Root(v Vertex) Vertex {
	if t := self.vertexTreeLookup[v.toInt()]; t != nil {
		return t.Root
	}

	return INVALID_VERTEX
}

func (self *forest) AddTree(t *tree) {
	if n := len(self.trees); t.g.currentVertexIndex != n {
		panic(
			errors.New(
				fmt.Sprintf(
					"Only trees with capacity of %v can be added to the forest of capacity %v",
					t.g.currentVertexIndex,
					n)))
	}

	self.trees[t.Root.toInt()] = t
	t.g.ForAllVertices(func(vertex Vertex, index int, done chan<- bool) {
		self.addVertexToLookup(vertex, t)
	})
}

func (self *forest) addVertexToLookup(vertex Vertex, t *tree) {
	if existing := self.Root(vertex); existing != INVALID_VERTEX {
		panic(
			errors.New(
				fmt.Sprintf(
					"Vertex %v cannot be added under root %v - "+
						"it already exists under another root (%v).",
					vertex,
					t.Root,
					existing)))
	}

	self.vertexTreeLookup[vertex.toInt()] = t
}

func (self *forest) Distance(a, b Vertex) {
	// Return the length of the path from a to b in this forest.
}
