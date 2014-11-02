package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strconv"
)

var currentfname string

func setOutputFile(filename string) {
	currentfname = filename
}

func writeln(data string) {
	file, err := os.OpenFile(currentfname, os.O_APPEND|os.O_CREATE, 0666)

	if nil != err {
		log.Fatal(err)
		return
	}

	file.WriteString(fmt.Sprintf("%v\n", data))
	file.Close()
}

func whoami() string {
	pc, _, _, ok := runtime.Caller(0)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	return me.Name()
}

type flags struct {
	runPattern *string
}

type dataFileDescriptor struct {
	path               string
	vertices           int
	degreeDistribution int
}

func defineFlags() (result flags) {
	result = flags{
		runPattern: flag.String("run", ".*", "Regexp for which measurements should be run"),
	}

	flag.Parse()
	return
}

var testCases map[string]func(string) = map[string]func(string){
	"MeasureBnb":                                      MeasureBnb,
	"MeasureBnbPreprocessing":                         MeasureBnbPreprocessing,
	"MeasureNaive":                                    MeasureNaive,
	"MeasureNaivePreprocessing":                       MeasureNaivePreprocessing,
	"MeasureVCCrownReduction":                         MeasureVCCrownReduction,
	"MeasureVCNetworkFlow":                            MeasureVCNetworkFlow,
	"MeasureVCCrownReductionPreprocessing":            MeasureVCCrownReductionPreprocessing,
	"MeasureVCNetworkFlowPreprocessing":               MeasureVCNetworkFlowPreprocessing,
	"MeasureKernelizationCrownReduction":              MeasureKernelizationCrownReduction,
	"MeasureKernelizationNetworkFlow":                 MeasureKernelizationNetworkFlow,
	"MeasureKernelizationCrownReductionPreprocessing": MeasureKernelizationCrownReductionPreprocessing,
	"MeasureKernelizationNetworkFlowPreprocessing":    MeasureKernelizationNetworkFlowPreprocessing,
}

var dataFiles []dataFileDescriptor

func forAllFilesInDir(dir string, match string, action func(os.FileInfo)) {
	infiles, err := ioutil.ReadDir(dir)
	if nil != err {
		panic(err)
	}
	for _, infile := range infiles {
		if regexp.MustCompile(match).MatchString(infile.Name()) {
			action(infile)
		}
	}
}

func listRandomInFiles(dir string) {
	dataFiles = make([]dataFileDescriptor, 0)
	forAllFilesInDir(dir, "\\d+_\\d+\\.dot", func(infile os.FileInfo) {
		input := regexp.MustCompile("\\d+").FindAllStringSubmatch(infile.Name(), 2)
		descriptor := dataFileDescriptor{
			path: path.Join(dir, infile.Name()),
		}

		descriptor.vertices, _ = strconv.Atoi(input[0][0])
		descriptor.degreeDistribution, _ = strconv.Atoi(input[1][0])
		dataFiles = append(dataFiles, descriptor)
	})

	sorter := &fileSorter{
		files: dataFiles,
	}

	sort.Sort(sorter)
}

func listExInFiles(dir string) {
	dataFiles = make([]dataFileDescriptor, 0)
	forAllFilesInDir(dir, "ex_\\d+\\.dot", func(infile os.FileInfo) {
		input := regexp.MustCompile("\\d+").FindAllStringSubmatch(infile.Name(), 1)
		descriptor := dataFileDescriptor{
			path: path.Join(dir, infile.Name()),
		}

		descriptor.vertices, _ = strconv.Atoi(input[0][0])
		dataFiles = append(dataFiles, descriptor)
	})

	sorter := &fileSorter{
		files: dataFiles,
	}

	sort.Sort(sorter)
}

func main() {
	listRandomInFiles("../results")
	currentFlags := defineFlags()
	for key, testCase := range testCases {
		if regexp.MustCompile(*(currentFlags.runPattern)).MatchString(key) {
			testCase(key)
		}
	}
}
