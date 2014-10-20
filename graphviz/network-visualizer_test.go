package graphviz

// import (
// 	"os"
// 	"testing"

// 	// "github.com/stretchr/testify/assert"
// )

// func TestColorEdges(t *testing.T) {
// 	g := MkGraph(3)
// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 3)
// 	net := mkNet(g)

// 	net.arcs[0][2].flow = 1
// 	nv := MkNetworkVisualizer()
// 	nv.Display(&net)
// }

// func TestMkJpgFromNet(t *testing.T) {
// 	g := MkGraph(3)
// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 3)
// 	net := mkNet(g)

// 	expectedFile, err := os.Open("expected_dot.svg")
// 	if nil != err {
// 		panic(err)
// 	}

// 	defer expectedFile.Close()
// 	ACTUAL_NAME := "net.svg"
// 	nv := MkNetworkVisualizer()
// 	err = nv.MkJpg(&net, "net")
// 	if nil != err {
// 		panic(err)
// 	}

// 	actualFile, err := os.Open(ACTUAL_NAME)
// 	if nil != err {
// 		panic(err)
// 	}

// 	defer actualFile.Close()

// 	expected, actual := make([]byte, 7000), make([]byte, 7000)
// 	expectedFile.Read(expected)
// 	actualFile.Read(actual)
// 	// assert.Equal(t, expected, actual)
// 	os.Remove(ACTUAL_NAME)
// }
