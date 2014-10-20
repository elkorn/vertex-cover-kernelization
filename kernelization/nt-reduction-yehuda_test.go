package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

func TestNtReductionYehuda(t *testing.T) {
	g := graph.MkPetersenGraph()
	nt := mkNtReductionYehuda(g, 10)
	utility.Debug("V0: %v", nt.V0)
	utility.Debug("C0: %v", nt.C0)
}
