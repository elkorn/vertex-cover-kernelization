package graph

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"github.com/deckarep/golang-set"
)

type gvAttr struct {
	key   string
	value string
}

type graphVisualizer struct {
	g     *Graph
	attrs [][]*gvAttr
}

func showGraph(g *Graph) {
	MkGraphVisualizer(g).Display()
}

func MkGraphVisualizer(g *Graph) *graphVisualizer {
	result := &graphVisualizer{
		g: g,
	}

	result.attrs = make([][]*gvAttr, g.currentVertexIndex)
	for i, _ := range result.attrs {
		result.attrs[i] = make([]*gvAttr, g.currentVertexIndex)
	}

	return result
}

func stob(str string) []byte {
	return []byte(str)
}

func stobn(str string) []byte {
	return append(stob(str), '\n')
}

func tstobn(str string) []byte {
	res := new(bytes.Buffer)
	res.Write(stob("\t"))
	res.Write(stobn(str))
	return res.Bytes()
}

func edgeToB(edge *Edge) []byte {
	return tstobn(fmt.Sprintf("%v -- %v;", edge.from, edge.to))
}

func vertexToB(v Vertex) []byte {
	return tstobn(fmt.Sprintf("%v;", v))
}

func (self *graphVisualizer) toDot(g *Graph, name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	connectedVertices := mapset.NewSet()
	g.ForAllEdges(func(edge *Edge, _ int, done chan<- bool) {
		// In this context it might be useful to use this range loop and e.g. display
		// the removed edge as dotted or grayed out.
		// for _, edge := range g.Edges {
		res.Write(edgeToB(edge))
		connectedVertices.Add(edge.from)
		connectedVertices.Add(edge.to)
	})

	for _, v := range g.Vertices {
		if connectedVertices.Contains(v) {
			continue
		}

		if g.hasVertex(v) {
			res.Write(vertexToB(v))
		}
	}

	res.Write(stob("}"))
	return res
}

func (self *graphVisualizer) dotToJpg(dot bytes.Buffer) bytes.Buffer {
	var res bytes.Buffer
	cmd := exec.Command("dot", "-T", "jpg")
	cmd.Stdout = &res
	cmd.Stdin = bytes.NewReader(dot.Bytes())
	err := cmd.Run()
	if nil != err {
		log.Fatal(err)
		return res
	}

	return res
}

func (self *graphVisualizer) mkJpg(name string) bytes.Buffer {
	return self.dotToJpg(self.toDot(self.g, name))
}

func (self *graphVisualizer) MkJpg(name string) error {
	file, err := os.Create(fmt.Sprintf("%v.jpg", name))
	if nil != err {
		return err
	}

	defer file.Close()
	buf := self.mkJpg(name)
	_, err = file.Write(buf.Bytes())
	return err
}

func (self *graphVisualizer) Display() {
	randname := fmt.Sprintf("%v", rand.Int63())
	filename := fmt.Sprintf("%v.jpg", randname)
	cmd := exec.Command("feh", filename)
	err := self.MkJpg(randname)
	if nil != err {
		log.Fatal(err)
	}

	defer os.Remove(filename)

	err = cmd.Run()
	if nil != err {
		log.Fatal(err)
	}
}
