package graph

type blossom struct {
	Root  Vertex
	edges Edges
}

// func getEndpoints(edges Edges)

// func (g *Graph) contractBlossom(b blossom) {
// 	contractionMap := make(NeighborMap, b.Root.toInt()+1)

// }

// TODO edge contraction within a graph has to be refactored to use rewiring.
// Thanks to that approach, blossom lifting will be possible.
