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

func findBlossom(forest []*nodeInformation, edge *Edge) *blossom {
	pathA, pathB := MkStack(len(forest)), MkStack(len(forest))
	for v := edge.from; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
		pathA.Push(v)
	}

	for v := edge.to; v != INVALID_VERTEX; v = forest[v.toInt()].Parent {
		pathB.Push(v)
	}

	commonAncestorIdx := 0
	for diffIdx := 0; diffIdx < pathA.Size() && diffIdx < pathB.Size(); diffIdx++ {
		if pathA.Peek(diffIdx) != pathB.Peek(diffIdx) {
			commonAncestorIdx = diffIdx - 1
			break
		}
	}

	// Both nodes belong to the same tree, they have the same root,
	// what guarantees having a non-zero common ancestor index.
	cycle := list.New()
	vertices := mapset.NewSet()
	for i := commonAncestorIdx; i < pathA.Size(); i++ {
		v := pathA.Peek(i)
		cycle.PushBack(v)
		vertices.Add(v)
	}

	for i := pathB.Size() - 1; i >= commonAncestorIdx; i-- {
		v := pathB.Peek(i)
		cycle.PushBack(v)
		vertices.Add(v)
	}

	return MkBlossom((pathA.Peek(commonAncestorIdx)).(Vertex), cycle, vertices)
}
