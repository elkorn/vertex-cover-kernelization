package graph

import (
	"fmt"

	"github.com/deckarep/golang-set"
)

type ntReduction struct {
	formulation *lpPrimalFormulation
}

func mkNtReduction(g *Graph, k int) *ntReduction {
	return &ntReduction{
		formulation: mklpPrimalFormulation(g, k),
	}
}

func (self *ntReduction) solve() (P, Q, R mapset.Set, err error) {
	err = self.formulation.solve()
	P, Q, R = mapset.NewSet(), mapset.NewSet(), mapset.NewSet()
	self.formulation.g.ForAllVertices(func(v Vertex, done chan<- bool) {
		i := int(v)
		val := self.formulation.lp.ColPrim(i)
		Debug("; %s = %g", self.formulation.lp.ColName(i), val)
		switch true {
		case val == 1:
			P.Add(v)
			break
		case val == 0.5:
			Q.Add(v)
			break
		case val == 0:
			R.Add(v)
			break
		default:
			panic(
				fmt.Sprintf(
					"Undefined case for val: %v (%v)",
					val,
					asPrimalLpVar(v)))
		}
	})

	return
}
