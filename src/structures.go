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

func mkGoodPairStruct(s ...Vertex) *structure {
	return MkStructure(-1, s...)
}

type goodPair struct {
	pair                                *structure
	u, z                                Vertex
	numNeighborhoodAlmostDominatedPairs int
	numNeighborhoodEdges                int
}

func mkGoodPair(s ...Vertex) *goodPair {
	result := &goodPair{
		pair: mkGoodPairStruct(s...),
	}

	return result
}

func (self *goodPair) countAlmostDominatedPairs(g *Graph) int {
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

func (self *goodPair) countNeighborhoodEdges(g *Graph) int {
	result := 0
	u := self.U()

	g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
		x := getOtherVertex(u, edge)
		g.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
			y := getOtherVertex(u, edge)
			if x == y {
				return
			}

			Debug("Looking for edge %v-%v", x, y)
			if g.hasEdge(x, y) {
				Debug("Found edge %v-%v", x, y)
				result++
			}

		})
	})

	self.numNeighborhoodEdges = result / 2
	// This is less computationally expensive than maintaining a set of
	// processed edges and within the neighborhood it's safe - each edge is
	// counted twice due to the graph being undirected.
	return self.numNeighborhoodEdges
}

func (self *goodPair) U() Vertex {
	if INVALID_VERTEX == self.u {
		iter := self.pair.S.Iter()
		self.u = (<-iter).(Vertex)
	}

	return self.u
}

func (self *goodPair) Z() Vertex {
	if INVALID_VERTEX == self.z {
		iter := self.pair.S.Iter()
		<-iter
		self.z = (<-iter).(Vertex)
	}

	return self.z
}

func (self *goodPair) IsValid() bool {
	// The pair must have u and z to be valid.
	return self.pair.S.Cardinality() == 2
}

type degree struct {
	v   Vertex
	val int
}

func forAllGoodPairs(set mapset.Set, action func(*goodPair)) {
	for gpi := range set.Iter() {
		action(gpi.(*goodPair))
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
	possibleGoodPairs := mapset.NewSet()
	invalidPairs := mapset.NewSet()
	updatePossibleGoodPairs := func(toRemove mapset.Set) {
		possibleGoodPairs = possibleGoodPairs.Difference(toRemove)
		toRemove.Clear()
	}
	// The first vertex in a good pair is found as follows:
	// 1. tag(u) is lex. max over tag(w) for all w of the same degree as u.
	Debug("Looking for U...")
	G.ForAllVertices(func(u Vertex, done chan<- bool) {
		deg := G.Degree(u)
		tagU := tags[u.toInt()]
		Debug("Tag of %v: %v", u, tagU.neighbors)
		foundValidU := true
		G.forAllVerticesOfDegree(deg, func(w Vertex) {
			if !foundValidU {
				return
			}

			comparison := tagU.Compare(tags[w.toInt()], G)
			if comparison == -1 {
				foundValidU = false
			}
		})

		if foundValidU {
			Debug("1) satisfied, adding possible pair with u: %v", u)
			possibleGoodPairs.Add(mkGoodPair(u))
		}
	})

	// 2. If the graph is regular, the number of pairs {x,y} \subseteq N(u) s.t.
	// y is almost-dominated by x is maximized.
	toRemove := mapset.NewSet()
	if G.IsRegular() {
		maxAlmostDominated := 0
		forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
			almostDominated := possibleGoodPair.countAlmostDominatedPairs(G)
			if almostDominated > maxAlmostDominated {
				maxAlmostDominated = almostDominated
			}
		})

		forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
			if possibleGoodPair.numNeighborhoodAlmostDominatedPairs != maxAlmostDominated {
				Debug("2) not satisfied, removing pair with u: %v", possibleGoodPair.U())
				toRemove.Add(possibleGoodPair)
			}
		})

		updatePossibleGoodPairs(toRemove)
	}

	// 3. The number of edges in the subgraph induced by N(u) is maximized.
	maxNumEdges := 0
	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		if possibleGoodPair.countNeighborhoodEdges(G) > maxNumEdges {
			maxNumEdges = possibleGoodPair.numNeighborhoodEdges
		}

		Debug(
			"Num. neighborhood edges for %v: %v",
			possibleGoodPair.U(),
			possibleGoodPair.numNeighborhoodEdges)
	})

	Debug("Max. neighborhood edges: %v", maxNumEdges)

	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		if possibleGoodPair.numNeighborhoodEdges < maxNumEdges {
			Debug("3) not satisfied, removing pair with u: %v", possibleGoodPair.U())
			toRemove.Add(possibleGoodPair)
		}
	})

	updatePossibleGoodPairs(toRemove)
	// Having chosen the first vertex u in a good pair, to choose the second
	// vertex, we pick a neighbor z of u such that the following conditions are
	// satisfied in their respective order.
	possibleZ := mapset.NewSet()
	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		// a) If there exist 2 neighbors of u: w,v s.t. v is almost-dominated
		// by w, then z is almost dominated by a neighbor of u.
		if possibleGoodPair.numNeighborhoodAlmostDominatedPairs > 0 {
			u := possibleGoodPair.U()
			G.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
				n := getOtherVertex(u, edge)
				G.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
					z := getOtherVertex(u, edge)
					if n == z {
						return
					}

					if n.almostDominates(z, G) {
						Debug("a) is satisfied, adding %v as possible z for u %v", z, u)
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

			// Now, there should be one vertex in possibleZ.
			// If that's not the case, it means that multiple vertices fulfill
			// the criteria a,b,c,d.
			// Is that even possible?
			// If so, additional good pairs of (u, z_1), (u, z_2)... should
			// probably be created.
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
		}

		if !possibleGoodPair.IsValid() {
			invalidPairs.Add(possibleGoodPair)
		}
	})

	return possibleGoodPairs.Difference(invalidPairs)
}

func identifyStructures(G *Graph, k int) *StructurePriorityQueueProxy {
	result := MkStructurePriorityQueue()
	// goodVertices := identifyGoodVertices(G)
	// goodPairs := identifyGoodPairs(G)

	// TODO: Compute the priority of the good pairs.
	// TODO: Put the good pairs in the queue.

	return result
}
