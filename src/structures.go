package graph

import "github.com/deckarep/golang-set"

// Structures are related to fold-struction.go.

type structure struct {
	S mapset.Set
	q int
}

func MkStructure(q int, s ...Vertex) *structure {
	result := &structure{
		q: q,
		S: mapset.NewSet(),
	}

	for _, s := range s {
		result.S.Add(s)
	}

	return result
}

func MkGoodPair(s ...Vertex) *structure {
	return MkStructure(-1, s...)
}

type goodPairInfo struct {
	pair                                *structure
	numNeighborhoodAlmostDominatedPairs int
	numNeighborhoodEdges                int
}

func mkGoodPairInfo(goodPair *structure) *goodPairInfo {
	result := &goodPairInfo{
		pair: goodPair,
	}

	return result
}

func (self *goodPairInfo) countAlmostDominatedPairs(g *Graph) int {
	result := 0
	var u Vertex
	for v := range self.pair.S.Iter() {
		u = v.(Vertex)
		break
	}

	g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
		x := getOtherVertex(u, edge)
		g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
			y := getOtherVertex(u, edge)
			if x == y {
				return
			}

			if x.almostDominates(y, g) {
				result++
			}
		})
	})

	self.numNeighborhoodAlmostDominatedPairs = result
	return result
}

func (self *goodPairInfo) countNeighborhoodEdges(g *Graph) int {
	result := 0
	var u Vertex
	for v := range self.pair.S.Iter() {
		u = v.(Vertex)
		break
	}

	g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
		x := getOtherVertex(u, edge)
		g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
			y := getOtherVertex(u, edge)
			if x == y {
				return
			}

			if g.hasEdge(x, y) {
				result++
			}

		})
	})

	self.numNeighborhoodEdges = result
	return result
}

func identifyStructures(G *Graph, k int) *StructurePriorityQueueProxy {
	forAllGoodPairInfos := func(set mapset.Set, action func(*goodPairInfo)) {
		for gpi := range set.Iter() {
			action(gpi.(*goodPairInfo))
		}
	}

	tags := computeTags(G)
	result := MkStructurePriorityQueue()
	// TODO: capacity is arbitrary.
	// TODO: Create a goodPair struct that has only two vertices.
	possibleGoodPairs := mapset.NewSet()
	// The first vertex in a good pair is found as follows:
	// 1. tag(u) is lex. max over tag(w) for all w of the same degree as u.
	G.ForAllVertices(func(u Vertex, done chan<- bool) {
		deg := G.Degree(u)
		tagU := tags[u.toInt()]
		foundValidU := true
		G.forAllVerticesOfDegree(deg, func(w Vertex) {
			comparison := tagU.Compare(tags[w.toInt()], G)
			if comparison == -1 {
				done <- true
				foundValidU = false
			}
		})

		if foundValidU {
			possibleGoodPairs.Add(mkGoodPairInfo(MkGoodPair(u)))
		}
	})

	// 2. If the graph is reguler, the number of pairs {x,y} \subseteq N(u) s.t.
	// y is almost-dominated by x is maximized.
	if G.IsRegular() {
		maxAlmostDominated := 0
		forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
			almostDominated := possibleGoodPair.countAlmostDominatedPairs(G)
			if almostDominated > maxAlmostDominated {
				maxAlmostDominated = almostDominated
			}
		})

		forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
			if possibleGoodPair.numNeighborhoodAlmostDominatedPairs != maxAlmostDominated {
				possibleGoodPairs.Remove(possibleGoodPair)
			}
		})
	}

	// 3. The number of edges in the subgraph induced by N(u) is maximized.
	maxNumEdges := 0
	forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
		if possibleGoodPair.countNeighborhoodEdges(G) > maxNumEdges {
			maxNumEdges = possibleGoodPair.numNeighborhoodEdges
		}
	})

	forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
		if possibleGoodPair.numNeighborhoodEdges < maxNumEdges {
			possibleGoodPairs.Remove(possibleGoodPair)
		}
	})

	// The second vertex is chosen as follows.
	possibleZ := mapset.NewSet()
	forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
		// a) If there exist 2 neighbors of u: w,v s.t. v is almost-dominated
		// by w, then z is almost dominated by a neighbor of u.
		// TODO: maintain an array of almost-dominated pairs of neighbors.
		if possibleGoodPair.numNeighborhoodAlmostDominatedPairs > 0 {
			// G.ForAllNeighbors(possibleGoodPair
			// b) the degree of z is max among N(u) satisfying a).

			// c) z is adjacent to the least number of N(u) satisfying a) and b)

			// d) The number of edges in a subgraph induced by N(u) is maximized.
			// TODO: This sounds a bit fishy. Double-check this with the paper.
		}
	})

	return result
}
