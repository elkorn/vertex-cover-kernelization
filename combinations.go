package main

import "fmt"

var memo map[uint64]uint64 = make(map[uint64]uint64)

func main() {
	fmt.Println("Hello, playground")
	for i := 0; i < 100; i++ {
		fmt.Printf("Combinations for %d elements: %d\n", i, numberOfCombinations(i))
	}
}

func numberOfCombinations(elements int) uint64 {
	var result uint64 = 0
	for i := 1; i <= elements; i++ {
		result += C(uint64(elements), uint64(i))
	}

	return result + 1 // account for empty set.
}

func C(n, m uint64) uint64 {
	return factorial(n) / (factorial(m) * factorial(n-m))
}

func factorial(n uint64) (result uint64) {
	//fmt.Println("n = ", n)
	if memo[n] != 0 {
		return memo[n]
	}

	if n == 0 {
		result = 1
	} else {
		result = n * factorial(n-1)
	}

	memo[n] = result
	return result
}
