package graph

import "github.com/lukpank/go-glpk/glpk"

// TODO: @start-from-here use the Operational Research book to guide yourself
// through implementing the NT-decomposition.
// Decide what should be returned (probably 3 sets of vertices)

func ntReduction(g *Graph, k int) {
	lp := glpk.New()
	lp.SetProbName("NT decomposition")
}
