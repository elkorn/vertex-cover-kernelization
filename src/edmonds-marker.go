package graph

type edmondsMarker struct {
	g                      *Graph
	markedVertex           []bool
	numberOfMarkedVertices int
	edgeMarkLookup         [][]bool
	markedEdgesFromVertex  []int
}

func mkEdmondsMarker(G *Graph) *edmondsMarker {
	return &edmondsMarker{
		markedVertex:           make([]bool, G.currentVertexIndex),
		numberOfMarkedVertices: 0,
		edgeMarkLookup:         mkBoolMatrix(G.currentVertexIndex, G.currentVertexIndex),
		markedEdgesFromVertex:  make([]int, G.currentVertexIndex),
		g: G,
	}
}

func (self *edmondsMarker) SetEdgeMarked(edge *Edge, state bool) {
	a, b := edge.GetIntEndpoints()
	if self.edgeMarkLookup[a][b] == state || self.edgeMarkLookup[b][a] == state {
		return
	}

	self.edgeMarkLookup[a][b] = state
	self.edgeMarkLookup[b][a] = state
	var incr int
	if state {
		incr = 1
	} else {
		incr = -1
	}

	self.markedEdgesFromVertex[a] += incr
	self.markedEdgesFromVertex[b] += incr
}

func (self *edmondsMarker) IsEdgeMarked(edge *Edge) bool {
	a, b := edge.GetIntEndpoints()
	return self.edgeMarkLookup[a][b] || self.edgeMarkLookup[b][a]
}

func (self *edmondsMarker) IsVertexMarked(v Vertex) bool {
	return self.markedVertex[v.toInt()]
}

func (self *edmondsMarker) ExistsUnmarkedEdgeFromVertex(v Vertex) bool {
	// IDEA: length of neighbors collection for each vertex should be
	// maintained in graph and used for the edmondsMarker.
	degree, err := self.g.Degree(v)
	if err != nil {
		panic(err)
	}

	return self.markedEdgesFromVertex[v.toInt()] < degree
}

func (self *edmondsMarker) SetVertexMarked(v Vertex, state bool) {
	self.markedVertex[v.toInt()] = state
}
