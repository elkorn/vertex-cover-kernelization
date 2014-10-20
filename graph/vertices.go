package graph

import "github.com/deckarep/golang-set"

type Vertices []Vertex

func (self Vertices) toSet() (result mapset.Set) {
	result = mapset.NewSet()
	for _, vertex := range self {
		result.Add(vertex)
	}

	return
}
