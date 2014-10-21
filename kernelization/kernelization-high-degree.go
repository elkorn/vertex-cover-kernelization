package kernelization

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
)

func removeVerticesWithDegreeGreaterThan(self *graph.Graph, k int) (graph.Neighbors, int) {
	toRemove := mapset.NewThreadUnsafeSet()
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
