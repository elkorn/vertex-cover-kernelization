package graph

var index int

func permutations(input []int) [][]int {
	index = 0
	res := make([][]int, factorial(len(input)))
	permutationIter(make([]int, 0), input, res)
	return res
}

func without(index int, arr []int) []int {
	c := make([]int, 0, len(arr)-1)
	return append(append(c, arr[:index]...), arr[index+1:]...)
}

func permutationIter(prefix []int, all []int, result [][]int) {
	Debug("all: %v, prefix: %v", all, prefix)
	n := len(all)
	if n == 0 {
		result[index] = prefix
		index = index + 1
		Debug("Added %v.", prefix)
	} else {
		for i := 0; i < n; i++ {
			Debug("%v", without(i, all))
			permutationIter(append(prefix, all[i]), without(i, all), result)
		}
	}
}
