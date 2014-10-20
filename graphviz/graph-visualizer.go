package graphviz

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

var allowedLayouts map[string]bool = map[string]bool{
	"dot":   true,
	"neato": true,
	"sfdp":  true,
}

var defaultOutputFormat string = "svg"

type graphVisualizer struct {
	g               *Graph
	edgeAttrs       [][]map[string]string
	vertexAttrs     []map[string]string
	layoutAlgorithm string
}

func ShowGraph(g *Graph) {
	MkGraphVisualizer(g).Display()
}

func mkGraphVisualizer(g *Graph, layoutAlgorithm string) *graphVisualizer {
	if !allowedLayouts[layoutAlgorithm] {
		panic(errors.New(fmt.Sprintf("Layout algorithm '%v' is not allowed.", layoutAlgorithm)))
	}

	result := &graphVisualizer{
		g:               g,
		layoutAlgorithm: layoutAlgorithm,
	}

	result.edgeAttrs = make([][]map[string]string, g.CurrentVertexIndex)
	for i, _ := range result.edgeAttrs {
		result.edgeAttrs[i] = make([]map[string]string, g.CurrentVertexIndex)
		for j, _ := range result.edgeAttrs[i] {
			result.edgeAttrs[i][j] = make(map[string]string)
		}
	}

	result.vertexAttrs = make([]map[string]string, g.CurrentVertexIndex)
	for i, _ := range result.vertexAttrs {
		result.vertexAttrs[i] = make(map[string]string)
	}

	return result
}

func MkGraphVisualizer(g *Graph) *graphVisualizer {
	if g.NVertices() > 100 {
		return mkGraphVisualizer(g, "sfdp")
	} else {
		return mkGraphVisualizer(g, "neato")
	}
}

func MkNeatoVisualizer(g *Graph) *graphVisualizer {
	return mkGraphVisualizer(g, "neato")
}

func MkSfdpVisualizer(g *Graph) *graphVisualizer {
	return mkGraphVisualizer(g, "sfdp")
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
	return tstobn(fmt.Sprintf("%v -- %v;", edge.From, edge.To))
}

func edgeToBWithAttrs(edge *Edge, attrs map[string]string) []byte {
	if len(attrs) > 0 {
		attrsStrs := make([]string, 0, len(attrs))

		for name, value := range attrs {
			attrsStrs = append(attrsStrs, fmt.Sprintf("%v=\"%v\"", name, value))
		}

		return tstobn(fmt.Sprintf("%v -- %v [%v];", edge.From, edge.To, strings.Join(attrsStrs, ", ")))
	} else {
		return edgeToB(edge)
	}
}

func vertexToB(v Vertex) []byte {
	return tstobn(fmt.Sprintf("%v;", v))
}

func (self *graphVisualizer) vertexToB(v Vertex) []byte {
	if len(self.vertexAttrs[v.ToInt()]) > 0 {
		attrsStrs := make([]string, 0, len(self.vertexAttrs))

		for name, value := range self.vertexAttrs[v.ToInt()] {
			attrsStrs = append(attrsStrs, fmt.Sprintf("%v=\"%v\"", name, value))
		}

		utility.Debug(strings.Join(attrsStrs, ", "))

		return tstobn(fmt.Sprintf("%v [%v];", v, strings.Join(attrsStrs, ", ")))
	} else {
		return vertexToB(v)
	}
}

func (self *graphVisualizer) edgeToB(edge *Edge) []byte {
	return edgeToBWithAttrs(edge, self.edgeAttrs[edge.From.ToInt()][edge.To.ToInt()])
}

func (self *graphVisualizer) setEdgeAttr(edge *Edge, name, val string) {
	self.setEdgeEndpointsAttr(edge.From, edge.To, name, val)
}

func (self *graphVisualizer) setEdgeEndpointsAttr(from, to Vertex, name, val string) {
	self.setEdgeCoordsAttr(from.ToInt(), to.ToInt(), name, val)
}

func (self *graphVisualizer) setEdgeCoordsAttr(from, to int, name, val string) {
	self.edgeAttrs[from][to][name] = val
	self.edgeAttrs[to][from][name] = val
}

func (self *graphVisualizer) HighlightEdge(edge *Edge, color string) {
	self.setEdgeAttr(edge, "color", color)
}

func (self *graphVisualizer) HighlightMatching(matching *Graph, color string) {
	matching.ForAllEdges(func(edge *Edge, done chan<- bool) {
		self.HighlightEdge(edge, color)
	})
}

func (self *graphVisualizer) HighlightMatchingSet(matching mapset.Set, color string) {
	for e := range matching.Iter() {
		self.HighlightEdge(e.(*Edge), color)
	}
}

func (self *graphVisualizer) HighlightVertex(v Vertex, color string) {
	self.vertexAttrs[v.ToInt()]["style"] = "filled"
	self.vertexAttrs[v.ToInt()]["fillcolor"] = color
}

func (self *graphVisualizer) HighlightCover(cover mapset.Set, color string) {
	for vInter := range cover.Iter() {
		self.HighlightVertex(vInter.(Vertex), color)
	}
}

func (self *graphVisualizer) HighlightCrown(crown *kernelization.Crown) {
	for vInter := range crown.I.Iter() {
		self.HighlightVertex(vInter.(Vertex), "lightgray")
	}

	for vInter := range crown.H.Iter() {
		self.HighlightVertex(vInter.(Vertex), "yellow")
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
		connectedVertices.Add(edge.From)
		connectedVertices.Add(edge.To)
	})

	for _, v := range self.g.Vertices {
		if connectedVertices.Contains(v) || verticesWithAttrs.Contains(v) {
			continue
		}

		if self.g.HasVertex(v) {
			res.Write(self.vertexToB(v))
		}
	}

	res.Write(stob("}"))
	return res
}

func (self *graphVisualizer) dotToImage(dot bytes.Buffer) bytes.Buffer {
	var res bytes.Buffer
	cmd := exec.Command(self.layoutAlgorithm, "-T", defaultOutputFormat)
	cmd.Stdout = &res
	cmd.Stdin = bytes.NewReader(dot.Bytes())
	utility.Debug("Converting dot to %v...", defaultOutputFormat)
	err := cmd.Run()
	if nil != err {
		log.Fatal(err)
		return res
	}

	utility.Debug("Converted dot to %v.", defaultOutputFormat)
	return res
}

func (self *graphVisualizer) mkImage(name string) bytes.Buffer {
	return self.dotToImage(self.toDot(name))
}

func (self *graphVisualizer) MkImage(name string) error {
	file, err := os.Create(fmt.Sprintf("%v.%v", name, defaultOutputFormat))
	if nil != err {
		return err
	}

	defer file.Close()
	buf := self.mkImage(name)
	_, err = file.Write(buf.Bytes())
	return err
}

func (self *graphVisualizer) Display() {
	randname := fmt.Sprintf("%v", rand.Int63())
	filename := fmt.Sprintf("%v.%v", randname, defaultOutputFormat)
	cmd := exec.Command("feh", filename)
	err := self.MkImage(randname)
	if nil != err {
		log.Fatal(err)
	}

	defer os.Remove(filename)

	err = cmd.Run()
	if nil != err {
		log.Fatal(err)
	}
}
