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
	u, z                                Vertex
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

func (self *goodPairInfo) U() Vertex {
	if INVALID_VERTEX == self.u {
		iter := self.pair.S.Iter()
		self.u = (<-iter).(Vertex)
	}

	return self.u
}

func (self *goodPairInfo) Z() Vertex {
	if INVALID_VERTEX == self.z {
		iter := self.pair.S.Iter()
		<-iter
		self.z = (<-iter).(Vertex)
	}

	return self.z
}

type degree struct {
	v   Vertex
	val int
}

func forAllGoodPairInfos(set mapset.Set, action func(*goodPairInfo)) {
	for gpi := range set.Iter() {
		action(gpi.(*goodPairInfo))
	}
}

func (self *Graph) forAllVerticesOfDegreeGeq(degree int, action func(Vertex)) {
	self.ForAllVertices(func(v Vertex, done chan<- bool) {
		if self.Degree(v) >= degree {
			action(v)
		}
	})
}

func identifyGoodVertices(G *Graph) mapset.Set {
	result := mapset.NewSet()
	G.forAllVerticesOfDegreeGeq(7, func(v Vertex) {
		result.Add(MkStructure(-1, v))
	})

	return result
}

func identifyGoodPairs(G *Graph) mapset.Set {
	tags := computeTags(G)
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

	// Having chosen the first vertex u in a good pair, to choose the second
	// vertex, we pick a neighbor z of u such that the following conditions are
	// satisfied in their respective order.
	possibleZ := mapset.NewSet()
	forAllGoodPairInfos(possibleGoodPairs, func(possibleGoodPair *goodPairInfo) {
		// a) If there exist 2 neighbors of u: w,v s.t. v is almost-dominated
		// by w, then z is almost dominated by a neighbor of u.

		if possibleGoodPair.numNeighborhoodAlmostDominatedPairs > 0 {
			// TODO: This step can be done earlier to avoid additional loops,
			// improving average complexity.
			G.ForAllNeighbors(possibleGoodPair.U(), func(edge *Edge, done chan<- bool) {
				n := getOtherVertex(possibleGoodPair.U(), edge)
				G.ForAllNeighbors(possibleGoodPair.U(), func(edge *Edge, done chan<- bool) {
					z := getOtherVertex(possibleGoodPair.U(), edge)
					if n == z {
						return
					}

					if n.almostDominates(z, G) {
						possibleZ.Add(z)
					}
				})
			})
			// b) the degree of z is max among N(u) satisfying a).
			maxDegreeOfZ := 0
			for zInter := range possibleZ.Iter() {
				z := zInter.(Vertex)
				if deg := G.Degree(z); deg > maxDegreeOfZ {
					maxDegreeOfZ = deg
				}
			}

			for zInter := range possibleZ.Iter() {
				z := zInter.(Vertex)
				if G.Degree(z) != maxDegreeOfZ {
					possibleZ.Remove(zInter)
				}
			}

			// c) z is adjacent to the least number of N(u) satisfying a) and b)
			minAdjacency := MAX_INT
			// TODO: This should be a priority queue.
			adjacencies := mapset.NewSet()
			for zInter := range possibleZ.Iter() {
				z := zInter.(Vertex)
				adjacency := 0
				G.ForAllNeighbors(
					possibleGoodPair.U(),
					func(edge *Edge, done chan<- bool) {
						if G.hasEdge(z, getOtherVertex(possibleGoodPair.U(), edge)) {
							adjacency++
						}
					})

				if adjacency < minAdjacency {
					minAdjacency = adjacency
				}

				adjacencies.Add(&degree{
					v:   z,
					val: adjacency,
				})
			}

			for adj := range adjacencies.Iter() {
				adjacency := adj.(*degree)
				if adjacency.val != minAdjacency {
					possibleZ.Remove(adjacency.v)
				}
			}

			// d) The number of shared neighbors between z and a neighbor of u is
			// maximized among N(u) satisfying a), b) and c).
			maxSharedNeighbors := 0
			sharedNeighbors := mapset.NewSet()
			for zInter := range possibleZ.Iter() {
				z := zInter.(Vertex)
				curSharedNeighbors := 0
				G.ForAllNeighbors(z, func(edge *Edge, done chan<- bool) {
					G.ForAllNeighbors(
						possibleGoodPair.U(),
						func(edge *Edge, done chan<- bool) {
							sharedNeighbor := getOtherVertex(
								possibleGoodPair.U(),
								edge)
							if G.hasEdge(z, sharedNeighbor) {
								curSharedNeighbors++
							}
						})
				})

				if curSharedNeighbors > maxSharedNeighbors {
					maxSharedNeighbors = curSharedNeighbors
				}

				sharedNeighbors.Add(&degree{
					v:   z,
					val: curSharedNeighbors,
				})
			}

			for shnInter := range sharedNeighbors.Iter() {
				shn := shnInter.(*degree)
				if shn.val == maxSharedNeighbors {
					// This is the z we're looking for.
					possibleGoodPair.pair.S.Add(shn.v)
					break
				} else {
					possibleZ.Remove(shn.v)
				}
			}

			// Now, there should be one vertex in possibleZ.
			// If that's not the case, it means that multiple vertices fulfill
			// the criteria a,b,c,d.
			// Is that even possible?
			// If so, additional good pairs of (u, z_1), (u, z_2)... should
			// probably be created.
		}
	})

	return possibleGoodPairs
}

func identifyStructures(G *Graph, k int) *StructurePriorityQueueProxy {
	result := MkStructurePriorityQueue()
	// goodVertices := identifyGoodVertices(G)
	// goodPairs := identifyGoodPairs(G)

	// TODO: Compute the priority of the good pairs.
	// TODO: Put the good pairs in the queue.

	return result
}
