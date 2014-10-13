package graph

import "github.com/deckarep/golang-set"

type ntReductionYehuda struct {
	V0, C0 mapset.Set
}

func mkNtReductionYehuda(g *Graph, k int) *ntReductionYehuda {
	border := g.currentVertexIndex
	V0, C0 := mapset.NewSet(), mapset.NewSet()
	edges, _ := fordFulkerson(mkNetworkFlow(g))
	CB := mapset.NewSet()
	for _, edge := range edges {
		if CB.Contains(edge.from) {
			CB.Add(edge.to)
		} else {
			CB.Add(edge.from)
		}
	}

	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		if CB.Contains(v) == CB.Contains(v+Vertex(border)) {
			if CB.Contains(v) {
				C0.Add(v)
			}
		} else {
			V0.Add(v)
		}
	})

	return &ntReductionYehuda{
		V0: V0,
		C0: C0,
	}
}
