package graph

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDot(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	expected := `graph test {
	1 -- 2;
	1 -- 3;
	2 -- 3;
}`

	result := g.ToDot("test")
	assert.Equal(t, expected, result.String())
}

func TestDotToJpg(t *testing.T) {
	g := mkGraphWithVertices(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	dot := g.ToDot("test")
	file, err := os.Open("expected_dot.jpg")
	expected := make([]byte, 0)
	if nil != err {
		panic("Cannot open file.")
	}

	defer file.Close()

	file.Read(expected)
	actual := DotToJpg(dot)
	assert.Equal(t, expected, actual.Bytes())
}
