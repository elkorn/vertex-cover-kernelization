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

//    INPUT:  Graph G, initial matching M on G
//    OUTPUT: maximum matching M* on G
// A1 function find_maximum_matching( G, M ) : M*
// A2     P ← find_augmenting_path( G, M )
// A3     if P is non-empty then
// A4          return find_maximum_matching( G, augment M along P )
// A5     else
// A6          return M
// A7     end if
// A8 end function

//     INPUT:  Graph G, matching M on G
//     OUTPUT: augmenting path P in G or empty path if none found
// B01 function find_augmenting_path( G, M ) : P
// B02    F ← empty forest
// B03    unmark all vertices and edges in G, mark all edges of M
// B05    for each exposed vertex v do
// B06        create a singleton tree { v } and add the tree to F
// B07    end for
// B08    while there is an unmarked vertex v in F with distance( v, root( v ) ) even do
// B09        while there exists an unmarked edge e = { v, w } do
// B10            if w is not in F then
//                    // w is matched, so add e and w's matched edge to F
// B11                x ← vertex matched to w in M
// B12                add edges { v, w } and { w, x } to the tree of v
// B13            else
// B14                if distance( w, root( w ) ) is odd then
//                        // Do nothing.
// B15                else
// B16                    if root( v ) ≠ root( w ) then
//                            // Report an augmenting path in F \cup { e }.
// B17                        P ← path ( root( v ) → ... → v ) → ( w → ... → root( w ) )
// B18                        return P
// B19                    else
//                            // Contract a blossom in G and look for the path in the contracted graph.
// B20                        B ← blossom formed by e and edges on the path v → w in T
// B21                        G’, M’ ← contract G and M by B
// B22                        P’ ← find_augmenting_path( G’, M’ )
// B23                        P ← lift P’ to G
// B24                        return P
// B25                    end if
// B26                end if
// B27            end if
// B28            mark edge e
// B29        end while
// B30        mark vertex v
// B31    end while
// B32    return empty path
// B33 end function

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
							gPrime := G.Copy()
							mPrime := M.Clone()
							B.Contract(gPrime, mPrime)
							blossoms[B.Root.toInt()] = B
							// B22 P’ ← find_augmenting_path( G’, M’ )
							pPrime := findAugmentingPath(gPrime, mPrime)
							// B23 P ← lift P’ to G
							result = lift(pPrime, M, blossoms, G)
							// B24 return P
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
	})
	// B31 end while

	if nil == result {
		// B32 return empty path
		result = make([]*Edge, 0, G.NEdges())
	}

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

func (self *blossom) Expand(target Vertex, matching mapset.Set, g *Graph) []*Edge {
	// the side of B, ( u’ → ... → w’ ),
	// going from u’ to w’ are chosen to ensure that the new path is still
	// alternating (u’ is exposed with respect to M ∩ B, \{ w', w \} ∈ E ⧵ M).
	// TODO: What about 'u’ is exposed with respect to M ∩ B' ?
	bGraph := MkGraph(g.currentVertexIndex)
	for e := range self.edges.Iter() {
		edge := e.(*Edge)
		bGraph.AddEdge(edge.from, edge.to)
	}

	gv := MkGraphVisualizer()
	gv.Display(bGraph)

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

	return ShortestPathInGraph(bGraph, self.Root, exitVertex)
}

func lift(path []*Edge, matching mapset.Set, blossoms []*blossom, g *Graph) (result []*Edge) {
	// If the path contains contracted blossoms, then the size of the result size
	// must be enlarged for each blossom by (n-1)/2, where n is a blossom's size.
	// if P’ traverses through a segment u → vB → w in G’,
	// then this segment is replaced with the segment u → ( u’ → ... → w’ ) → w in G,
	// where blossom vertices u’ and w’ and the side of B, ( u’ → ... → w’ ),
	// going from u’ to w’ are chosen to ensure that the new path is still
	// alternating (u’ is exposed with respect to M ∩ B, \{ w', w \} ∈ E ⧵ M).
	// TODO: @refactor add a 'checkHasBlossom' function.
	processedBlossoms := make([]bool, len(blossoms))
	result = make([]*Edge, 0, cap(blossoms))
	for i, n := 0, len(path); i < n; i++ {
		curEdge := path[i]
		fi := curEdge.from.toInt()
		ti := curEdge.to.toInt()
		if nil == blossoms[fi] {
			result = append(result, curEdge)
			if b := blossoms[ti]; nil != b {
			} else {
				// u := ti
				w := getOtherVertex(curEdge.to, path[i+1])
				result = append(result, b.Expand(w, matching, g)...)
				processedBlossoms[ti] = true
			}
		}
	}

	return result
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
