package graph

import "github.com/deckarep/golang-set"

func isAlternatingPathWithMatching(path []int, matching mapset.Set) bool {
	previous := MkEdgeFromInts(path[0], path[1])
	// P is an alternating path if:
	//  - P is a path in G
	//  - for each subsequent pair of edges from P one of them
	//    belongs to M and the other one does not.

	for i := 2; i < len(path); i++ {
		current := MkEdgeFromInts(path[i-1], path[i])
		if matching.Contains(previous) == matching.Contains(current) {
			return false
		}

		previous = current
	}

	return true
}

func (from Vertex) isReachableWithMatchingThroughAlternatingPath(to Vertex, net Net, matching mapset.Set) bool {
	exists, path, _ := shortestPath(net, from, to)
	if len(path) <= 2 {
		// TODO I'm not sure whether in case where only a single edge is in the path,
		// the mere fact of it existing is enough - does it have to belong/not belong to M?
		return exists
	}

	if exists {
		return isAlternatingPathWithMatching(path, matching)
	}

	return false
}
