package graphviz

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

type networkVisualizer struct {
	gv *graphVisualizer
}

func edgeToBWithColor(edge *Edge, color string) []byte {
	if color != "" {
		return tstobn(fmt.Sprintf("%v -- %v [color=\"%v\"];", edge.From, edge.To, color))
	}

	return tstobn(fmt.Sprintf("%v -- %v;", edge.From, edge.To))
}

func MkNetworkVisualizer(g *Graph) *networkVisualizer {
	return &networkVisualizer{
		MkGraphVisualizer(g),
	}
}

func (self *networkVisualizer) toDot(net *Net, name string) bytes.Buffer {
	var res bytes.Buffer
	res.Write(stob("graph "))
	res.Write(stob(name))
	res.Write(stobn(" {"))
	arcs := (*net).arcs
	for y := range arcs {
		for x, arc := range arcs[y] {
			if nil == arc || nil == arc.edge {
				continue
			}

			if arc.residuum() == 0 || (net.arcs[x][y] != nil && net.arcs[x][y].residuum() == 0) {
				res.Write(edgeToBWithColor(arc.edge, "red"))
			} else {
				res.Write(edgeToB(arc.edge))
			}
		}
	}

	res.Write(stob("}"))
	return res
}

func (self *networkVisualizer) mkImage(net *Net, name string) bytes.Buffer {
	return self.gv.dotToImage(self.toDot(net, name))
}

func (self *networkVisualizer) MkImage(net *Net, name string) error {
	file, err := os.Create(fmt.Sprintf("%v.svg", name))
	if nil != err {
		return err
	}

	defer file.Close()
	buf := self.mkImage(net, name)
	_, err = file.Write(buf.Bytes())
	return err
}

func (self *networkVisualizer) Display(net *Net) {
	randname := fmt.Sprintf("%v", rand.Int63())
	filename := fmt.Sprintf("%v.svg", randname)
	cmd := exec.Command("feh", filename)
	err := self.MkImage(net, randname)
	if nil != err {
		log.Fatal(err)
	}

	defer os.Remove(filename)

	err = cmd.Run()
	if nil != err {
		log.Fatal(err)
	}
}
