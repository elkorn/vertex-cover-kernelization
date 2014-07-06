package graph

type coverageMap map[Vertex]bool

// TODO copy the graph instead of mutating
func removeOnce(g *Graph, removed coverageMap) func(Vertex) {
	return func(v Vertex) {
		if !removed[v] {
			err := g.RemoveVertex(v)
			if nil != err {
				removed[v] = true
			}
		}
	}
}
