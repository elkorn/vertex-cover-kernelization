package main

import (
	"fmt"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/elkorn/vertex-cover-kernelization/vc"
)

var k int = 250
var rng []int = []int{50, 100, 200, 500}
var n int = len(rng)
var filenameBnB = "./results_bnb"
var filenameNF = "./results_nf"
var filenameCrown = "./results_cr"

func bnb(g *graph.Graph) (bool, int) {
	result := vc.BranchAndBound(g, nil, utility.MAX_INT)
	return result.Cardinality() > 0, result.Cardinality()
}

func naive(g *graph.Graph) (bool, int) {
	found, result := vc.NaiveVC(g, utility.MAX_INT)
	return found, result.Cardinality()
}

func MeasureBnb() {
	setOutputFile(filenameBnB)
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			m := takeMeasurement(
				"bnb",
				graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2)),
				bnb)
			m.positional = calcPositional(i, j)
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}

func MeasureNaive() {
	setOutputFile(filenameBnB)
	writeln(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			m := takeMeasurement(
				"naive",
				graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2)),
				naive)
			m.positional = calcPositional(i, j)
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}

func MeasureKernelizationCrownReduction() {
	setOutputFile(filenameCrown)
	writeln(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			g := graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2))
			reduction := kernelization.ReduceAllCrowns(g, k)
			m := takeMeasurement(
				"bnb_crown",
				g,
				bnb)
			m.coverSize += reduction
			m.positional = calcPositional(i, j)
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}

func MeasureKernelizationNetworkFlow() {
	setOutputFile(filenameNF)
	writeln(measurementHeader())
	for i, val1 := range rng {
		for j, val2 := range rng {
			g := graph.ScanDot(fmt.Sprintf("../results/%v_%v.dot", val1, val2))
			reduction := kernelization.KernelizationNetworkFlow(g, k)
			m := takeMeasurement(
				"bnb_nf",
				g,
				bnb)
			m.positional = calcPositional(i, j)
			m.coverSize += reduction
			writeln(m.Str())
			fmt.Println(m.Str())
		}
	}
}
