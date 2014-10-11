package graph

import "github.com/lukpank/go-glpk/glpk"

func ntReduction(g *Graph, k int) {
	lp := glpk.New()
	lp.SetProbName("NT decomposition")
}
