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
	self.buf.WriteByte('\t')
	self.buf.WriteString(strconv.Itoa(from))
	self.buf.WriteString(" -- ")
	self.buf.WriteString(strconv.Itoa(to))
	self.buf.WriteString(";\n")
}

func MkExampleGraph(vertices int) []byte {
	var buf bytes.Buffer
	ew := mkEdgeWriter(&buf)
	buf.WriteString("a b ")
	buf.WriteString(strconv.Itoa(vertices))
	buf.WriteByte(' ')
	buf.WriteString(strconv.Itoa(vertices*vertices - vertices))
	buf.WriteString("{\n")
	for from := 1; from <= vertices; from++ {
		for to := 1; to <= vertices; to++ {
			if to == from {
				continue
			}

			ew.writeEdge(from, to)
		}
	}

	buf.WriteString("}")
	return buf.Bytes()
}

func WriteExampleGraph(filename string, vertices int) {
	ioutil.WriteFile(filename, MkExampleGraph(vertices), 0777)
}
