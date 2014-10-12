package graph

import (
	"fmt"

	"github.com/deckarep/golang-set"
	"github.com/lukpank/go-glpk/glpk"
)

type lpFormulation struct {
	g            *Graph
	k            int
	lp           *glpk.Prob
	coefficients [][]float64
}

type lpDualFormulation struct {
	lpFormulation
}

func (self *Edge) lpVarStr() string {
	return fmt.Sprintf("y(%v,%v)", self.from, self.to)
}

// Maximization dual of the LP kernelization:
// one variable y (u,v) for every edge (u, v)
// Maximize the sum of y(u,v)
// The sum of y (u,v) all edges containing v should be <= 1 (which is the weight of v)
// y (u,v) >= 0

// e.g. for graph 1-2, 1-3, 2-3, 3-4
// Max. sum of y(u,v)
// y(1,2) + y(1,3) <= 1
// y(1,2) + y(2,3) <= 1
// y(1,3) + y(2,3) + y(3,4) <= 1
// y(3,4) <= 1
// y(1,2), y(1,3), y(2,3), y(3,4) >= 0

func mklpDualFormulation(g *Graph, k int) (result *lpDualFormulation) {
	result = &lpDualFormulation{
		lpFormulation{
			g:  g,
			k:  k,
			lp: glpk.New(),
		},
	}

	result.lp.SetProbName("NT reduction")
	result.lp.SetObjName("sum:Y(u,v)")
	result.lp.SetObjDir(glpk.MAX)

	result.coefficients = make([][]float64, g.currentVertexIndex+1)
	for i, _ := range result.coefficients {
		result.coefficients[i] = make([]float64, g.NEdges()+1)
	}

	result.lp.AddRows(g.NVertices())
	i := 1
	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		Debug("Adding row %v (v%v)", i, v)
		result.lp.SetRowName(i, fmt.Sprintf("v%v", v))
		result.lp.SetRowBnds(i, glpk.UP, 0, 1)
		j := 1
		g.ForAllEdges(func(edge *Edge, done chan<- bool) {
			if edge.IsCoveredBy(v) {
				result.coefficients[v][j] = 1
			}

			j++
		})
		i++
	})

	Debug("Coefficients:\n%v", result.coefficients)

	result.lp.AddCols(g.NEdges())
	j := 1
	g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		result.lp.SetColName(j, edge.lpVarStr())
		result.lp.SetColBnds(j, glpk.LO, 0, 1) // the ub should not matter here.
		Debug("Col[%v]: %v", j, edge.lpVarStr())
		// All the edges belong to the objective function.
		result.lp.SetObjCoef(j, 1)
		j++
	})

	// Set the indices for the y(u,v) variables in the constraints.
	ind := make([]int32, j)
	for idx, _ := range ind {
		ind[idx] = int32(idx)
	}

	// Set the coefficients for the constraints.
	i = 1
	g.ForAllVertices(func(v Vertex, done chan<- bool) {
		result.lp.SetMatRow(i, ind, result.coefficients[v])
		Debug("Matrix[%v]:\n%v", i, result.coefficients[v])
		i++
	})

	return
}

func (self *lpDualFormulation) solve() (matching mapset.Set, err error) {
	err = self.lp.Simplex(nil)
	matching = mapset.NewSet()
	Debug("%s = %g", self.lp.ObjName(), self.lp.ObjVal())
	for i := 1; i <= self.g.NEdges(); i++ {
		val := self.lp.ColPrim(i)
		Debug("; %s = %g", self.lp.ColName(i), val)
		if val > 0 {
			matching.Add(self.g.Edges[i-1])
		}
	}

	return
}