package kernelization

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type ntReductionYehuda struct {
	V0, C0 mapset.Set
}

func mkNtReductionYehuda(g *graph.Graph, k int) *ntReductionYehuda {
	border := g.CurrentVertexIndex
	V0, C0 := mapset.NewSet(), mapset.NewSet()
	edges, _ := fordFulkerson(mkNetworkFlow(g))
	CB := mapset.NewSet()
	for _, edge := range edges {
		if CB.Contains(edge.From) {
			CB.Add(edge.To)
		} else {
			CB.Add(edge.From)
		}
	}

	g.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		if CB.Contains(v) == CB.Contains(v+graph.Vertex(border)) {
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
