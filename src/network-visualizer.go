package graph

type networkVisualizer struct {
	gv *graphVisualizer
}

func MkNetworkVisualizer() *networkVisualizer {
	return &networkVisualizer{
		MkGraphVisualizer(),
	}
}

func convertToGraph(net Net) *Graph {
	result := mkGraphWithVertices(len(net))
	for x := range net {
		for y, arc := range net[x] {
			if nil == arc {
				continue
			}

			result.AddEdge(Vertex(y+1), Vertex(x+1))
		}
	}

	return result
}
