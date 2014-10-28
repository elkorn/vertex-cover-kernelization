package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/elkorn/vertex-cover-kernelization/vc"
)

var k int = 250
var rng []int = []int{50, 100, 200, 500}
var n int = len(rng)
var filename = "./results"

func writeln(data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)

	if nil != err {
		log.Fatal(err)
		return
	}

	file.WriteString(fmt.Sprintf("%v\n", data))
	file.Close()
}

func calcPos(i, j int) int {
	return (i * (n)) + j + 1
}

func bnb(g *graph.Graph) (bool, int) {
	result := vc.BranchAndBound(g, nil, utility.MAX_INT)
	return result.Cardinality() > 0, result.Cardinality()
}

func TestBnbGenerated(t *testing.T) {
	writeln("TestBnbGenerated")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			m := takeMeasurement(
				"bnb",
				graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2)),
				bnb)
			m.positional = calcPos(i, j)
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}

func TestKernelizationCrownReduction(t *testing.T) {
	writeln("TestKernelizationCrownReduction")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			g := graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2))
			_, h := kernelization.ReduceCrown(g, nil, k)
			m := takeMeasurement(
				"bnb_crown",
				g,
				bnb)
			m.coverSize += h.Cardinality()
			m.positional = calcPos(i, j)
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}

func TestKernelizationNetworkFlow(t *testing.T) {
	writeln("TestKernelizationNetworkFlow")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			g := graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2))
			reduction := kernelization.KernelizationNetworkFlow(g, k)
			m := takeMeasurement(
				"bnb_nf",
				g,
				bnb)
			m.positional = calcPos(i, j)
			m.coverSize += reduction
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}
