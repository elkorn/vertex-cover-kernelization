package graph

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

type graphVisualizer struct{}

func MkGraphVisualizer() *graphVisualizer {
	return &graphVisualizer{}
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

func (self *graphVisualizer) toDot(g *Graph, name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	for _, edge := range g.Edges {
		res.Write(edgeToB(edge))
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

func (self *graphVisualizer) mkJpg(g *Graph, name string) bytes.Buffer {
	return self.dotToJpg(self.toDot(g, name))
}

func (self *graphVisualizer) MkJpg(g *Graph, name string) error {
	file, err := os.Create(fmt.Sprintf("%v.jpg", name))
	if nil != err {
		return err
	}

	defer file.Close()
	buf := self.mkJpg(g, name)
	_, err = file.Write(buf.Bytes())
	return err
}

func (self *graphVisualizer) Display(g *Graph) {
	randname := fmt.Sprintf("%v", rand.Int63())
	filename := fmt.Sprintf("%v.jpg", randname)
	cmd := exec.Command("feh", filename)
	err := self.MkJpg(g, randname)
	if nil != err {
		log.Fatal(err)
	}

	defer os.Remove(filename)

	err = cmd.Run()
	if nil != err {
		log.Fatal(err)
	}
}
