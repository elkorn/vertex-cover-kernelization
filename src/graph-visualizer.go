package graph

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

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

func toDot(g *Graph, name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	for _, edge := range g.Edges {
		res.Write(tstobn(fmt.Sprintf("%v -- %v;", edge.from, edge.to)))
	}

	res.Write(stob("}"))
	return res
}

func dotToJpg(dot bytes.Buffer) bytes.Buffer {
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

func MkJpg(g *Graph, name string) error {
	file, err := os.Create(fmt.Sprintf("%v.jpg", name))
	if nil != err {
		return err
	}

	defer file.Close()
	buf := dotToJpg(toDot(g, name))
	_, err = file.Write(buf.Bytes())
	return err
}
