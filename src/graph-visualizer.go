package graph

import (
	"bytes"
	"fmt"
)

func stob(str string) []byte {
	return []byte(str)
}

func stobn(str string) []byte {
	return append(stob(str), '\n')
}

func ttstobn(str string) []byte {
	res := new(bytes.Buffer)
	res.Write(stob("\t\t"))
	res.Write(stobn(str))
	return res.Bytes()
}

func (self *Graph) ToDot(name string) string {
	res := new(bytes.Buffer)
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	for _, edge := range self.Edges {
		res.Write(ttstobn(fmt.Sprintf("%v -- %v;", edge.from, edge.to)))
	}

	res.Write(stob("}"))
	return res.String()
}
