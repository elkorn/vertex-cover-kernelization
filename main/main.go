package main

import (
	"fmt"
	"time"

	"github.com/elkorn/vertex-cover-kernelization/graph"
)

type measurement struct {
	positional                 int
	name                       string
	time                       time.Duration
	coverFound                 bool
	vertices, edges, coverSize int
}

func measurementHeader() string {
	return fmt.Sprintf("\t%v\t%v\t%v\t%v",
		"|V|",
		"|E|",
		"|C|",
		" T ")
}

func takeMeasurement(name string, g *graph.Graph, action func(*graph.Graph) (bool, int)) (result *measurement) {
	result = &measurement{
		name:     name,
		vertices: g.NVertices(),
		edges:    g.NEdges(),
	}

	result.time = measure(func() {
		result.coverFound, result.coverSize = action(g)
	})

	return
}

func (self *measurement) withPositional(str string) string {
	if self.positional > 0 {
		return fmt.Sprintf("%v.\t%v", self.positional, str)
	}

	return fmt.Sprintf("\t%v", str)
}

func (self *measurement) Str() string {
	return self.withPositional(fmt.Sprintf("%v\t%v\t%v\t%v",
		self.vertices,
		self.edges,
		self.coverSize,
		self.time))
}

func (self *measurement) StrSeconds() string {
	return self.withPositional(fmt.Sprintf("%v\t%v\t%v\t%v",
		self.vertices,
		self.edges,
		self.coverSize,
		self.time.Seconds()))
}

func measure(action func()) time.Duration {
	start := time.Now()
	action()
	return time.Since(start)
}

func main() {
}
