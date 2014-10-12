package graph

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/lukpank/go-glpk/glpk"
)

type lpPrimalFormulation struct {
	lpFormulation
}

func asPrimalLpVar(v Vertex) string {
	return fmt.Sprintf("X%v", v)
}

func asPrimalLpConstraint(edge *Edge) string {
	return fmt.Sprintf("%v + %v >= 1",
		asPrimalLpVar(edge.from),
		asPrimalLpVar(edge.to))
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
func mklpPrimalFormulation(g *Graph, k int) (result *lpPrimalFormulation) {
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

	result.coefficients = make([][]float64, g.NEdges()+1)
	for i, _ := range result.coefficients {
		result.coefficients[i] = make([]float64, g.currentVertexIndex+1)
	}

	result.lp.AddRows(g.NEdges())
	i := 1
	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		Debug("Constraint: %v", asPrimalLpConstraint(edge))
		result.lp.SetRowName(i, fmt.Sprintf("%v + %v >= 1", asPrimalLpVar(edge.from), asPrimalLpVar(edge.to)))
		result.lp.SetRowBnds(i, glpk.LO, 1, 1)
		result.coefficients[i][edge.from] = 1
		result.coefficients[i][edge.to] = 1
		Debug("Coeff.: %v", result.coefficients[i])
		i++
	})

	result.lp.AddCols(g.currentVertexIndex)
	j := 1
	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		result.lp.SetColName(j, asPrimalLpVar(v))
		result.lp.SetColBnds(j, glpk.LO, 0, 0) // edge.Str(), the ub should not matter here.
		Debug("Col[%v]: %v", j, asPrimalLpVar(v))
		// All the vertices belong to the objective function.
		result.lp.SetObjCoef(j, 1)
		j++
	})

	// Set the indices for the y(u,v) variables in the constraints.
	ind := make([]int32, j)
	for idx, _ := range ind {
		ind[idx] = int32(idx)
	}

	Debug("ind: %v", ind)

	// Set the coefficients for the constraints.
	i = 1
	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		// SetMatRow sets (replaces) i-th row. It sets
		//
		//     matrix[i, ind[j]] = val[j]
		//
		// for j=1..len(ind). ind[0] and val[0] are ignored. Requires
		// len(ind) = len(val).
		// !!!! ind[0] and val[0] are ignored !!!!
		result.lp.SetMatRow(i, ind, result.coefficients[i])
		Debug("Coeff.[%v (%v)]:\n%v", i, asPrimalLpConstraint(edge), result.coefficients[i])
		i++
	})

	return
}

func (self *lpPrimalFormulation) solve() (P, Q, R mapset.Set, err error) {
	err = self.lp.Simplex(nil)
	P, Q, R = mapset.NewSet(), mapset.NewSet(), mapset.NewSet()
	Debug("%s = %g", self.lp.ObjName(), self.lp.ObjVal())
	self.g.ForAllVertices(func(v Vertex, done chan<- bool) {
		i := int(v)
		val := self.lp.ColPrim(i)
		Debug("; %s = %g", self.lp.ColName(i), val)
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
