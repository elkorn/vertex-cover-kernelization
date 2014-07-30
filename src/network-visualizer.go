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
	result := mkGraphWithVertices(len(*net))
	for y := range *net {
		for x, arc := range (*net)[y] {
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
