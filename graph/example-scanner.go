package graph

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type header struct {
	name1    string
	name2    string
	vertices int
	edges    int
}

func processHeader(scanner *bufio.Scanner) header {
	if !scanner.Scan() {
		panic(errors.New("Nothing to scan!"))
	}

	line := scanner.Text()
	// Format: p sh 839 5860
	segments := strings.Split(line, " ")
	vertices, err := strconv.Atoi(segments[2])
	if nil != err {
		log.Fatal(err)
		return header{}
	}

	edges, err := strconv.Atoi(segments[3])
	if nil != err {
		log.Fatal(err)
		return header{}
	}

	return header{
		name1:    segments[0],
		name2:    segments[1],
		vertices: vertices,
		edges:    edges,
	}
}

func processLine(line string, g *Graph) {
	// Format: e 0 98
	segments := strings.Split(line, " ")
	from, err := strconv.Atoi(segments[1])
	if nil != err {
		log.Fatal(err)
		return
	}

	to, err := strconv.Atoi(segments[2])
	if nil != err {
		log.Fatal(err)
		return
	}

	g.AddEdge(MkVertex(from), MkVertex(to))
}

func withFile(path string, action func(scanner *bufio.Scanner)) {
	file, err := os.Open(path)

	if nil != err {
		log.Fatal(err)
		return
	}

	defer file.Close()
	action(bufio.NewScanner(file))
}

func ScanGraph(path string) (graph *Graph) {
	withFile(path, func(scanner *bufio.Scanner) {
		header := processHeader(scanner)
		graph = MkGraph(header.vertices)
		for scanner.Scan() {
			processLine(scanner.Text(), graph)
		}
	})

	return
}

type dotScanner struct {
	g         *Graph
	maxVertex int
	edges     Edges
}

func mkDotScanner() *dotScanner {
	return &dotScanner{
		edges:     make(Edges, 0, 1000),
		maxVertex: -1,
	}
}

func (self *dotScanner) checkMaxVertex(val int) {
	if val > self.maxVertex {
		self.maxVertex = val
	}
}

func (self *dotScanner) processLine(line string) *dotScanner {
	result := regexp.MustCompile("\\d+").FindAllStringSubmatch(line, 2)
	getVal := func(input []string) int {
		result, err := strconv.Atoi(input[0])
		if nil != err {
			panic(err)
		}

		self.checkMaxVertex(result)
		return result
	}

	if len(result) == 2 {
		edge := MkEdgeFromInts(getVal(result[0]), getVal(result[1]))
		self.edges = append(self.edges, edge)
	}

	return self
}

func (self *dotScanner) ProcessFile(path string) *dotScanner {
	withFile(path, func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			self.processLine(scanner.Text())
		}
	})

	return self
}

func (self *dotScanner) MakeGraph() *Graph {
	self.g = MkGraph(self.maxVertex + 1)
	for _, edge := range self.edges {
		self.g.AddEdge(edge.From, edge.To)
	}

	return self.g
}

func ScanDot(path string) *Graph {
	return mkDotScanner().ProcessFile(path).MakeGraph()
}
