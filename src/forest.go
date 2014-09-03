package graph

import (
	"errors"
	"fmt"
)

// TODO: refactor to allow adding only vertices and edges through the public interface.

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
	if t := self.lookup(v); t != nil {
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

func (self *forest) Distance(a, b Vertex) int {
	// Return the length of the path from a to b in this forest.
	return self.lookup(a).Distance(a, b)
}

func (self *forest) ForAllVertices(fn func(Vertex, chan<- bool)) {
	done := make(chan bool, 1)

	for vi, tree := range self.vertexTreeLookup {
		if nil == tree {
			continue
		}

		fn(MkVertex(vi), done)

		select {
		case <-done:
			return
		default:
		}
	}
}

func (self *forest) HasVertex(v Vertex) bool {
	return self.lookup(v) != nil
}

// TODO: Create a method AddEdgeFromPtr which reuses provided Edge.
func (self *forest) AddEdge(root Vertex, edge *Edge) {
	Debug("Adding edge %v-%v to tree %v", edge.from, edge.to, root)
	tree := self.lookup(root)
	addIfNotRoot := func(v Vertex) {
		existingRoot := self.Root(v)
		if v != root && existingRoot == INVALID_VERTEX {
			self.addVertexToLookup(v, tree)
		}
	}

	addIfNotRoot(edge.from)
	addIfNotRoot(edge.to)
	tree.AddEdge(edge.from, edge.to)
}

func (self *forest) addVertexToLookup(vertex Vertex, t *tree) {
	// This is supposed to enforce the trees of a forest to be disjoint.
	Debug("Adding vertex %v to lookup %v", vertex, t.Root)
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

func (self *forest) lookup(v Vertex) *tree {
	return self.vertexTreeLookup[v.toInt()]
}

func (self *forest) checkRoot(v Vertex) {
	if INVALID_VERTEX == self.Root(v) {
		panic(errors.New(fmt.Sprintf("Vertex %v does not exist in the forest", v)))
	}
}

func (self *forest) Path(treePathEndpoints ...*treePath) (result []*Edge) {
	// Returning an array of edges (instead of ints - vertex indices)
	// due to the fact that in the graph, coordinates (a,b) and (b,a)
	// point to the same instance of an edge.
	paths := make([][]*Edge, len(treePathEndpoints))
	totalLength := 0
	for i, endpoint := range treePathEndpoints {
		self.checkRoot(endpoint.to)
		tree := self.lookup(endpoint.from)
		if nil == tree {
			panic(errors.New(fmt.Sprintf("Vertex %v does not exist in the forest", endpoint.from)))
		}

		Debug("Searching for path %v-%v in %v", endpoint.from, endpoint.to, tree.Root)
		paths[i] = tree.Path(endpoint.from, endpoint.to)
		totalLength += len(paths[i])
	}

	result = make([]*Edge, 0, totalLength)
	for _, path := range paths {
		result = append(result, path...)
	}

	return result
}
