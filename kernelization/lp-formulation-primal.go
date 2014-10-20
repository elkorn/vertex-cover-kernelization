package kernelization

import (
	"fmt"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/lukpank/go-glpk/glpk"
)

type lpPrimalFormulation struct {
	lpFormulation
}

func asPrimalLpVar(v graph.Vertex) string {
	return fmt.Sprintf("X%v", v)
}

func asPrimalLpConstraint(edge *graph.Edge) string {
	return fmt.Sprintf("%v + %v >= 1",
		asPrimalLpVar(edge.From),
		asPrimalLpVar(edge.To))
}

// Assign a value Xu >= 0 to each vertex u.
// Minimize sum of Xu over all vertices.
// Xu+Xv >= 1 for every edge (u,v).

// e.g. for graph 1-2, 1-3, 2-3, 3-4
// X1 + X2 >= 1
// X1 + X3 >= 1
// X2 + X3 >= 1
// X3 + X4 >= 1
//
func mklpPrimalFormulation(g *graph.Graph, k int) (result *lpPrimalFormulation) {
	result = &lpPrimalFormulation{
		lpFormulation{
			g:  g,
			k:  k,
			lp: glpk.New(),
		},
	}

	result.lp.SetProbName("NT reduction")
	result.lp.SetObjName("sum X(u)")
	result.lp.SetObjDir(glpk.MIN)

	result.coefficients = make([][]float64, g.NEdges())
	for i, _ := range result.coefficients {
		result.coefficients[i] = make([]float64, g.CurrentVertexIndex+1)
	}

	result.lp.AddRows(g.NEdges())
	i := 1
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		utility.Debug("Constraint: %v", asPrimalLpConstraint(edge))
		result.lp.SetRowName(i, fmt.Sprintf("%v + %v >= 1", asPrimalLpVar(edge.From), asPrimalLpVar(edge.To)))
		result.lp.SetRowBnds(i, glpk.LO, 1, 1)
		result.coefficients[i-1][edge.From] = 1
		result.coefficients[i-1][edge.To] = 1
		utility.Debug("Coeff.: %v", result.coefficients[i-1])
		i++
	})

	result.lp.AddCols(g.CurrentVertexIndex)
	j := 1
	g.ForAllVertices(func(v graph.Vertex, done chan<- bool) {
		result.lp.SetColName(j, asPrimalLpVar(v))
		result.lp.SetColBnds(j, glpk.LO, 0, 0) // edge.Str(), the ub should not matter here.
		utility.Debug("Col[%v]: %v", j, asPrimalLpVar(v))
		// All the vertices belong to the objective function.
		result.lp.SetObjCoef(j, 1)
		j++
	})

	// Set the indices for the y(u,v) variables in the constraints.
	ind := make([]int32, j)
	for idx, _ := range ind {
		ind[idx] = int32(idx)
	}

	utility.Debug("ind: %v", ind)

	// Set the coefficients for the constraints.
	i = 1
	g.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		// SetMatRow sets (replaces) i-th row. It sets
		//
		//     matrix[i, ind[j]] = val[j]
		//
		// for j=1..len(ind). ind[0] and val[0] are ignored. Requires
		// len(ind) = len(val).
		// !!!! ind[0] and val[0] are ignored !!!!
		result.lp.SetMatRow(i, ind, result.coefficients[i-1])
		utility.Debug("Coeff.[%v (%v)]:\n%v", i, asPrimalLpConstraint(edge), result.coefficients[i-1])
		i++
	})

	return
}

func (self *lpPrimalFormulation) solve() (err error) {
	smcp := glpk.NewSmcp()
	smcp.SetMsgLev(glpk.MSG_ERR)
	err = self.lp.Simplex(smcp)
	utility.Debug("%s = %g", self.lp.ObjName(), self.lp.ObjVal())
	return
}
