package vc

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type naiveVC struct {
	result        bool
	solved        bool
	cover         mapset.Set
	kernelization func(*graph.Graph, int) int
}

func (self *naiveVC) solveIter(g *graph.Graph, k int, cover mapset.Set) (bool, mapset.Set) {
	if g.NEdges() == 0 {
		return true, cover
	}

	if k == 0 {
		return false, cover
	}

	edge := self.getFirstExistingEdge(g)
	g1, g2 := g.Copy(), g.Copy()
	if self.kernelization != nil {
		reduction := self.kernelization(g1, k)
		if reduction == -1 {
			return false, cover
		}

		self.kernelization(g2, k)
		k -= reduction
	}

	g1.RemoveVertex(edge.From)
	g2.RemoveVertex(edge.To)
	result1, cov1 := self.solveIter(g1, k-1, cover /*.Clone()*/)
	result2, cov2 := self.solveIter(g2, k-1, cover /*.Clone()*/)
	// var resultCover mapset.Set
	if result1 {
		cover = cov1
		// cover.Add(edge.From)
	} else if result2 {
		cover = cov2
		// cover.Add(edge.To)
	}

	return result1 || result2, cover
}

func (self *naiveVC) getFirstExistingEdge(g *graph.Graph) (result *graph.Edge) {
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		result = edge
		done <- true
	})

	return
}

func (self *naiveVC) solve(g *graph.Graph, k int) (bool, mapset.Set) {
	if self.solved {
		return self.result, self.cover
	}

	self.solved = true
	self.result, self.cover = self.solveIter(g, k, self.cover)
	return self.result, self.cover
}

func NaiveVC(G *graph.Graph, k int) (bool, mapset.Set) {
	instance := &naiveVC{
		cover: mapset.NewThreadUnsafeSet(),
	}
	return instance.solve(G, k)
}

func KernelizedNaiveVC(G *graph.Graph, k int, kernelization func(*graph.Graph, int) int) (bool, mapset.Set) {
	instance := &naiveVC{
		cover:         mapset.NewThreadUnsafeSet(),
		kernelization: kernelization,
	}

	return instance.solve(G, k)
}
