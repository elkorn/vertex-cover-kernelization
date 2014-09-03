package graph

import (
	"container/list"

	"github.com/deckarep/golang-set"
)

type blossom struct {
	Root     Vertex
	cycle    *list.List
	vertices mapset.Set
}

func MkBlossom(root Vertex, cycle *list.List, vertices mapset.Set) *blossom {
	return &blossom{
		Root:     root,
		cycle:    cycle,
		vertices: vertices,
	}
}

func prepend(slice []Vertex, v Vertex) []Vertex {
	return append([]Vertex{v}, slice...)
}

func printForest(forest []*nodeInformation) {
	for i, f := range forest {
		Debug("%v - parent: %v, root, %v, outer: %v", MkVertex(i), f.Parent, f.Root, f.IsOuter)
	}
}

func findBlossom(forest []*nodeInformation, edge *Edge) *blossom {
	// A maximum matching has cardinality at most n/2.
	n := len(forest)
	Debug("Searching for blossom with edge %v in forest", edge)
	printForest(forest)
	pathA, pathB := make([]Vertex, 0, n/2), make([]Vertex, 0, n/2)
	for v := edge.from; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
		Debug("Add %v to path A", v)
		pathA = prepend(pathA, v)
	}

	for v := edge.to; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
		Debug("Add %v to path B", v)
		pathB = prepend(pathB, v)
	}

	commonAncestorIdx := 0
	nA, nB := len(pathA), len(pathB)
	Debug("len(A): %v, len(B): %v", nA, nB)
	for diffIdx := 0; diffIdx < nA && diffIdx < nB; diffIdx++ {
		Debug("[Ancestor %v] A: %v, B: %v", diffIdx, pathA[diffIdx], pathB[diffIdx])
		if pathA[diffIdx] != pathB[diffIdx] {
			commonAncestorIdx = diffIdx - 1
			Debug("Common ancestor index: %v", commonAncestorIdx)
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

	Debug("Blossom cycle: %v", cycle)
	return MkBlossom(pathA[commonAncestorIdx], cycle, vertices)
}
