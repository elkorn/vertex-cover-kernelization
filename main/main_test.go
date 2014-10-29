package main

import (
	"fmt"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/graphviz"
)

func TestGenExamples(t *testing.T) {
	for i := 1; i <= 20; i++ {
		name := fmt.Sprintf("ex_%v", i*100)
		g := graph.ScanGraph(name)
		graphviz.MkGraphVisualizer(g).SaveDot(fmt.Sprintf("%v.dot", name))
	}
}
