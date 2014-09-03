package graph

import "github.com/deckarep/golang-set"

// TODO: Refactor Path to return a queue.

type tree struct {
	Root Vertex
	g    *Graph
}

type treePath struct {
	from, to Vertex
}

func MkTreePath(from, to Vertex) *treePath {
	// This swap is intentional - searching for paths in the tree is not as
	// unintuitive with 'reverse' arguments order as searching for paths in a
	// forest would be.
	// TODO: see if it is possible to change for tree path lookup to use the
	// 'straightforward' argument order as well.
	return &treePath{
		from: from,
		to:   to,
	}
}

func MkTree(root Vertex, capacity int) (result *tree) {
	result = &tree{
		Root: root,
		g:    MkGraph(capacity),
	}

	result.g.ForAllVertices(func(vertex Vertex, index int, done chan<- bool) {
		result.g.isVertexDeleted[vertex.toInt()] = vertex != root
	})

	return result
}

func (self *tree) AddEdge(a, b Vertex) {
	// TODO: should the existence of `a` be required?
	self.addVertex(a)
	self.addVertex(b)
	self.g.AddEdge(a, b)
}

func (self *tree) Path(a, b Vertex) []*Edge {
	// // If the main graph itself is a tree AND it happens to be wholly included
	// // in this tree, then this tree needs to have the same amount of edges.
	// result = make([]*Edge, 0, self.g.NEdges())
	// self.forAllEdgesInPath(a, b, func(edge *Edge, done chan<- bool) {
	// 	Debug("Adding edge %v-%v to path", edge.from, edge.to)
	// 	result = append(result, edge)
	// })

	return ShortestPathInGraph(self.g, a, b)
}

func (self *tree) Distance(a, b Vertex) int {
	return ShortestDistanceInGraph(self.g, a, b)
}

func (self *tree) CommonAncestor(a, b Vertex) (ancestor Vertex) {
	existsPath, edgeToPath, _ := shortestPathInGraph(self.g, a, b)
	if !existsPath {
		return INVALID_VERTEX
	}

	existsRef, edgeToRef, _ := shortestPathInGraph(self.g, b, self.Root)
	if !existsRef {
		return INVALID_VERTEX
	}

	coordsInPath := mapset.NewSet()

	forEachCoordsInPath(a, b, edgeToPath, self.g, func(from, to int, done chan<- bool) {
		coordsInPath.Add(from)
		coordsInPath.Add(to)
	})

	ancestor = INVALID_VERTEX
	forEachCoordsInPath(b, self.Root, edgeToRef, self.g, func(from, to int, done chan<- bool) {
		if coordsInPath.Contains(from) {
			ancestor = MkVertex(from)
			done <- true
			return
		}

		if coordsInPath.Contains(to) {
			ancestor = MkVertex(to)
			done <- true
			return
		}
	})

	return ancestor
}

func (self *tree) addVertex(vertex Vertex) {
	self.g.isVertexDeleted[vertex.toInt()] = false
}
