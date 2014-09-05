package graph

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/deckarep/golang-set"
)

type graphVisualizer struct {
	g           *Graph
	edgeAttrs   [][]map[string]string
	vertexAttrs []map[string]string
}

func showGraph(g *Graph) {
	MkGraphVisualizer(g).Display()
}

func MkGraphVisualizer(g *Graph) *graphVisualizer {
	result := &graphVisualizer{
		g: g,
	}

	result.edgeAttrs = make([][]map[string]string, g.currentVertexIndex)
	for i, _ := range result.edgeAttrs {
		result.edgeAttrs[i] = make([]map[string]string, g.currentVertexIndex)
		for j, _ := range result.edgeAttrs[i] {
			result.edgeAttrs[i][j] = make(map[string]string)
		}
	}

	result.vertexAttrs = make([]map[string]string, g.currentVertexIndex)
	for i, _ := range result.vertexAttrs {
		result.vertexAttrs[i] = make(map[string]string)
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

func edgeToBWithAttrs(edge *Edge, attrs map[string]string) []byte {
	if len(attrs) > 0 {
		attrsStrs := make([]string, 0, len(attrs))

		for name, value := range attrs {
			attrsStrs = append(attrsStrs, fmt.Sprintf("%v=\"%v\"", name, value))
		}

		return tstobn(fmt.Sprintf("%v -- %v [%v];", edge.from, edge.to, strings.Join(attrsStrs, ", ")))
	} else {
		return edgeToB(edge)
	}
}

func vertexToB(v Vertex) []byte {
	return tstobn(fmt.Sprintf("%v;", v))
}

func (self *graphVisualizer) vertexToB(v Vertex) []byte {
	if len(self.vertexAttrs[v.toInt()]) > 0 {
		attrsStrs := make([]string, 0, len(self.vertexAttrs))

		for name, value := range self.vertexAttrs[v.toInt()] {
			attrsStrs = append(attrsStrs, fmt.Sprintf("%v=\"%v\"", name, value))
		}

		Debug(strings.Join(attrsStrs, ", "))

		return tstobn(fmt.Sprintf("%v [%v];", v, strings.Join(attrsStrs, ", ")))
	} else {
		return vertexToB(v)
	}
}

func (self *graphVisualizer) edgeToB(edge *Edge) []byte {
	return edgeToBWithAttrs(edge, self.edgeAttrs[edge.from.toInt()][edge.to.toInt()])
}

func (self *graphVisualizer) setEdgeAttr(edge *Edge, name, val string) {
	self.setEdgeEndpointsAttr(edge.from, edge.to, name, val)
}

func (self *graphVisualizer) setEdgeEndpointsAttr(from, to Vertex, name, val string) {
	self.setEdgeCoordsAttr(from.toInt(), to.toInt(), name, val)
}

func (self *graphVisualizer) setEdgeCoordsAttr(from, to int, name, val string) {
	self.edgeAttrs[from][to][name] = val
	self.edgeAttrs[to][from][name] = val
}

func (self *graphVisualizer) HighlightEdge(edge *Edge, color string) {
	self.setEdgeAttr(edge, "color", color)
}

func (self *graphVisualizer) HighlightVertex(v Vertex, color string) {
	self.vertexAttrs[v.toInt()]["style"] = "filled"
	self.vertexAttrs[v.toInt()]["fillcolor"] = color
}

func (self *graphVisualizer) HighlightCover(cover mapset.Set, color string) {
	for vInter := range cover.Iter() {
		self.HighlightVertex(vInter.(Vertex), color)
	}
}

func (self *graphVisualizer) toDot(name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	verticesWithAttrs := mapset.NewSet()
	for i, attrs := range self.vertexAttrs {
		if len(attrs) == 0 {
			continue
		}

		v := MkVertex(i)
		verticesWithAttrs.Add(v)
		res.Write(self.vertexToB(v))
	}

	connectedVertices := mapset.NewSet()
	self.g.ForAllEdges(func(edge *Edge, done chan<- bool) {
		// In this context it might be useful to use this range loop and e.g. display
		// the removed edge as dotted or grayed out.
		// for _, edge := range g.Edges {
		res.Write(self.edgeToB(edge))
		connectedVertices.Add(edge.from)
		connectedVertices.Add(edge.to)
	})

	for _, v := range self.g.Vertices {
		if connectedVertices.Contains(v) || verticesWithAttrs.Contains(v) {
			continue
		}

		if self.g.hasVertex(v) {
			res.Write(self.vertexToB(v))
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
	Debug("Converting dot to jpg...")
	err := cmd.Run()
	if nil != err {
		log.Fatal(err)
		return res
	}

	Debug("Converted dot to jpg.")
	return res
}

func (self *graphVisualizer) mkJpg(name string) bytes.Buffer {
	return self.dotToJpg(self.toDot(name))
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
