package graph

import "github.com/deckarep/golang-set"

type blossom struct {
	Root  Vertex
	edges Edges
}

// func getEndpoints(edges Edges)

// func (g *Graph) contractBlossom(b blossom) {
// 	contractionMap := make(NeighborMap, b.Root.toInt()+1)

// }

// TODO edge contraction within a graph has to be refactored to use rewiring.
// Thanks to that approach, blossom lifting will be possible.

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

func findAugmentingPath(G *Graph, M mapset.Set) []int {
	// TODO what should the capacity be?
	// B02    F ← empty forest
	F := MkForest(100)
	// B03    unmark all vertices and edges in G, mark all edges of M
	markedVertex := make([]bool, G.currentVertexIndex)
	edgeMarkMatrix := mkBoolMatrix(G.currentVertexIndex, G.currentVertexIndex)
	for edge := range M.Iter() {
		setEdgeMarked(edgeMarkMatrix, edge, true)
	}
}

func setEdgeMarked(mx [][]bool, edge *Edge, state bool) {
	a, b := edge.from.toInt(), edge.to.toInt()
	mx[a][b] = state
	mx[b][a] = state
}
