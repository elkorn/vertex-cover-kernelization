package kernelization

import "github.com/deckarep/golang-set"

func (self *graph.Graph) removeVerticesWithDegreeGreaterThan(k int) (graph.Neighbors, int) {
	toRemove := mapset.NewSet()
	removed := 0

	self.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		if self.Degree(v) > k {
			removed++
			toRemove.Add(v)
		}
	})

	result := make(graph.Neighbors, 0, removed)
	for vInter := range toRemove.Iter() {
		vertex := vInter.(graph.Vertex)
		result = append(result, vertex)
		self.RemoveVertex(vertex)
	}

	return result, removed
}
