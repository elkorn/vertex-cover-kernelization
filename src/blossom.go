package graph

import (
	"errors"
	"fmt"

	"github.com/deckarep/golang-set"
)

type blossom struct {
	Root     Vertex
	edges    mapset.Set
	vertices mapset.Set
}

func MkBlossom(root Vertex, completor *Edge, edges ...*Edge) *blossom {
	result := &blossom{
		Root:     root,
		edges:    mapset.NewSet(),
		vertices: mapset.NewSet(),
	}

	add := func(edge *Edge) {
		result.edges.Add(edge)
		result.vertices.Add(edge.from)
		result.vertices.Add(edge.to)
	}

	for _, edge := range edges {
		add(edge)
	}

	add(completor)

	return result
}

func (self *blossom) Contract(g *Graph, matching mapset.Set) {
	g.ForAllNeighbors(self.Root, func(edge *Edge, idx int, done chan<- bool) {
		neighbor := getOtherVertex(self.Root, edge)
		if !self.vertices.Contains(neighbor) {
			return
		}

		g.ForAllNeighbors(neighbor, func(edge *Edge, idx int, done chan<- bool) {
			distantNeighbor := getOtherVertex(neighbor, edge)
			if distantNeighbor == self.Root {
				return
			}

			g.rewireEdge(edge, neighbor, self.Root)
			if nil != matching && matching.Contains(edge) {
				matching.Remove(edge)
			}
		})

		g.RemoveVertex(neighbor)
	})
}

func (self *blossom) Expand(target Vertex, matching mapset.Set, g *Graph) ([]*Edge, Vertex) {
	// the side of B, ( u’ → ... → w’ ),
	// going from u’ to w’ are chosen to ensure that the new path is still
	// alternating (u’ is exposed with respect to M ∩ B, \{ w', w \} ∈ E ⧵ M).
	// TODO: What about 'u’ is exposed with respect to M ∩ B'?
	// Is this guaranteed by the Edmonds logic?
	bGraph := MkGraph(g.currentVertexIndex)
	for e := range self.edges.Iter() {
		edge := e.(*Edge)
		bGraph.AddEdge(edge.from, edge.to)
	}

	var exitVertex Vertex

	g.ForAllNeighbors(target, func(edge *Edge, index int, done chan<- bool) {
		// { w', w } ∈ E ⧵ M
		exit := getOtherVertex(target, edge)
		Debug("Checking %v-%v, matched: %v, in blossom: %v", edge.from, edge.to, matching.Contains(edge), bGraph.hasVertex(exit))
		if exit != self.Root && !matching.Contains(edge) && bGraph.hasVertex(exit) {
			Debug("Found exit %v", exit)
			exitVertex = exit
			done <- true
		}
	})

	return ShortestPathInGraph(bGraph, self.Root, exitVertex), exitVertex
}

func (self *Graph) setEdgeAtCoords(from, to int, value *Edge) {
	self.neighbors[from][to] = value
	self.neighbors[to][from] = value
}

func (self *Edge) changeEndpoint(which, newEndpoint Vertex) {
	if self.from == which {
		self.from = newEndpoint
	} else if self.to == which {
		self.to = newEndpoint
	}
}

func (self *Graph) rewireEdge(edge *Edge, from, newAnchor Vertex) {
	to := getOtherVertex(from, edge)

	if newAnchor == to {
		panic(errors.New(fmt.Sprintf("Cannot rewire edge %v-%v to %v-%v", from, to, newAnchor, to)))
	}

	fi := from.toInt()
	nAi := newAnchor.toInt()
	ti := to.toInt()

	edge.changeEndpoint(from, newAnchor)

	self.setEdgeAtCoords(fi, ti, nil)
	self.setEdgeAtCoords(nAi, ti, edge)

	self.degrees[fi]--
	self.degrees[nAi]++
}
