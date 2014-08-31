package graph

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
	// TODO see if it is possible to change for tree path lookup to use the
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
	// TODO should the existence of `a` be required?
	self.addVertex(a)
	self.addVertex(b)
	self.g.AddEdge(a, b)
}

func (self *tree) Path(a, b Vertex) (result []*Edge) {
	// If the main graph itself is a tree AND it happens to be wholly included
	// in this tree, then this tree needs to have the same amount of edges.
	result = make([]*Edge, 0, self.g.NEdges())
	self.forAllEdgesInPath(a, b, func(edge *Edge, done chan<- bool) {
		Debug("Adding edge %v-%v to path", edge.from, edge.to)
		result = append(result, edge)
	})

	return result
}

func (self *tree) Distance(a, b Vertex) (distance int) {
	distance = 0

	self.forAllEdgesInPath(a, b, func(edge *Edge, done chan<- bool) {
		distance++
	})

	return distance
}

func (self *tree) addVertex(vertex Vertex) {
	self.g.isVertexDeleted[vertex.toInt()] = false
}

func (self *tree) forAllEdgesInPath(a, b Vertex, fn func(*Edge, chan<- bool)) {
	var edge *Edge
	from := a.toInt()
	bi := b.toInt()
	to := -1
	done := make(chan bool, 1)
	prevFrom := -1 // Used to avoid locking when using reverse path ordering.
	// This is safe in the context of a tree, where cycles should not exist.
	for to != bi {
		for to = 0; to < self.g.currentVertexIndex; to++ {
			if to == prevFrom {
				continue
			}

			if edge = self.g.getEdgeByCoordinates(from, to); edge != nil {
				Debug("%v -> %v: edge (%v)", MkVertex(from), MkVertex(to), edge)
				fn(edge, done)

				select {
				case <-done:
					return
				default:
				}

				prevFrom = from
				from = to
				break

			}

			Debug("%v -> %v: no edge", MkVertex(from), MkVertex(to))
		}
	}
}
