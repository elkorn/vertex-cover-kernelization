package graphviz

import (
	"os"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/stretchr/testify/assert"
)

func TesttoDot(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	gv := MkGraphVisualizer(g)

	expected := `graph test {
	1 -- 2;
	1 -- 3;
	2 -- 3;
}`

	result := gv.toDot("test")
	assert.Equal(t, expected, result.String())
}

func TestdotToImage(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	gv := MkGraphVisualizer(g)

	dot := gv.toDot("test")
	file, err := os.Open("expected_dot." + GetOutputFormat())
	expected := make([]byte, 0)
	if nil != err {
		panic("Cannot open reference file.")
	}

	defer file.Close()

	file.Read(expected)
	actual := gv.dotToImage(dot)
	utility.Debug("", actual)
	// assert.Equal(t, expected, actual.Bytes())
}

func TestMkImage(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	gv := MkGraphVisualizer(g)

	expectedFile, err := os.Open("expected_dot." + GetOutputFormat())
	if nil != err {
		panic(err)
	}

	defer expectedFile.Close()

	err = gv.MkImage("actual_dot")
	if nil != err {
		panic(err)
	}

	actualFile, err := os.Open("actual_dot." + GetOutputFormat())
	if nil != err {
		panic(err)
	}

	defer actualFile.Close()

	expected, actual := make([]byte, 7000), make([]byte, 7000)
	expectedFile.Read(expected)
	actualFile.Read(actual)
	// assert.Equal(t, expected, actual)
	os.Remove("actual_dot." + GetOutputFormat())
}

func TestColor(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	gv := MkGraphVisualizer(g)

	gv.HighlightEdge(g.Edges[0], "red")
	gv.HighlightEdge(g.Edges[1], "green")
	gv.HighlightEdge(g.Edges[2], "purple")
	// gv.Display()
}

func TestSaveDot(t *testing.T) {
	g := graph.ScanGraph("../examples/sh2/sh2-3.dim")
	gv := MkGraphVisualizer(g)
	gv.SaveDot("TestSaveDot.dot")
	g2 := graph.ScanDot("TestSaveDot.dot")
	assert.Equal(t, g.NEdges(), g2.NEdges())
}
