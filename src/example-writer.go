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
	var buf bytes.Buffer
	ew := mkEdgeWriter(&buf)
	buf.WriteString("a b ")
	buf.WriteString(strconv.Itoa(vertices))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(vertices*vertices - vertices))
	buf.WriteByte('\n')
	for from := 0; from < vertices; from++ {
		for to := 0; to < vertices; to++ {
			if to == from {
				continue
			}

			ew.writeEdge(from, to)
		}
	}

	return buf.Bytes()
}

func WriteExampleGraph(filename string, vertices int) {
	ioutil.WriteFile(filename, MkExampleGraph(vertices), 0777)
}
