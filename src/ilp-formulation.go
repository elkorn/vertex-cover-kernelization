package graph

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func objectiveFunction(feasibleSolutions []map[Vertex]int) map[Vertex]int {
	res := make(map[Vertex]int)
	minWeight := maxInt
	for _, solution := range feasibleSolutions {
		totalWeight := 0
		for _, weight := range solution {
			totalWeight = totalWeight + weight
		}

		if totalWeight < minWeight {
			res = solution
			minWeight = totalWeight
		}
	}

	return res
}
