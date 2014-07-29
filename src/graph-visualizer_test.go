package graph

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesttoDot(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	expected := `graph test {
	1 -- 2;
	1 -- 3;
	2 -- 3;
}`

	result := toDot(g, "test")
	assert.Equal(t, expected, result.String())
}

func TestdotToJpg(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	dot := toDot(g, "test")
	file, err := os.Open("expected_dot.jpg")
	expected := make([]byte, 0)
	if nil != err {
		panic("Cannot open reference file.")
	}

	defer file.Close()

	file.Read(expected)
	actual := dotToJpg(dot)
	assert.Equal(t, expected, actual.Bytes())
}

func TestMkJpg(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	expectedFile, err := os.Open("expected_dot.jpg")
	if nil != err {
		panic(err)
	}

	defer expectedFile.Close()

	err = MkJpg(g, "actual_dot")
	if nil != err {
		panic(err)
	}

	actualFile, err := os.Open("actual_dot.jpg")
	if nil != err {
		panic(err)
	}

	defer actualFile.Close()

	expected, actual := make([]byte, 7000), make([]byte, 7000)
	expectedFile.Read(expected)
	actualFile.Read(actual)
	assert.Equal(t, expected, actual)
	os.Remove("actual_dot.jpg")
}
