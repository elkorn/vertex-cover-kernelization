package main

var k int = 250

// func TestHighDegreeKernelization(b *testing.T) {
// 	g := graph.ScanGraph("../examples/sh2/sh2-3.dim")
// 	pre, _ := preprocessing.Preprocessing(g)
// 	preprocessing.Preprocessing(g)

// 	log.Println(pre)
// 	_, removed := kernelization.KernelizationHighDegree(g, 100)
// 	log.Println(removed)
// 	log.Println("Start.")
// 	vc.BranchAndBound(g, nil, 250-pre-removed)
// 	log.Println("End.")
// }

// func BenchmarkHighDegreeKernelization(b *testing.B) {
// 	g := graph.ScanGraph("../examples/sh2/sh2-3.dim")
// 	pre, _ := preprocessing.Preprocessing(g)
// 	preprocessing.Preprocessing(g)

// 	_, removed := kernelization.KernelizationHighDegree(g, 250-pre)

// 	graphs := make([]*graph.Graph, 1000)
// 	utility.InVerboseContext(func() {
// 		utility.Debug("Kernelized.")
// 		for i := 0; i < 1000; i++ {
// 			graphs[i] = g.Copy()
// 		}

// 		utility.Debug("Copied.")
// 	})
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		vc.BranchAndBound(graphs[i], nil, 250-pre-removed)
// 	}
// }
