package graph

type Neighbors []Vertex

func (self *Graph) getNeighbors(v Vertex) Neighbors {
	result := Neighbors{}
	for i := range self.getCoveredEdgePositions(v) {
		edge := self.Edges[i]
		if edge.from != v && !contains(result, edge.from) {
			result = append(result, edge.from)
		} else if edge.to != v && !contains(result, edge.to) {
			result = append(result, edge.to)
		}
	}

	return result
}
