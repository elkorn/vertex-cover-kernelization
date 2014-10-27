package vc

import (
	"fmt"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
)

const examplePathFormat = "../examples/example_%v"

var written = false

func writeExamples() {
	for i := 1; i <= 11; i++ {
		graph.WriteExampleGraph(fmt.Sprintf(examplePathFormat, i), 10*i)
	}

	written = true
}

func benchBnB(b *testing.B, no int) {
	if !written {
		writeExamples()
	}

	g := graph.ScanGraph(fmt.Sprintf(examplePathFormat, no))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		branchAndBound(g)
	}
}

func BenchmarkBnBPessimistic1(b *testing.B) {
	benchBnB(b, 1)
}

func BenchmarkBnBPessimistic2(b *testing.B) {
	benchBnB(b, 2)
}

func BenchmarkBnBPessimistic3(b *testing.B) {
	benchBnB(b, 3)
}

func BenchmarkBnBPessimistic5(b *testing.B) {
	benchBnB(b, 5)
}

func BenchmarkBnBPessimistic6(b *testing.B) {
	benchBnB(b, 6)
}

func BenchmarkBnBPessimistic7(b *testing.B) {
	benchBnB(b, 7)
}

func BenchmarkBnBPessimistic8(b *testing.B) {
	benchBnB(b, 8)
}

func BenchmarkBnBPessimistic9(b *testing.B) {
	benchBnB(b, 9)
}

func BenchmarkBnBPessimistic10(b *testing.B) {
	benchBnB(b, 10)
}

func BenchmarkBnBPessimistic11(b *testing.B) {
	benchBnB(b, 11)
}
