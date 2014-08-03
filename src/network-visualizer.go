package graph

type networkVisualizer struct {
	gv *graphVisualizer
}

func MkNetworkVisualizer() *networkVisualizer {
	return &networkVisualizer{
		MkGraphVisualizer(),
	}
}

func convertToGraph(net *Net) *Graph {
	// FIXME this will not work when the network is represented as a dense 2D array.
	arcs := (*net).arcs
	result := mkGraphWithVertices(len(arcs))
	for y := range arcs {
		for x, arc := range arcs[y] {
			if nil == arc {
				continue
			}

			result.AddEdge(Vertex(y+1), Vertex(x+1))
		}
	}

	return result
}

func (self *networkVisualizer) MkJpg(net *Net) error {
	return self.gv.MkJpg(convertToGraph(net), "net")
}

func (self *networkVisualizer) Display(net *Net) {
	self.gv.Display(convertToGraph(net))
}
