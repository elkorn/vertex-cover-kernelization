package graph

import "testing"

func TestNtReductionYehuda(t *testing.T) {
	g := mkPetersenGraph()
	nt := mkNtReductionYehuda(g, 10)
	Debug("V0: %v", nt.V0)
	Debug("C0: %v", nt.C0)
}
