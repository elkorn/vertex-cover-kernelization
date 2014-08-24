package graph

import "errors"

func getOtherVertex(v Vertex, edge *Edge) Vertex {
	if edge.from != v {
		return edge.from
	}

	if edge.to != v {
		return edge.to
	}

	panic(errors.New("An edge with the same vertex as both endpoints may not exist."))
}