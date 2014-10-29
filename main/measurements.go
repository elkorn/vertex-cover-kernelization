package main

import (
	"fmt"
	"time"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/elkorn/vertex-cover-kernelization/vc"
)

var k int = 250
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
	for i, dataFile := range dataFiles {
		m := takeMeasurement(
			fmt.Sprintf("bnb %v_%v", dataFile.vertices, dataFile.degreeDistribution),
			graph.ScanDot(fmt.Sprintf(dataFile.path)),
			bnb)
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureNaive() {
	setOutputFile(filenameBnB)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		m := takeMeasurement(
			fmt.Sprintf("naive %v_%v", dataFile.vertices, dataFile.degreeDistribution),
			graph.ScanDot(dataFile.path),
			naive)
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationCrownReduction() {
	setOutputFile(filenameCrown)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		fmt.Println(dataFile.path)
		g := graph.ScanDot(dataFile.path)
		reduction := kernelization.ReduceAllCrowns(g, k)
		var m *measurement
		if reduction == -1 {
			m = &measurement{
				name:       fmt.Sprintf("bnb_crown %v_%v", dataFile.vertices, dataFile.degreeDistribution),
				vertices:   g.NVertices(),
				edges:      g.NEdges(),
				time:       time.Since(time.Now()),
				coverFound: false,
				coverSize:  0,
			}
		} else {
			m := takeMeasurement(
				fmt.Sprintf("bnb_crown %v_%v", dataFile.vertices, dataFile.degreeDistribution),
				g,
				naive)
			m.coverSize += reduction
		}

		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationNetworkFlow() {
	setOutputFile(filenameNF)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		reduction := kernelization.KernelizationNetworkFlow(g, k)
		m := takeMeasurement(
			fmt.Sprintf("bnb_nf %v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			naive)
		m.positional = i + 1
		m.coverSize += reduction
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}
