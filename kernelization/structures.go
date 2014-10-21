package kernelization

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

// Structures are related to fold-struction.go.

type structure struct {
	S        mapset.Set
	Elements graph.Vertices
	q        int
}

func MkStructure(q int, s ...graph.Vertex) *structure {
	result := &structure{
		q:        q,
		S:        mapset.NewThreadUnsafeSet(),
		Elements: s,
	}

	for _, s := range s {
		result.S.Add(s)
	}

	return result
}

func MkStructureWithSet(q int, S mapset.Set, s ...graph.Vertex) *structure {
	return &structure{
		q:        q,
		S:        S,
		Elements: s,
	}
}

func MkStructureWithCapacity(q, capacity int, s ...graph.Vertex) *structure {
	result := &structure{
		q:        q,
		S:        mapset.NewThreadUnsafeSet(),
		Elements: make(graph.Vertices, len(s), capacity),
	}

	copy(result.Elements, s)

	for _, s := range s {
		result.S.Add(s)
	}

	return result
}

func mkGoodPairStruct(s ...graph.Vertex) *structure {
	return MkStructureWithCapacity(-1, 2, s...)
}

// In Part 6. of the proof states that when having a good pair where
// d(u)=3 and N(u) are all deg. 5 verts that do not share any common neighbors
// besides u, 2 graph.vertices in N(u) may be connected (it seems to be a valid case),
// and then CS applies.
// SOLUTION:
// A graph.vertex cannot be its own neighbor. It means that when having
// e.g. a set {v,w,z} where if v,w are adj., then if z is not adj. to v or w,
// the structure remains valid (v is a neighbor of w and w is a neighbor of v,
// neither of which are shared with z).
// An if statement has to be written for that.
func (self *structure) neighborsOfUShareCommonVertexOtherThanU(u, z graph.Vertex, g *graph.Graph) (neighborsShareCommonVertexOtherThanU, neighborsAreDisjoint bool) {
	neighborsAreDisjoint = true
	g.ForAllNeighbors(u, func(e *graph.Edge, done chan<- bool) {
		if neighborsShareCommonVertexOtherThanU {
			done <- true
			return
		}

		v1 := graph.GetOtherVertex(u, e)

		g.ForAllNeighbors(u, func(e *graph.Edge, done chan<- bool) {
			if neighborsShareCommonVertexOtherThanU {
				done <- true
				return
			}

			v2 := graph.GetOtherVertex(u, e)
			if v1 == v2 {
				return
			}

			if g.HasEdge(v1, v2) {
				neighborsAreDisjoint = false
				utility.Debug("N(%v) are not disjoint, %v-%v exists", u, v1, v2)
			}

			g.ForAllNeighbors(v1, func(e *graph.Edge, done chan<- bool) {
				if neighborsShareCommonVertexOtherThanU {
					done <- true
					return
				}

				n1 := graph.GetOtherVertex(v1, e)

				g.ForAllNeighbors(v2, func(e *graph.Edge, done chan<- bool) {
					utility.Debug("Checking %v|%v", v1, v2)
					n2 := graph.GetOtherVertex(v2, e)
					if n1 == n2 && n1 != u {
						neighborsShareCommonVertexOtherThanU = true
						utility.Debug("N(%v) share common graph.vertex %v", u, n1)
						done <- true
					}
				})
			})
		})
	})

	if !neighborsShareCommonVertexOtherThanU {
		utility.Debug("N(%v) do not share other common graph.vertices", u)
	}

	if neighborsAreDisjoint {
		utility.Debug("All N(%v) are disjoint", u)
	}

	return
}

func (self *structure) countDegree5Neighbors(u graph.Vertex, g *graph.Graph) (degree5NeighborsCount int, hasOnlyDegree5Neighbors bool) {
	hasOnlyDegree5Neighbors = true
	g.ForAllNeighbors(u, func(e *graph.Edge, done chan<- bool) {
		w := graph.GetOtherVertex(u, e)
		deg := g.Degree(w)
		if deg == 5 {
			degree5NeighborsCount++
		} else {
			utility.Debug("N(%v): %v is of deg. %v", u, w, deg)
			hasOnlyDegree5Neighbors = false
		}
	})

	utility.Debug("There are %v N(%v) of deg. 5", degree5NeighborsCount, u)
	if hasOnlyDegree5Neighbors {
		utility.Debug("All N(%v) are of deg. 5", u)
	}
	return
}

type goodPair struct {
	pair                                *structure
	numNeighborhoodAlmostDominatedPairs int
	numNeighborhoodEdges                int
}

func mkGoodPair(s ...graph.Vertex) *goodPair {
	result := &goodPair{
		pair: mkGoodPairStruct(s...),
	}

	return result
}

func (self *goodPair) countAlmostDominatedPairs(g *graph.Graph) int {
	result := 0
	u := self.U()

	utility.Debug("Counting almost-dominated pairs for u: %v", u)
	g.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
		x := graph.GetOtherVertex(u, edge)
		g.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
			y := graph.GetOtherVertex(u, edge)
			if x == y {
				return
			}

			if almostDominates(x, y, g) {
				result++
			}
		})
	})

	self.numNeighborhoodAlmostDominatedPairs = result
	utility.Debug("%v almost dominated pairs in N(%v)", result, u)
	return result
}

func (self *goodPair) countNeighborhoodEdges(g *graph.Graph) int {
	result := 0
	u := self.U()

	g.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
		x := graph.GetOtherVertex(u, edge)
		g.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
			y := graph.GetOtherVertex(u, edge)
			if x == y {
				return
			}

			// utility.Debug("Looking for edge %v-%v", x, y)
			if g.HasEdge(x, y) {
				// utility.Debug("Found edge %v-%v", x, y)
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

func (self *goodPair) U() graph.Vertex {
	return self.pair.Elements[0]
}

func (self *goodPair) Z() graph.Vertex {
	return self.pair.Elements[1]
}

func (self *goodPair) setZ(z graph.Vertex) {
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
	v   graph.Vertex
	val int
}

func forAllGoodPairs(set mapset.Set, action func(*goodPair)) {
	for gpi := range set.Iter() {
		action(gpi.(*goodPair))
	}
}

func forAllVerticesOfDegreeGeq(self *graph.Graph, degree int, action func(graph.Vertex)) {
	self.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		if self.Degree(v) >= degree {
			action(v)
		}
	})
}

func identifyGoodVertices(G *graph.Graph) mapset.Set {
	result := mapset.NewThreadUnsafeSet()
	forAllVerticesOfDegreeGeq(G, 7, func(v graph.Vertex) {
		result.Add(MkStructure(-1, v))
	})

	return result
}

func trimSet(source mapset.Set, toRemove *mapset.Set) mapset.Set {
	result := source.Difference(*toRemove)
	(*toRemove).Clear()
	return result
}

func identifyGoodPairs(G *graph.Graph) mapset.Set {
	tags := computeTags(G)
	possibleGoodPairs := mapset.NewThreadUnsafeSet()
	invalidPairs := mapset.NewThreadUnsafeSet()
	// The first graph.vertex in a good pair is found as follows:
	// 1. tag(u) is lex. max over tag(w) for all w of the same degree as u.
	utility.Debug("Looking for U...")
	G.ForAllVertices(func(u graph.Vertex, done chan<- bool) {
		deg := G.Degree(u)
		tagU := tags[u.ToInt()]
		utility.Debug("Tag of %v: %v", u, tagU.neighbors)
		foundValidU := true
		G.ForAllVerticesOfDegree(deg, func(w graph.Vertex) {
			if !foundValidU {
				return
			}

			if tagU.Compare(tags[w.ToInt()], G) == -1 {
				foundValidU = false
			}
		})

		if foundValidU {
			utility.Debug("1) satisfied, adding possible pair with u: %v (deg. %v)", u, deg)
			possibleGoodPairs.Add(mkGoodPair(u))
		}
	})

	// 2. If the graph is regular, the number of pairs {x,y} \subseteq N(u) s.t.
	// y is almost-dominated by x is maximized.
	toRemove := mapset.NewThreadUnsafeSet()
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
				utility.Debug("2) not satisfied, removing pair with u: %v", possibleGoodPair.U())
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

		utility.Debug(
			"Num. neighborhood edges for %v: %v",
			possibleGoodPair.U(),
			possibleGoodPair.numNeighborhoodEdges)
	})

	utility.Debug("Max. neighborhood edges: %v", maxNumEdges)

	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		if possibleGoodPair.numNeighborhoodEdges < maxNumEdges {
			utility.Debug("3) not satisfied, removing pair with u: %v", possibleGoodPair.U())
			toRemove.Add(possibleGoodPair)
		}
	})

	possibleGoodPairs = trimSet(possibleGoodPairs, &toRemove)

	// Having chosen the first graph.vertex u in a good pair, to choose the second
	// graph.vertex, we pick a neighbor z of u such that the following conditions are
	// satisfied in their respective order.
	// TODO: Verify whether this is feasible.
	additionalPairs := mapset.NewThreadUnsafeSet()
	forAllGoodPairs(possibleGoodPairs, func(possibleGoodPair *goodPair) {
		utility.Debug("\n")
		utility.Debug("Looking for Z for %v...", possibleGoodPair.U())
		var possibleZ mapset.Set
		// a) If there exist 2 neighbors of u: w,v s.t. v is almost-dominated
		// by w, then z is almost dominated by a neighbor of u.
		u := possibleGoodPair.U()
		if possibleGoodPair.countAlmostDominatedPairs(G) > 0 {
			possibleZ = mapset.NewThreadUnsafeSet()
			G.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
				n := graph.GetOtherVertex(u, edge)
				G.ForAllNeighbors(u, func(edge *graph.Edge, done chan<- bool) {
					z := graph.GetOtherVertex(u, edge)
					if n == z {
						return
					}

					if isAlmostDominatedBy(z, n, G) {
						utility.Debug("a) satisfied, adding %v", z)
						possibleZ.Add(z)
					}
				})
			})
		} else {
			// If no graph.vertex in N ( u ) is almost-dominated by another graph.vertex in
			// N ( u ) , then (a) is vacuously satisfied by every graph.vertex in N ( u ),
			// and z will be a neighbor of u of maximum degree.
			_, possibleZ = G.GetNeighborsWithSet(u)
		}

		utility.Debug("Satisfying a): %v", possibleZ)
		// b) the degree of z is max among N(u) satisfying a).
		maxDegreeOfZ := 0
		for zInter := range possibleZ.Iter() {
			z := zInter.(graph.Vertex)
			if deg := G.Degree(z); deg > maxDegreeOfZ {
				maxDegreeOfZ = deg
			}
		}

		utility.Debug("Max. degree of z: %v", maxDegreeOfZ)
		for zInter := range possibleZ.Iter() {
			z := zInter.(graph.Vertex)
			if G.Degree(z) != maxDegreeOfZ {
				toRemove.Add(z)
			}
		}

		possibleZ = trimSet(possibleZ, &toRemove)
		utility.Debug("Satisfying a),b): %v", possibleZ)

		// c) z is adjacent to the least number of N(u) satisfying a) and b)
		minAdjacency := utility.MAX_INT
		// TODO: This should be a priority queue.
		adjacencies := mapset.NewThreadUnsafeSet()
		for zInter := range possibleZ.Iter() {
			z := zInter.(graph.Vertex)
			adjacency := 0
			G.ForAllNeighbors(
				possibleGoodPair.U(),
				func(edge *graph.Edge, done chan<- bool) {
					if G.HasEdge(z, graph.GetOtherVertex(possibleGoodPair.U(), edge)) {
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

			utility.Debug("Adjacency of %v: %v", adj.v, adj.val)
			adjacencies.Add(adj)
		}

		utility.Debug("Min. adjacency of z: %v", minAdjacency)

		for adj := range adjacencies.Iter() {
			adjacency := adj.(*degree)
			if adjacency.val != minAdjacency {
				toRemove.Add(adjacency.v)
			}
		}

		possibleZ = trimSet(possibleZ, &toRemove)
		utility.Debug("Satisfying a),b),c): %v", possibleZ)
		// d) The number of shared neighbors between z and a neighbor of u is
		// maximized among N(u) satisfying a), b) and c).
		maxSharedNeighbors := 0
		sharedNeighbors := mapset.NewThreadUnsafeSet()
		for zInter := range possibleZ.Iter() {
			z := zInter.(graph.Vertex)
			curSharedNeighbors := 0
			G.ForAllNeighbors(z, func(edge *graph.Edge, done chan<- bool) {
				G.ForAllNeighbors(
					possibleGoodPair.U(),
					func(edge *graph.Edge, done chan<- bool) {
						sharedNeighbor := graph.GetOtherVertex(
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

		// Now, there should be one graph.vertex in possibleZ.
		// If that's not the case, it means that multiple graph.vertices fulfill
		// the criteria a,b,c,d.
		// Is that even possible?
		// If so, additional good pairs of (u, z_1), (u, z_2)... should
		// probably be created.
		utility.Debug("Shn: %v, maxShn: %v", sharedNeighbors, maxSharedNeighbors)
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
		utility.Debug("U: %v, pairs: %v, edges: %v", pgp.U(), pgp.numNeighborhoodAlmostDominatedPairs, pgp.numNeighborhoodEdges)
	})

	utility.Debug("additional pairs: %v", additionalPairs)
	return possibleGoodPairs.Union(additionalPairs).Difference(invalidPairs)
}

func identifyStructures(G *graph.Graph, k int) *StructurePriorityQueueProxy {
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
