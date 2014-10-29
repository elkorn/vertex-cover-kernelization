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
	"strings"
)

var currentfname string

func setOutputFile(filename string) {
	currentfname = filename
	ioutil.WriteFile(filename, []byte{}, os.ModeAppend)
}

func writeln(data string) {
	file, err := os.OpenFile(currentfname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

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
	// "MeasureNaive":                       MeasureNaive,
	"MeasureBnb":                         MeasureBnb,
	"MeasureKernelizationCrownReduction": MeasureKernelizationCrownReduction,
	"MeasureKernelizationNetworkFlow":    MeasureKernelizationNetworkFlow,
}

var dataFiles []dataFileDescriptor

func listInFiles(dir string) {
	infiles, err := ioutil.ReadDir(dir)
	if nil != err {
		panic(err)
	}

	dataFiles = make([]dataFileDescriptor, 0, len(infiles))
	for _, infile := range infiles {
		if strings.HasSuffix(infile.Name(), ".dot") {
			input := regexp.MustCompile("\\d+").FindAllStringSubmatch(infile.Name(), 2)
			descriptor := dataFileDescriptor{
				path: path.Join(dir, infile.Name()),
			}

			descriptor.vertices, _ = strconv.Atoi(input[0][0])
			descriptor.degreeDistribution, _ = strconv.Atoi(input[1][0])
			dataFiles = append(dataFiles, descriptor)
		}
	}

	sorter := &fileSorter{
		files: dataFiles,
	}

	sort.Sort(sorter)
}

func main() {
	listInFiles("../results")
	currentFlags := defineFlags()
	for key, testCase := range testCases {
		if regexp.MustCompile(*(currentFlags.runPattern)).MatchString(key) {
			testCase()
		}
	}
}
