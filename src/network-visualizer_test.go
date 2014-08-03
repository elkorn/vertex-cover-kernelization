package graph

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToGraph(t *testing.T) {
	expected := mkGraphWithVertices(4)
	expected.AddEdge(1, 2)
	expected.AddEdge(1, 3)
	expected.AddEdge(2, 4)

	net := mkNet(expected)
	actual := convertToGraph(&net)
	for vertex := range expected.Vertices {
		assert.True(t, actual.hasVertex(vertex))
	}

	for _, edge := range expected.Edges {
		assert.True(t, actual.hasEdge(edge.from, edge.to))
	}
}

func TestMkJpgFromNet(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	net := mkNet(g)

	expectedFile, err := os.Open("expected_dot.jpg")
	if nil != err {
		panic(err)
	}

	defer expectedFile.Close()
	ACTUAL_NAME := "net.jpg"
	nv := MkNetworkVisualizer()
	err = nv.MkJpg(&net)
	if nil != err {
		panic(err)
	}

	actualFile, err := os.Open(ACTUAL_NAME)
	if nil != err {
		panic(err)
	}

	defer actualFile.Close()

	expected, actual := make([]byte, 7000), make([]byte, 7000)
	expectedFile.Read(expected)
	actualFile.Read(actual)
	// assert.Equal(t, expected, actual)
	os.Remove(ACTUAL_NAME)
}
