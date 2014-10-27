package vc

import (
	"fmt"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
)

func benchNaive(b *testing.B, no int) {
	if !written {
		writeExamples()
	}

	g := graph.ScanGraph(fmt.Sprintf(examplePathFormat, no))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NaiveVC(g, 10000000)
	}
}

func BenchmarkNaivePessimistic1(b *testing.B) {
	benchNaive(b, 1)
}

func BenchmarkNaivePessimistic2(b *testing.B) {
	benchNaive(b, 2)
}

func BenchmarkNaivePessimistic3(b *testing.B) {
	benchNaive(b, 3)
}

func BenchmarkNaivePessimistic5(b *testing.B) {
	benchNaive(b, 5)
}

func BenchmarkNaivePessimistic6(b *testing.B) {
	benchNaive(b, 6)
}

func BenchmarkNaivePessimistic7(b *testing.B) {
	benchNaive(b, 7)
}

func BenchmarkNaivePessimistic8(b *testing.B) {
	benchNaive(b, 8)
}

func BenchmarkNaivePessimistic9(b *testing.B) {
	benchNaive(b, 9)
}

func BenchmarkNaivePessimistic10(b *testing.B) {
	benchNaive(b, 10)
}

func BenchmarkNaivePessimistic11(b *testing.B) {
	benchNaive(b, 11)
}
