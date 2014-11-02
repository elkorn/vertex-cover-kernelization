package main

import (
	"fmt"
	"time"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/kernelization"
	"github.com/elkorn/vertex-cover-kernelization/preprocessing"
	"github.com/elkorn/vertex-cover-kernelization/utility"
	"github.com/elkorn/vertex-cover-kernelization/vc"
)

var k int = utility.MAX_INT

func bnb(g *graph.Graph) (bool, int) {
	result := vc.BranchAndBound(g, nil, k)
	return result.Cardinality() > 0, result.Cardinality()
}

func naive(g *graph.Graph) (bool, int) {
	found, result := vc.NaiveVC(g, k)
	return found, result.Cardinality()
}

func MeasureBnb(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func MeasureBnbPreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func MeasureNaive(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func MeasureNaivePreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	fmt.Println(measurementHeader())
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		r1, _ := preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("naive:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			naive)
		m.degreeDistribution = dataFile.degreeDistribution
		m.coverSize += r1
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureVCCrownReduction(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2, r3 int
		g := graph.ScanDot(dataFile.path)
		r2 = kernelization.ReduceAllCrowns(g, k)

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

func MeasureVCNetworkFlow(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func MeasureVCCrownReductionPreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func WriteGraphSizes(whoami string) {
	fmt.Println(whoami)
	for _, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		fmt.Println(g.NVertices(), g.NEdges())
	}
}

func MeasureVCNetworkFlowPreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
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

func MeasureKernelizationCrownReduction(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		var m *measurement
		m = takeMeasurement(
			fmt.Sprintf("k_crown:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			func(g *graph.Graph) (bool, int) {
				return true, kernelization.ReduceAllCrowns(g, k)
			})
		m.degreeDistribution = dataFile.degreeDistribution
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationNetworkFlow(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		m := takeMeasurement(
			fmt.Sprintf("k_nf:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			func(g *graph.Graph) (bool, int) {
				return true, kernelization.KernelizationNetworkFlow(g, k)
			})
		m.positional = i + 1
		m.degreeDistribution = dataFile.degreeDistribution
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationCrownReductionPreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	fmt.Println(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2 int
		g := graph.ScanDot(dataFile.path)
		r1, _ = preprocessing.Preprocessing(g)
		var m *measurement
		m = takeMeasurement(
			fmt.Sprintf("k_crown_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			func(g *graph.Graph) (bool, int) {
				r2 = kernelization.ReduceAllCrowns(g, k)
				return true, r1 + r2
			})
		m.degreeDistribution = dataFile.degreeDistribution
		m.positional = i + 1
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasureKernelizationNetworkFlowPreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		var r1, r2 int
		g := graph.ScanDot(dataFile.path)
		r1, _ = preprocessing.Preprocessing(g)
		m := takeMeasurement(
			fmt.Sprintf("k_nf_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			func(g *graph.Graph) (bool, int) {
				r2 = kernelization.KernelizationNetworkFlow(g, k)
				return true, r1 + r2
			})
		m.positional = i + 1
		m.degreeDistribution = dataFile.degreeDistribution
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}

func MeasurePreprocessing(whoami string) {
	setOutputFile(whoami)
	fmt.Println(whoami)
	writeln(measurementHeader())
	for i, dataFile := range dataFiles {
		g := graph.ScanDot(dataFile.path)
		m := takeMeasurement(
			fmt.Sprintf("k_nf_p:%v_%v", dataFile.vertices, dataFile.degreeDistribution),
			g,
			func(g *graph.Graph) (bool, int) {
				r1, _ := preprocessing.Preprocessing(g)
				return true, r1
			})
		m.positional = i + 1
		m.degreeDistribution = dataFile.degreeDistribution
		writeln(m.Str())
		fmt.Println(m.Str())
	}
}
