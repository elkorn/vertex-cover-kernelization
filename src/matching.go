package graph

import "github.com/deckarep/golang-set"

// A matching, M , of G is a subset of the edges E, such that no vertex
// in V is incident to more that one edge in M .
// Intuitively we can say that no two edges in M have a common vertex.

// A matching M is said to be maximal if M is not properly contained in
// any other matching.
// Formally, M !⊂ M' for any matching M' of G.
// Intuitively, this is equivalent to saying that a matching is maximal if we cannot
// add any edge to the existing set
func maximalMatching(g *Graph) (matching mapset.Set, outsiders mapset.Set) {
	matching = mapset.NewSet()
	outsiders = mapset.NewSet()
	added := make([]bool, len(g.Vertices))
	g.ForAllEdges(func(edge *Edge, index int, done chan<- bool) {
		if !(added[edge.from.toInt()] || added[edge.to.toInt()]) {
			matching.Add(edge)
			added[edge.from.toInt()] = true
			added[edge.to.toInt()] = true
		} else {
			outsiders.Add(edge)
		}
	})

	return matching, outsiders
}

// Given G = (V, E) and a matching M of G, a vertex v is exposed,
// if no edge of M is incident with v.
func (self Vertex) isExposed(matching mapset.Set) bool {
	for element := range matching.Iter() {
		edge := element.(Edge)
		if edge.IsCoveredBy(self) {
			return false
		}
	}

	return true
}

// An augmenting path P is an alternating path that starts and ends
// at two distinct exposed vertices.
func isAugmentingPath(path []int, matching mapset.Set) bool {
	start := MkVertex(path[0])
	end := MkVertex(path[len(path)-1])

	return start.isExposed(matching) &&
		end.isExposed(matching) &&
		isAlternatingPathWithMatching(path, matching)
}

// A matching augmentation along an augmenting path P
// is the operation of replacing M with a new matching M1 = M⊕P = (M⧵P)∪(P⧵M).
func matchingAugmentation(path []int, M mapset.Set) mapset.Set {
	return M.SymmetricDifference(pathToSet(path))
}

func pathToSet(path []int) (P mapset.Set) {
	P = mapset.NewSet()
	addedFirst := false
	forAllCoordPairsInPath(path, func(prevFrom, prevTo, curFrom, curTo int, done chan<- bool) {
		if !addedFirst {
			P.Add(MkEdgeValFromInts(prevFrom, prevTo))
			addedFirst = true
		}

		P.Add(MkEdgeValFromInts(curFrom, curTo))
	})

	return P
}

func findAugmentingPath(G *Graph, M mapset.Set) (result []*Edge) {
	blossoms := make([]*blossom, G.currentVertexIndex)
	// TODO what should the capacity be?
	// B02 F ← empty forest
	F := MkForest(G.currentVertexIndex)
	// B03 unmark all vertices and edges in G, mark all edges of M
	marker := mkEdmondsMarker(G)
	matchedEdges := make([]*Edge, G.currentVertexIndex)
	for e := range M.Iter() {
		edge := e.(*Edge)
		marker.SetEdgeMarked(edge, true)
		// There cannot exist two edges beginning in the same vertex in
		// a matching.
		matchedEdges[edge.from.toInt()] = edge
	}

	// B05 for each exposed vertex v do
	G.ForAllVertices(func(vertex Vertex, index int, done chan<- bool) {
		if vertex.isExposed(M) {
			// B06 create a singleton tree { v } and add the tree to F
			F.AddTree(MkTree(vertex, G.currentVertexIndex))
		}
	})

	// B08 while there is an unmarked vertex v in F with distance( v, root( v ) ) even do
	F.ForAllVertices(func(v Vertex, done chan<- bool) {
		if nil != result {
			done <- true
			return
		}

		if marker.IsVertexMarked(v) {
			return
		}

		var distance int
		if distance = F.Distance(v, F.Root(v)); distance%2 != 0 {
			return
		}

		// B09 while there exists an unmarked edge e = { v, w } do
		for marker.ExistsUnmarkedEdgeFromVertex(v) {
			G.ForAllNeighbors(v, func(e *Edge, index int, done chan<- bool) {
				if marker.IsEdgeMarked(e) {
					return
				}

				// B10            if w is not in F then
				w := getOtherVertex(v, e)
				if !F.HasVertex(w) {
					// w is matched, so add e and w's matched edge to F
					// B11 x ← vertex matched to w in M
					matchedEdge := matchedEdges[w.toInt()]
					// x := getOtherVertex(w, matchedEdge)
					// B12 add edges { v, w } and { w, x } to the tree of v
					F.AddEdge(v, e)           // { v, w }
					F.AddEdge(v, matchedEdge) // { w, x }
				} else { // B13 else
					// B14 if distance( w, root( w ) ) is odd then
					if F.Distance(w, F.Root(w))%2 == 1 {
						// Do nothing.
					} else { // B15 else
						// B16 if root( v ) ≠ root( w ) then
						if vRoot, wRoot := F.Root(v), F.Root(w); vRoot != wRoot {
							// Report an augmenting path in F \cup { e }.
							// B17  P ← path ( root( v ) → ... → v ) → ( w → ... → root( w ) )
							result = F.Path(MkTreePath(vRoot, v), MkTreePath(w, wRoot))
							// B18 return P
							done <- true
							return
						} else {
							// Contract a blossom in G and look for the path in the contracted graph.
							// B20 B ← blossom formed by e and edges on the path v → w in T
							blossomRoot := F.lookup(vRoot).CommonAncestor(v, w)
							blossomPath := F.Path(MkTreePath(v, w))
							B := MkBlossom(blossomRoot, e, blossomPath...)
							// B21 G’, M’ ← contract G and M by B
							// TODO: @refactor possible performance issue
							// Rewiring should be undoable so there is no need to copy the whole graph each time.
							gPrime := G.Copy()
							mPrime := M.Clone()
							B.Contract(gPrime, mPrime)
							blossoms[B.Root.toInt()] = B
							// B22 P’ ← find_augmenting_path( G’, M’ )
							pPrime := findAugmentingPath(gPrime, mPrime)
							// B23 P ← lift P’ to G
							// B24 return P
							result = lift(pPrime, M, blossoms, G)
							done <- true
							return
						}
						// B25 end if
					}
					// B26 end if
				}
				// B27 end if
				// B28 mark edge e
				marker.SetEdgeMarked(e, true)
			})
		}
		// B29 end while
		// B30 mark vertex v
		marker.SetVertexMarked(v, true)
	}) // B31 end while

	if nil == result {
		// B32 return empty path
		result = make([]*Edge, 0, G.NEdges())
	}

	return result
}

func lift(path []*Edge, matching mapset.Set, blossoms []*blossom, g *Graph) (result []*Edge) {
	// If the path contains contracted blossoms, then the size of the result size
	// must be enlarged for each blossom by (n-1)/2, where n is a blossom's size.
	// if P’ traverses through a segment u → vB → w in G’,
	// then this segment is replaced with the segment u → ( u’ → ... → w’ ) → w in G,
	// where blossom vertices u’ and w’ and the side of B, ( u’ → ... → w’ ),
	// going from u’ to w’ are chosen to ensure that the new path is still
	// alternating (u’ is exposed with respect to M ∩ B, \{ w', w \} ∈ E ⧵ M).

	processedBlossoms := make([]bool, len(blossoms))
	isProcessable := func(b *blossom) bool {
		return nil != b && !processedBlossoms[b.Root.toInt()]
	}
	result = make([]*Edge, 0, cap(blossoms))
	exits := make([]Vertex, g.currentVertexIndex)
	getLiftedExitEdge := func(edge *Edge) *Edge {
		fi := edge.from.toInt()
		ti := edge.to.toInt()
		if exits[fi] != 0 {
			return g.getEdgeByCoordinates(exits[fi].toInt(), ti)
		}

		if exits[ti] != 0 {
			return g.getEdgeByCoordinates(fi, exits[ti].toInt())
		}

		// Or just return edge?
		return g.getEdgeByCoordinates(fi, ti)
	}

	for i, n := 0, len(path); i < n; i++ {
		curEdge := path[i]
		fi := curEdge.from.toInt()
		ti := curEdge.to.toInt()
		b := blossoms[fi]
		Debug("Processing %v", curEdge)
		var expansion []*Edge
		if nil == b {
			Debug("Adding %v", curEdge)
			result = append(result, getLiftedExitEdge(curEdge))
			if b = blossoms[ti]; isProcessable(b) {
				w := getOtherVertex(curEdge.to, path[i+1])
				expansion, exits[ti] = b.Expand(w, matching, g)
				Debug("Blossom after edge, adding %v", expansion)
				result = append(result, expansion...)
				processedBlossoms[ti] = true
			}
		} else if isProcessable(b) {
			expansion, exits[fi] = b.Expand(curEdge.to, matching, g)
			Debug("Blossom before edge, adding %v", expansion)
			result = append(result, expansion...)
			processedBlossoms[fi] = true
			Debug("Adding %v", curEdge)
			result = append(result, getLiftedExitEdge(curEdge))
		} else {
			Debug("No blossoms to process, adding %v", *curEdge)
			result = append(result, getLiftedExitEdge(curEdge))
		}
	}

	return result
}
