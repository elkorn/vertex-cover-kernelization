package utility

import (
	"fmt"
	"log"
	"math"

	"github.com/deckarep/golang-set"
)

const MAX_UINT = ^uint(0)
const MAX_INT = int(MAX_UINT >> 1)

func mkBoolMatrix(n, cap int) [][]bool {
	result := make([][]bool, n, cap)
	for i := range result {
		result[i] = make([]bool, n, cap)
	}

	return result
}

func PrintSet(set mapset.Set) {
	for s := range set.Iter() {
		Debug("%v", s)
	}
}

func Debug(format string, args ...interface{}) {
	if options.Verbose {
		log.Print(fmt.Sprintf(format, args...))
	}
}

func InVerboseContext(fn func()) {
	SetOptions(Options{Verbose: true})
	fn()
	SetOptions(Options{Verbose: false})
}

func rng(args ...int) []int {
	c := make([]int, len(args))
	copy(c, args)
	return c
}

func IntAbs(val int) int {
	return int(math.Abs(float64(val)))
}
