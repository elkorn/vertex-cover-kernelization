package graph

import (
	"container/list"
	"fmt"
	"testing"
)

// func benchScannedGraph(b *testing.B, filename string, fn func(g *Graph)) {
// 	original := ScanGraph(filename)
// 	gs := make([]*Graph, b.N)
// 	for i := 0; i < b.N; i++ {
// 		gs[i] = original.Copy()
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		fn(gs[i])
// 	}
// }

// func benchExample(b *testing.B, exampleNum, k int) {
// 	benchScannedGraph(b, fmt.Sprintf("example_%v", exampleNum), func(g *Graph) {
// 		networkFlowKernelization(g, k)
// 	})
// }

// func BenchmarkNetworkFlow1(b *testing.B) {
// 	benchExample(b, 1, 3)
// }

// func BenchmarkNetworkFlow2(b *testing.B) {
// 	benchExample(b, 2, 3)
// }

// func BenchmarkNetworkFlow3(b *testing.B) {
// 	benchExample(b, 3, 3)
// }

// func BenchmarkNetworkFlow4(b *testing.B) {
// 	benchExample(b, 4, 3)
// }

// func BenchmarkNetworkFlow5(b *testing.B) {
// 	benchExample(b, 5, 3)
// }

// func BenchmarkNetworkFlow6(b *testing.B) {
// 	benchExample(b, 6, 3)
// }

// func BenchmarkNetworkFlow7(b *testing.B) {
// 	benchExample(b, 7, 3)
// }

// func BenchmarkNetworkFlow8(b *testing.B) {
// 	benchExample(b, 8, 3)
// }

// func BenchmarkNetworkFlow9(b *testing.B) {
// 	benchExample(b, 9, 3)
// }

// func BenchmarkNetworkFlow10(b *testing.B) {
// 	benchExample(b, 10, 3)
// }

// func BenchmarkNetworkFlow11(b *testing.B) {
// 	benchExample(b, 11, 3)
// }

// func BenchmarkNetworkFlow12(b *testing.B) {
// 	benchExample(b, 12, 3)
// }

// func BenchmarkNetworkFlow13(b *testing.B) {
// 	benchExample(b, 13, 3)
// }

// func BenchmarkNetworkFlow14(b *testing.B) {
// 	benchExample(b, 14, 3)
// }

// func BenchmarkNetworkFlow15(b *testing.B) {
// 	benchExample(b, 15, 3)
// }

func BenchmarkPrependList(b *testing.B) {
	list := list.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			list.PushFront(j)
		}

		get(list, 9999)
	}
}

func BenchmarkPrependSlice(b *testing.B) {
	slice := make([]Vertex, 0, 10000)
	var p Vertex
	Debug("%v", p)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			slice = prepend(slice, Vertex(j))
		}

		p = slice[9999]
	}
}

func BenchmarkPrependSliceWithTemplate(b *testing.B) {
	slice := make([]Vertex, 0, 10000)
	template := make([]Vertex, 1)
	var p Vertex
	Debug("%v", p)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			template[0] = Vertex(j)
			slice = append(template, slice...)
		}

		p = slice[9999]
	}
}

func BenchmarkBSTConflictResolverStraightPetersenLT(b *testing.B) {
	g := mkPetersenGraph()
	CONFLICT_RESOLVER = func(g *Graph, d1, d2 int) bool {
		return d1 < d2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		branchAndBound(g)
	}
}

func BenchmarkBSTConflictResolverStraightPetersenGT(b *testing.B) {
	g := mkPetersenGraph()
	CONFLICT_RESOLVER = func(g *Graph, d1, d2 int) bool {
		return d1 > d2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		branchAndBound(g)
	}
}

func BenchmarkBSTConflictResolverStraightPetersenLTEQ(b *testing.B) {
	g := mkPetersenGraph()
	CONFLICT_RESOLVER = func(g *Graph, d1, d2 int) bool {
		return d1 <= d2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		branchAndBound(g)
	}
}

func BenchmarkBSTConflictResolverStraightPetersenGTEQ(b *testing.B) {
	g := mkPetersenGraph()
	CONFLICT_RESOLVER = func(g *Graph, d1, d2 int) bool {
		return d1 >= d2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		branchAndBound(g)
	}
}

func benchBnB(b *testing.B, no int) {
	g := ScanGraph(fmt.Sprintf("example_%v", no))
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
