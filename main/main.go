package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
)

var currentfname string

func setOutputFile(filename string) {
	currentfname = filename
	ioutil.WriteFile(filename, []byte{}, os.ModeAppend)
}

func writeln(data string) {
	file, err := os.OpenFile(currentfname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	if nil != err {
		log.Fatal(err)
		return
	}

	file.WriteString(fmt.Sprintf("%v\n", data))
	file.Close()
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

var testCases map[string]func() = map[string]func(){
	"MeasureNaive":                                      MeasureNaive,
	"MeasureVCBnb":                                      MeasureBnb,
	"MeasureVCKernelizationCrownReduction":              MeasureKernelizationCrownReduction,
	"MeasureVCKernelizationNetworkFlow":                 MeasureKernelizationNetworkFlow,
	"MeasureVCPreprocessingBnb":                         MeasureBnbPreprocessing,
	"MeasureVCPreprocessingKernelizationCrownReduction": MeasureKernelizationCrownReductionPreprocessing,
	"MeasureVCPreprocessingKernelizationNetworkFlow":    MeasureKernelizationNetworkFlowPreprocessing,
	// "MeasureKernelizationCrownReduction":                MeasureKernelizationCrownReduction,
	// "MeasureKernelizationNetworkFlow":                   MeasureKernelizationNetworkFlow,
	// "MeasurePreprocessingKernelizationCrownReduction":   MeasureKernelizationCrownReduction,
	// "MeasurePreprocessingKernelizationNetworkFlow":      MeasureKernelizationNetworkFlow,
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
			testCase()
		}
	}
}
