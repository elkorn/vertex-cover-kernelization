package kernelization

import (
	"sort"

	"github.com/elkorn/vertex-cover-kernelization/graph"
)

// Tags are related to fold-struction.go.

type tag struct {
	v         graph.Vertex
	neighbors graph.Neighbors
	g         *graph.Graph
}

func (a tag) Len() int      { return len(a.neighbors) }
func (a tag) Swap(i, j int) { a.neighbors[i], a.neighbors[j] = a.neighbors[j], a.neighbors[i] }
func (a tag) Less(i, j int) bool {
	return a.g.Degree(a.neighbors[i]) > a.g.Degree(a.neighbors[j])
}

func (self *tag) Compare(other *tag, g *graph.Graph) int {
	selfN, otherN := len(self.neighbors), len(other.neighbors)
	for i := 0; i < selfN && i < otherN; i++ {
		dSelf, dOther := g.Degree(self.neighbors[i]), g.Degree(other.neighbors[i])
		if dSelf > dOther {
			return 1
		} else if dSelf < dOther {
			return -1
		}
	}

	// In lexicographic comparison, if the words are equal up to this point,
	// the longer one is greater than the shorter one.
	if selfN > otherN {
		return 1
	} else if selfN < otherN {
		return -1
	}

	return 0
}

func MkTag(v graph.Vertex, g *graph.Graph) *tag {
	result := &tag{
		v:         v,
		g:         g,
		neighbors: g.GetNeighbors(v),
	}

	sort.Sort(result)
	return result
}

func computeTags(g *graph.Graph) []*tag {
	result := make([]*tag, g.CurrentVertexIndex)
	g.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		result[v.ToInt()] = MkTag(v, g)
	})

	return result
}
