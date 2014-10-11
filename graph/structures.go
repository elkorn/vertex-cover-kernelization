package graph

import "github.com/deckarep/golang-set"

// Structures are related to fold-struction.go.

type structure struct {
	S        mapset.Set
	Elements Vertices
	q        int
}

func MkStructure(q int, s ...Vertex) *structure {
	result := &structure{
		q:        q,
		S:        mapset.NewSet(),
		Elements: s,
	}

	for _, s := range s {
		result.S.Add(s)
	}

	return result
}

func MkStructureWithSet(q int, S mapset.Set, s ...Vertex) *structure {
	return &structure{
		q:        q,
		S:        S,
		Elements: s,
	}
}

func MkStructureWithCapacity(q, capacity int, s ...Vertex) *structure {
	result := &structure{
		q:        q,
		S:        mapset.NewSet(),
		Elements: make(Vertices, len(s), capacity),
	}

	copy(result.Elements, s)

	for _, s := range s {
		result.S.Add(s)
	}

	return result
}

func mkGoodPairStruct(s ...Vertex) *structure {
	return MkStructureWithCapacity(-1, 2, s...)
}

// In Part 6. of the proof states that when having a good pair where
// d(u)=3 and N(u) are all deg. 5 verts that do not share any common neighbors
// besides u, 2 vertices in N(u) may be connected (it seems to be a valid case),
// and then CS applies.
// SOLUTION:
// A vertex cannot be its own neighbor. It means that when having
// e.g. a set {v,w,z} where if v,w are adj., then if z is not adj. to v or w,
// the structure remains valid (v is a neighbor of w and w is a neighbor of v,
// neither of which are shared with z).
// An if statement has to be written for that.
func (self *structure) neighborsOfUShareCommonVertexOtherThanU(u, z Vertex, g *Graph) (neighborsShareCommonVertexOtherThanU, neighborsAreDisjoint bool) {
	neighborsAreDisjoint = true
	g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
		if neighborsShareCommonVertexOtherThanU {
			done <- true
			return
		}

		v1 := getOtherVertex(u, e)

		g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
			if neighborsShareCommonVertexOtherThanU {
				done <- true
				return
			}

			v2 := getOtherVertex(u, e)
			if v1 == v2 {
				return
			}

			if g.HasEdge(v1, v2) {
				neighborsAreDisjoint = false
				Debug("N(%v) are not disjoint, %v-%v exists", u, v1, v2)
				neighborsShareCommonVertexOtherThanU = true
				Debug("N(%v) share common vertices %v, %v", u, v1, v2)
				done <- true
				return
			}

			g.ForAllNeighbors(v1, func(e *Edge, done chan<- bool) {
				if neighborsShareCommonVertexOtherThanU {
					done <- true
					return
				}

				n1 := getOtherVertex(v1, e)

				g.ForAllNeighbors(v2, func(e *Edge, done chan<- bool) {
					Debug("Checking %v|%v", v1, v2)
					n2 := getOtherVertex(v2, e)
					if n1 == n2 && n1 != u {
						neighborsShareCommonVertexOtherThanU = true
						Debug("N(%v) share common vertex %v", u, n1)
						done <- true
					}
				})
			})
		})
	})

	if !neighborsShareCommonVertexOtherThanU {
		Debug("N(%v) do not share other common vertices", u)
	}

	if neighborsAreDisjoint {
		Debug("All N(%v) are disjoint", u)
	}

	return
}

func (self *structure) countDegree5Neighbors(u Vertex, g *Graph) (degree5NeighborsCount int, hasOnlyDegree5Neighbors bool) {
	hasOnlyDegree5Neighbors = true
	g.ForAllNeighbors(u, func(e *Edge, done chan<- bool) {
		w := getOtherVertex(u, e)
		deg := g.Degree(w)
		if deg == 5 {
			degree5NeighborsCount++
		} else {
			Debug("N(%v): %v is of deg. %v", u, w, deg)
			hasOnlyDegree5Neighbors = false
		}
	})

	Debug("There are %v N(%v) of deg. 5", degree5NeighborsCount, u)
	if hasOnlyDegree5Neighbors {
		Debug("All N(%v) are of deg. 5", u)
	}
	return
}

type goodPair struct {
	pair                                *structure
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
	u := self.U()

	Debug("Counting almost-dominated pairs for u: %v", u)
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
	Debug("%v almost dominated pairs in N(%v)", result, u)
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

			// Debug("Looking for edge %v-%v", x, y)
			if g.HasEdge(x, y) {
				// Debug("Found edge %v-%v", x, y)
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
	return self.pair.Elements[0]
}

func (self *goodPair) Z() Vertex {
	return self.pair.Elements[1]
}

func (self *goodPair) setZ(z Vertex) {
	n := len(self.pair.Elements)
	if n != 1 && n != 2 {
		panic("Invalid good pair!")
	}

	if n == 1 {
		self.pair.Elements = append(self.pair.Elements, z)
	} else if n == 2 {
		self.pair.Elements[1] = z
	}

	self.pair.S.Add(z)
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

func trimSet(source mapset.Set, toRemove *mapset.Set) mapset.Set {
	result := source.Difference(*toRemove)
	(*toRemove).Clear()
	return result
}

func identifyGoodPairs(G *Graph) mapset.Set {
	tags := computeTags(G)
	possibleGoodPairs := mapset.NewSet()
	invalidPairs := mapset.NewSet()
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

			if tagU.Compare(tags[w.toInt()], G) == -1 {
				foundValidU = false
			}
		})

		if foundValidU {
			Debug("1) satisfied, adding possible pair with u: %v (deg. %v)", u, deg)
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

		possibleGoodPairs = trimSet(possibleGoodPairs, &toRemove)
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

	possibleGoodPairs = trimSet(possibleGoodPairs, &toRemove)

	// Having chosen the first vertex u in a good pair, to choose the second
	// vertex, we pick a neighbor z of u such that the following conditions are
	// satisfied in their respective order.
	// TODO: Verify whether this is feasible.
	additionalPairs := mapset.NewSet()
	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		Debug("\n")
		Debug("Looking for Z for %v...", possibleGoodPair.U())
		var possibleZ mapset.Set
		// a) If there exist 2 neighbors of u: w,v s.t. v is almost-dominated
		// by w, then z is almost dominated by a neighbor of u.
		u := possibleGoodPair.U()
		if possibleGoodPair.countAlmostDominatedPairs(G) > 0 {
			possibleZ = mapset.NewSet()
			G.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
				n := getOtherVertex(u, edge)
				G.ForAllNeighbors(u, func(edge *Edge, done chan<- bool) {
					z := getOtherVertex(u, edge)
					if n == z {
						return
					}

					if z.isAlmostDominatedBy(n, G) {
						Debug("a) satisfied, adding %v", z)
						possibleZ.Add(z)
					}
				})
			})
		} else {
			// If no vertex in N ( u ) is almost-dominated by another vertex in
			// N ( u ) , then (a) is vacuously satisfied by every vertex in N ( u ),
			// and z will be a neighbor of u of maximum degree.
			_, possibleZ = G.getNeighborsWithSet(u)
		}

		Debug("Satisfying a): %v", possibleZ)
		// b) the degree of z is max among N(u) satisfying a).
		maxDegreeOfZ := 0
		for zInter := range possibleZ.Iter() {
			z := zInter.(Vertex)
			if deg := G.Degree(z); deg > maxDegreeOfZ {
				maxDegreeOfZ = deg
			}
		}

		Debug("Max. degree of z: %v", maxDegreeOfZ)
		for zInter := range possibleZ.Iter() {
			z := zInter.(Vertex)
			if G.Degree(z) != maxDegreeOfZ {
				toRemove.Add(z)
			}
		}

		possibleZ = trimSet(possibleZ, &toRemove)
		Debug("Satisfying a),b): %v", possibleZ)

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
					if G.HasEdge(z, getOtherVertex(possibleGoodPair.U(), edge)) {
						adjacency++
					}
				})

			if adjacency < minAdjacency {
				minAdjacency = adjacency
			}

			adj := &degree{
				v:   z,
				val: adjacency,
			}

			Debug("Adjacency of %v: %v", adj.v, adj.val)
			adjacencies.Add(adj)
		}

		Debug("Min. adjacency of z: %v", minAdjacency)

		for adj := range adjacencies.Iter() {
			adjacency := adj.(*degree)
			if adjacency.val != minAdjacency {
				toRemove.Add(adjacency.v)
			}
		}

		possibleZ = trimSet(possibleZ, &toRemove)
		Debug("Satisfying a),b),c): %v", possibleZ)
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
						if G.HasEdge(z, sharedNeighbor) {
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
		Debug("Shn: %v, maxShn: %v", sharedNeighbors, maxSharedNeighbors)
		for shnInter := range sharedNeighbors.Iter() {
			shn := shnInter.(*degree)
			if shn.val == maxSharedNeighbors {
				// This is the z we're looking for.
				if possibleGoodPair.IsValid() {
					// There is a pair with that U already.
					// Create another one.
					additionalPairs.Add(mkGoodPair(possibleGoodPair.U(), shn.v))
				} else {
					possibleGoodPair.setZ(shn.v)
				}
			} else {
				toRemove.Add(shn.v)
			}
		}

		possibleZ = trimSet(possibleZ, &toRemove)

		if !possibleGoodPair.IsValid() {
			invalidPairs.Add(possibleGoodPair)
		}
	})

	forAllGoodPairs(possibleGoodPairs, func(pgp *goodPair) {
		Debug("U: %v, pairs: %v, edges: %v", pgp.U(), pgp.numNeighborhoodAlmostDominatedPairs, pgp.numNeighborhoodEdges)
	})

	Debug("additional pairs: %v", additionalPairs)
	return possibleGoodPairs.Union(additionalPairs).Difference(invalidPairs)
}

func identifyStructures(G *Graph, k int) *StructurePriorityQueueProxy {
	result := MkStructurePriorityQueue()
	goodVertices := identifyGoodVertices(G)
	goodPairs := identifyGoodPairs(G)

	for gvInter := range goodVertices.Iter() {
		result.Push(gvInter.(*structure), G)
	}

	for gpInter := range goodPairs.Iter() {
		result.Push((gpInter.(*goodPair)).pair, G)
	}

	return result
}
