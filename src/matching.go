package graph

import (
	"container/list"
	"errors"

	"github.com/deckarep/golang-set"
)

type nodeInformation struct {
	Parent  Vertex
	Root    Vertex
	IsOuter bool
}

func mkNodeInformation(parent, root Vertex, isOuter bool) *nodeInformation {
	return &nodeInformation{
		Parent:  parent,
		Root:    root,
		IsOuter: isOuter,
	}
}

// A matching, M , of G is a subset of the edges E, such that no vertex
// in V is incident to more that one edge in M .
// Intuitively we can say that no two edges in M have a common vertex.

// A matching M is said to be maximal if M is not properly contained in
// any other matching.
// Formally, M !⊂ M' for any matching M' of G.
// Intuitively, this is equivalent to saying that a matching is maximal if we cannot
// add any edge to the existing set
func FindMaximalMatching(g *Graph) (matching mapset.Set, outsiders mapset.Set) {
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

func FindMaximumMatching(G *Graph) (result *Graph) {
	if G.NEdges() == 0 {
		return MkGraph(0)
	}

	result = MkGraph(G.currentVertexIndex)

	for {
		path := findAugmentingPath(G, result)
		Debug("Found aug. path: %v", path)
		if nil == path {
			return result // maximum is found.
		}

		updateMatching(path, result)
	}
}

func findAugmentingPath(G, M *Graph) (result *list.List) {
	forest := make([]*nodeInformation, G.currentVertexIndex)
	workList := MkQueue(G.NEdges())
	G.ForAllVertices(func(v Vertex, index int, done chan<- bool) {
		Debug("Checking vertex %v", v)
		// The forest must be initially seeded with singleton nodes only.
		if deg, _ := M.Degree(v); deg > 0 {
			Debug("Has %v matched edges.", deg)
			return
		}

		Debug("%v is exposed, adding to forest.", v)
		forest[v.toInt()] = mkNodeInformation(INVALID_VERTEX, v, true)

		G.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
			e := MkEdge(v, getOtherVertex(v, edge))
			Debug("Adding %v-%v to work list", e.from, e.to)
			// This ordering has to be enforced.
			workList.Push(e)
		})
	})

	Debug("Looking for an augmenting path.")
	for !workList.Empty() {
		cur := (workList.Pop()).(*Edge)
		Debug("Processing edge %v-%v.", cur.from, cur.to)
		if M.hasEdge(cur.from, cur.to) {
			continue
		}

		startInfo := forest[cur.from.toInt()]
		endInfo := forest[cur.to.toInt()]

		Debug("Got startInfo: %v, endInfo: %v", startInfo, endInfo)

		if nil != endInfo {
			if endInfo.IsOuter && startInfo.Root == endInfo.Root {
				// Case 1.
				// Both endpoints are outer nodes in the same tree -
				// a blossom is present.
				// Contract the blossom, repeat the search in the contracted graph
				// and expand the result.
				Debug("Case 1.: %v-%v", cur.from, cur.to)
				blossom := findBlossom(forest, cur)
				Debug("Found blossom %v", blossom.vertices)
				path := findAugmentingPath(
					contractGraph(G, blossom),
					contractGraph(M, blossom))

				if nil == path {
					return path
				}

				return expandPath(path, G, forest, blossom)
			} else if endInfo.IsOuter && startInfo.Root != endInfo.Root {
				// Case 2.
				// Both endpoints are outer nodes in different trees.
				// The augmenting path goes from the root of one tree
				// down through the other.
				// (root(v) → … → v) → (w → … → root(w))

				Debug("Case 2.: %v-%v", cur.from, cur.to)
				result = list.New()
				for v := cur.from; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
					// The path has to be added in reverse order.
					result.PushFront(v)
				}

				for v := cur.to; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
					// The path has to be added in reverse order.
					result.PushBack(v)
				}

				return result
			} else {
				// Case 3.
				// One endpoint is an outer node, and the second one is an inner node.
				// The path that we would end up taking from the root of the first tree
				// through this edge would not end up at the root of the other tree -
				// the only way it could be done while alternating would trail away from
				// the root.
				// This edge can be skipped.
				Debug("Case 3.: %v-%v", cur.from, cur.to)
			}
		} else {
			// There is no info on this edge - it must correspond to a matched
			// node, since all exposed nodes have been added to the forest.
			// The node can be added as an inner node to the tree
			// containing the start of the endpoint, then add the node for its
			// endpoint to the tree as an outer node.

			Debug("Corresponds to a matched vertex: %v-%v", cur.from, cur.to)
			forest[cur.to.toInt()] = mkNodeInformation(cur.from, startInfo.Root, false)

			// The endpoint of the unique matched edge corresp. to this node
			// will become an outer node of this tree.
			var endpoint Vertex
			M.ForAllNeighbors(cur.to, func(edge *Edge, index int, done chan<- bool) {
				endpoint = getOtherVertex(cur.to, edge)
				done <- true
			})

			forest[endpoint.toInt()] = mkNodeInformation(cur.to, startInfo.Root, true)

			G.ForAllNeighbors(endpoint, func(edge *Edge, index int, done chan<- bool) {
				e := MkEdge(endpoint, getOtherVertex(endpoint, edge))
				Debug("Adding fringe edge %v-%v to work list", e.from, e.to)
				workList.Push(e)
			})
		}
	}

	// Reaching here means that a maximum forest without finding any augmenting
	// paths has been created.
	return nil
}

func indexOf(el interface{}, list *list.List) int {
	for e, i := list.Front(), 0; e != nil; e = e.Next() {
		if e.Value == el {
			return i
		}

		i++
	}

	return -1
}

func updateMatching(path *list.List, matching *Graph) {
	// P ⊕ M
	for e, f := path.Front(), path.Front().Next(); f != nil; e, f = e.Next(), f.Next() {
		from, to := e.Value.(Vertex), f.Value.(Vertex)
		if matching.hasEdge(from, to) {
			Debug("Removing edge %v-%v from matching", from, to)
			matching.RemoveEdge(from, to)
		} else {
			Debug("Adding edge %v-%v to matching", from, to)
			matching.AddEdge(from, to)
		}
	}
}

func contractGraph(g *Graph, blossom *blossom) *Graph {
	result := MkGraph(g.currentVertexIndex)
	// Only the nodes not contained in a blossom belong to the result.
	g.ForAllVertices(func(v Vertex, index int, done chan<- bool) {
		if blossom.vertices.Contains(v) && v != blossom.Root {
			result.RemoveVertex(v)
		}
	})

	g.ForAllVertices(func(v Vertex, index int, done chan<- bool) {
		if blossom.vertices.Contains(v) {
			return
		}

		g.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
			w := getOtherVertex(v, edge)
			if blossom.vertices.Contains(w) {
				result.AddEdge(v, blossom.Root)
			} else {
				result.AddEdge(v, w)
			}
		})
	})

	return result
}

func expandPath(path *list.List, g *Graph, forest []*nodeInformation, blossom *blossom) *list.List {
	index := indexOf(blossom.Root, path)

	// If the blossom is not included in the path, it does not need
	// to be contracted.
	if index == -1 {
		return path
	}

	if index%2 == 1 {
		path = reverse(path)
	}

	result := list.New()
	for p := path.Front(); p != nil; p = p.Next() {
		v := (p.Value).(Vertex)

		if v != blossom.Root {
			result.PushBack(v)
		} else {
			result.PushBack(blossom.Root)
			exitNode := findBlossomExit(g, blossom, p.Next().Value.(Vertex))
			exitIndex := indexOf(exitNode, blossom.cycle)

			Debug("Exit node : %v, Exit index: %v", exitNode, exitIndex)
			var start, step int
			// The path taken to the exit of the cycle must end by following a
			// matched edge, to maintani the invariant of '{w', w} ∈ E ⧵ M'
			if exitIndex%2 == 0 {
				// The clockwise path will end in the matched edge.
				start = 1
				step = 1
			} else {
				// The anti-clockwise path will end in the matched edge.
				start = blossom.cycle.Len() - 2
				step = -1
			}

			for k := start; k != exitIndex+step; k += step {
				result.PushBack(get(blossom.cycle, k))
			}
		}
	}

	return result
}

// TODO: @refactor use an array instead.
func get(list *list.List, index int) interface{} {
	e := list.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
	}

	return e.Value
}

func reverse(input *list.List) (result *list.List) {
	result = list.New()

	for el := input.Front(); el != nil; el = el.Next() {
		result.PushFront(el.Value)
	}

	return result
}

func findBlossomExit(g *Graph, blossom *blossom, v Vertex) Vertex {
	for cv := range blossom.vertices.Iter() {
		cycleVertex := cv.(Vertex)
		if g.hasEdge(cycleVertex, v) {
			return cycleVertex
		}
	}

	panic(errors.New("The blossom has no exit in given graph!"))
}
