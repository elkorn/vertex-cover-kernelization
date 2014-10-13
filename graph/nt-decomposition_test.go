package graph

import "testing"

func TestNtDecomposition(t *testing.T) {
	g := mkPetersenGraph()
	nt := mkNtDecomposition(g, 6)
	InVerboseContext(func() {
		Debug(nt.Str())
	})
}
