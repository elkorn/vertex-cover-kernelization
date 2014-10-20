package matching

import (
	"container/list"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type blossom struct {
	Root     graph.Vertex
	cycle    *list.List
	vertices mapset.Set
}

func MkBlossom(root graph.Vertex, cycle *list.List, vertices mapset.Set) *blossom {
	return &blossom{
		Root:     root,
		cycle:    cycle,
		vertices: vertices,
	}
}

func prepend(slice []graph.Vertex, v graph.Vertex) []graph.Vertex {
	return append([]graph.Vertex{v}, slice...)
}

func printForest(forest []*NodeInformation) {
	for i, f := range forest {
		if nil == f {
			utility.Debug("graph.Vertex %v is not in the forest. ", graph.MkVertex(i))
		} else {
			utility.Debug("%v - parent: %v, root, %v, outer: %v", graph.MkVertex(i), f.Parent, f.Root, f.IsOuter)
		}
	}
}

func findBlossom(forest []*NodeInformation, edge *graph.Edge) *blossom {
	// A maximum matching has cardinality at most n/2.
	n := len(forest)
	utility.Debug("Searching for blossom with edge %v in forest", edge)
	printForest(forest)
	pathA, pathB := make([]graph.Vertex, 0, n/2), make([]graph.Vertex, 0, n/2)
	for v := edge.From; v != graph.INVALID_VERTEX; v = forest[v.ToInt()].Parent {
		utility.Debug("Add %v to path A", v)
		pathA = prepend(pathA, v)
	}

	for v := edge.To; v != graph.INVALID_VERTEX; v = forest[v.ToInt()].Parent {
		utility.Debug("Add %v to path B", v)
		pathB = prepend(pathB, v)
	}

	commonAncestorIdx := 0
	nA, nB := len(pathA), len(pathB)
	utility.Debug("len(A): %v, len(B): %v", nA, nB)
	for diffIdx := 0; diffIdx < nA && diffIdx < nB; diffIdx++ {
		utility.Debug("[Ancestor %v] A: %v, B: %v", diffIdx, pathA[diffIdx], pathB[diffIdx])
		if pathA[diffIdx] != pathB[diffIdx] {
			commonAncestorIdx = diffIdx - 1
			utility.Debug("Common ancestor index: %v", commonAncestorIdx)
			break
		}
	}

	// Both nodes belong to the same tree, they have the same root,
	// what guarantees having a non-zero common ancestor index.
	cycle := list.New()
	vertices := mapset.NewSet()
	for i := commonAncestorIdx; i < nA; i++ {
		v := pathA[i]
		cycle.PushBack(v)
		vertices.Add(v)
	}

	for i := nB - 1; i >= commonAncestorIdx; i-- {
		v := pathB[i]
		cycle.PushBack(v)
		vertices.Add(v)
	}

	utility.Debug("Blossom cycle: %v", cycle)
	return MkBlossom(pathA[commonAncestorIdx], cycle, vertices)
}
