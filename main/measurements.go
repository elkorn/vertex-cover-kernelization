package main

import (
	"fmt"
	"time"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/preprocessing"
	"github.com/elkorn/vertex-cover-kernelization/vc"
)

var k int = 805
var filenameBnB = "./results_bnb"
var filenameNaive = "./results_naive"
var filenameNF = "./results_nf"
var filenameCrown = "./results_cr"

func bnb(g *graph.Graph) (bool, int) {
	result := vc.BranchAndBound(g, nil, k)
	return result.Cardinality() > 0, result.Cardinality()
}

func naive(g *graph.Graph) (bool, int) {
	found, result := vc.NaiveVC(g, k)
	return found, result.Cardinality()
}

func MeasureBnb() {
	setOutputFile(filenameBnB)
	fmt.Println("MeasureBnb")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1 int
		g := graph.ScanDot(fmt.Sprintf(dataFile.path))
		// r1, _ = preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("bnb:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			bnb)
		m.positional = i + 1
		m.coverSize += r1
		m.degreeDistribution = dataFile.degreeDistribution
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureBnbPreprocessing() {
	setOutputFile(filenameBnB)
	fmt.Println("MeasureBnb")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1 int
		g := graph.ScanDot(fmt.Sprintf(dataFile.path))
		r1, _ = preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("bnb:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			bnb)
		m.positional = i + 1
		m.coverSize += r1
		m.degreeDistribution = dataFile.degreeDistribution
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureNaive() {
	setOutputFile(filenameNaive)
	fmt.Println("MeasureNaive")
	fmt.Println(measurementHeader())
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		m := takeMeasurement(
			fmt.Sprintf("naive:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			graph.ScanDot(dataFile.path),
			naive)
		m.degreeDistribution = dataFile.degreeDistribution
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationCrownReduction() {
	setOutputFile(filenameCrown)
	fmt.Println("MeasureKernelizationCrownReduction")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2, r3 int
		g := graph.ScanDot(dataFile.path)
		// r1, _ = preprocessing.Preprocessing(g)
		r2 = kernelization.ReduceAllCrowns(g, k)
		// r3, _ = preprocessing.Preprocessing(g)

		var m *measurement
		if r2 == -1 {
			m = &measurement{
				name:       fmt.Sprintf("bnb_crown:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
				vertices:   g.NVertices(),
				edges:      g.NEdges(),
				time:       time.Since(time.Now()),
				coverFound: false,
				coverSize:  0,
			}
		} else {
			m = takeMeasurement(
				fmt.Sprintf("bnb_crown:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
				g,
				bnb)
			m.coverSize += r1 + r2 + r3
		}

		m.degreeDistribution = dataFile.degreeDistribution
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationNetworkFlow() {
	setOutputFile(filenameNF)
	fmt.Println("MeasureKernelizationNetworkFlow")
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2, r3 int
		g := graph.ScanDot(dataFile.path)
		// r1, _ = preprocessing.Preprocessing(g)
		r2 = kernelization.KernelizationNetworkFlow(g, k)
		// r3, _ = preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("bnb_nf:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			bnb)
		m.positional = i + 1
		m.degreeDistribution = dataFile.degreeDistribution
		m.coverSize += r1 + r2 + r3
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationCrownReductionPreprocessing() {
	setOutputFile(filenameCrown)
	fmt.Println("MeasureKernelizationCrownReduction")
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2, r3 int
		g := graph.ScanDot(dataFile.path)
		r1, _ = preprocessing.Preprocessing(g)
		r2 = kernelization.ReduceAllCrowns(g, k)
		// r3, _ = preprocessing.Preprocessing(g)

		var m *measurement
		if r2 == -1 {
			m = &measurement{
				name:       fmt.Sprintf("bnb_crown_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
				vertices:   g.NVertices(),
				edges:      g.NEdges(),
				time:       time.Since(time.Now()),
				coverFound: false,
				coverSize:  0,
			}
		} else {
			m = takeMeasurement(
				fmt.Sprintf("bnb_crown_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
				g,
				bnb)
			m.coverSize += r1 + r2 + r3
		}

		m.degreeDistribution = dataFile.degreeDistribution
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationNetworkFlowPreprocessing() {
	setOutputFile(filenameNF)
	fmt.Println("MeasureKernelizationNetworkFlow")
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2, r3 int
		g := graph.ScanDot(dataFile.path)
		r1, _ = preprocessing.Preprocessing(g)
		r2 = kernelization.KernelizationNetworkFlow(g, k)
		// r3, _ = preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("bnb_nf_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			bnb)
		m.positional = i + 1
		m.degreeDistribution = dataFile.degreeDistribution
		m.coverSize += r1 + r2 + r3
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}
