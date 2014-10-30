package graph

import (
	"bytes"
	"io/ioutil"
	"strconv"
)

var bs = make([]byte, 4)

type edgeWriter struct {
	buf *bytes.Buffer
}

func mkEdgeWriter(buf *bytes.Buffer) *edgeWriter {
	return &edgeWriter{
		buf: buf,
	}
}

func (self *edgeWriter) writeEdge(from, to int) {
	self.buf.WriteString("e ")
	self.buf.WriteString(strconv.Itoa(from))
	self.buf.WriteByte(' ')
	self.buf.WriteString(strconv.Itoa(to))
	self.buf.WriteString("\n")
}

func MkExampleGraph(vertices int) []byte {
	written := make([][]bool, vertices)
	for i := range written {
		written[i] = make([]bool, vertices)
	}

	var buf bytes.Buffer
	ew := mkEdgeWriter(&buf)
	buf.WriteString("a b ")
	buf.WriteString(strconv.Itoa(vertices))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa((vertices*vertices - vertices) / 2))
	buf.WriteByte('\n')
	for from := 0; from < vertices; from++ {
		for to := 0; to < vertices; to++ {
			if to == from || written[from][to] || written[to][from] {
				continue
			}

			ew.writeEdge(from, to)
			written[from][to] = true
		}
	}

	return buf.Bytes()
}

func WriteExampleGraph(filename string, vertices int) {
	ioutil.WriteFile(filename, MkExampleGraph(vertices), 0777)
}
