package graph

import (
	"bytes"
	"fmt"
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

func (self *Graph) ToDot(name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	for _, edge := range self.Edges {
		res.Write(tstobn(fmt.Sprintf("%v -- %v;", edge.from, edge.to)))
	}

	res.Write(stob("}"))
	return res
}

func DotToJpg(dot bytes.Buffer) bytes.Buffer {
	var res bytes.Buffer
	cmd := exec.Command("dot", "T")
	cmd.Stdout = &res
	cmd.Stdin = &dot
	cmd.Run()
	return res
}
