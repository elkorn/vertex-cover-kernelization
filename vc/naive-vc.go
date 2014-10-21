package vc

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type naiveVC struct {
	result bool
	solved bool
	cover  mapset.Set
}

func (self *naiveVC) solveIter(g *graph.Graph, k int) bool {
	if g.NEdges() == 0 {
		return false
	}

	if k == 0 {
		return true
	}

	edge := self.getFirstExistingEdge(g)
	g1, g2 := g.Copy(), g.Copy()
	g1.RemoveVertex(edge.From)
	g2.RemoveVertex(edge.To)
	result1, result2 := self.solveIter(g1, k-1), self.solveIter(g2, k-1)

	if result1 {
		self.cover.Add(edge.From)
	} else if result2 {
		self.cover.Add(edge.To)
	}

	return result1 || result2
}

func (self *naiveVC) getFirstExistingEdge(g *graph.Graph) (result *graph.Edge) {
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		result = edge
		done <- true
	})

	return
}

func (self *naiveVC) solve(g *graph.Graph, k int) bool {
	if self.solved {
		return self.result
	}

	self.solved = true
	self.result = self.solveIter(g, k)
	return self.result
}

func NaiveVC(G *graph.Graph, k int) (bool, mapset.Set) {
	instance := &naiveVC{
		cover: mapset.NewSet(),
	}

	instance.solve(G, k)
	return instance.result, instance.cover
}
