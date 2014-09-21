package graph

type coverageMap map[Vertex]bool

// TODO: copy the graph instead of mutating
func removeOnce(g *Graph, removed coverageMap) func(Vertex) bool {
	return func(v Vertex) bool {
		if !removed[v] {
			err := g.RemoveVertex(v)
			if nil != err {
				removed[v] = true
				return true
			}
		}

		return false
	}
}
