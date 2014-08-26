package graph

type tree struct {
	Root Vertex
	g    *Graph
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
	// TODO should the existence of `a` be required?
	self.addVertex(a)
	self.addVertex(b)
	self.g.AddEdge(a, b)
}

func (self *tree) addVertex(vertex Vertex) {
	self.g.isVertexDeleted[vertex.toInt()] = false
}

func (self *tree) Distance(a, b Vertex) (distance int) {
	distance = 0
	var edge *Edge
	from := a.toInt()
	bi := b.toInt()
	to := -1
	// This is safe in the context of a tree, where cycles should not exist.
	for to != bi {
		Debug("Iterating...")
		for to = 0; to < self.g.currentVertexIndex; to++ {
			if edge = self.g.getEdgeByCoordinates(from, to); edge != nil {
				Debug("%v -> %v: edge", from, to)
				distance++
				from = to
				break
			}

			Debug("%v -> %v: no edge", from, to)
		}
	}

	return distance
}
