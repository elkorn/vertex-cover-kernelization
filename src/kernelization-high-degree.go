package graph

import "github.com/deckarep/golang-set"

func (self *Graph) removeVerticesWithDegreeGreaterThan(k int) (Neighbors, int) {
	toRemove := mapset.NewSet()
	removed := 0

	self.ForAllVertices(func(v Vertex, done chan<- bool) {
		if self.Degree(v) > k {
			removed++
			toRemove.Add(v)
		}
	})

	result := make(Neighbors, 0, removed)
	for vInter := range toRemove.Iter() {
		vertex := vInter.(Vertex)
		result = append(result, vertex)
		self.RemoveVertex(vertex)
	}

	return result, removed
}
