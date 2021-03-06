package kernelization

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

type ntReductionAbuKhzam struct {
	formulation *lpPrimalFormulation
}

func mkNtReductionAbuKhzam(g *graph.Graph, k int) *ntReductionAbuKhzam {
	return &ntReductionAbuKhzam{
		formulation: mklpPrimalFormulation(g, k),
	}
}

func (self *ntReductionAbuKhzam) solve() (P, Q, R mapset.Set, err error) {
	err = self.formulation.solve()
	P, Q, R = mapset.NewThreadUnsafeSet(), mapset.NewThreadUnsafeSet(), mapset.NewThreadUnsafeSet()
	self.formulation.g.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		i := int(v)
		val := self.formulation.lp.ColPrim(i)
		utility.Debug("; %s = %g", self.formulation.lp.ColName(i), val)
		switch true {
		case val > 0.5:
			P.Add(v)
			break
		case val == 0.5:
			Q.Add(v)
			break
		case val < 0.5:
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
