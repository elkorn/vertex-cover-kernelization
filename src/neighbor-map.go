package graph

type NeighborMap map[Vertex]Neighbors

func (self NeighborMap) AddNeighborOfVertex(v, n Vertex) {
	if self[v] == nil {
		self[v] = Neighbors{n}
	} else {
		if !contains(self[v], n) {
			self[v] = append(self[v], n)
		}
	}
}
