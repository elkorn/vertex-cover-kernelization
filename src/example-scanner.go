package graph

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type header struct {
	name1    string // TODO: find out what these names actually are.
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

func processLine(line string, graph *Graph) {
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

	graph.AddEdge(MkVertex(from), MkVertex(to))
}

func ScanGraph(path string) *Graph {
	file, err := os.Open(path)

	if nil != err {
		log.Fatal(err)
		return nil
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	header := processHeader(scanner)
	graph := MkGraph(header.vertices)
	for scanner.Scan() {
		processLine(scanner.Text(), graph)
	}

	return graph
}
