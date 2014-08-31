package graph

import "github.com/deckarep/golang-set"

// TODO Refactor to use shortest-path.
func forAllCoordPairsInPath(path []int, fn func(int, int, int, int, chan<- bool)) {
	done := make(chan bool, 1)

	prevFrom, prevTo := path[0], path[1]
	n := len(path)

	for i := 2; i < n; i++ {
		curFrom, curTo := path[i-1], path[i]
		fn(prevFrom, prevTo, curFrom, curTo, done)
		select {
		case <-done:
			return
		default:
		}

		prevFrom, prevTo = curFrom, curTo
	}

}

func (self *Graph) isAlternatingPathWithMatching(path []int, matching mapset.Set) (result bool) {
	// This version of the method compares by pointers!
	result = true
	forAllCoordPairsInPath(path, func(prevFrom, prevTo, curFrom, curTo int, done chan<- bool) {
		previous := self.getEdgeByCoordinates(prevFrom, prevTo)
		// P is an alternating path if:
		//  - P is a path in G
		//  - for each subsequent pair of edges from P one of them
		//    belongs to M and the other one does not.
		current := self.getEdgeByCoordinates(curFrom, curTo)
		containsPrevious := matching.Contains(previous)
		containsCurrent := matching.Contains(current)
		Debug("Previous: %v, contains: %v", previous, containsPrevious)
		Debug("Current: %v, contains: %v", current, containsCurrent)
		if containsPrevious == containsCurrent {
			result = false
			done <- true
		}
	})

	return result
}

func isAlternatingPathWithMatching(path []int, matching mapset.Set) (result bool) {
	// This version of the method compares by values!
	result = true
	forAllCoordPairsInPath(path, func(prevFrom, prevTo, curFrom, curTo int, done chan<- bool) {
		// P is an alternating path if:
		//  - P is a path in G
		//  - for each subsequent pair of edges from P one of them
		//    belongs to M and the other one does not.

		previous := MkEdgeValFromInts(prevFrom, prevTo)
		current := MkEdgeValFromInts(curFrom, curTo)
		containsPrevious := matching.Contains(previous)
		containsCurrent := matching.Contains(current)
		Debug("Previous: %v, contains: %v", previous, containsPrevious)
		Debug("Current: %v, contains: %v", current, containsCurrent)
		if containsPrevious == containsCurrent {
			result = false
			done <- true
		}
	})

	return result
}

func (from Vertex) isReachableWithMatchingThroughAlternatingPath(to Vertex, netFlow *NetworkFlow, matching mapset.Set) bool {
	exists, path, _ := shortestPath(netFlow.net, from, to)
	if len(path) <= 2 {
		// TODO I'm not sure whether in case where only a single edge is in the path,
		// the mere fact of it existing is enough - does it have to belong/not belong to M?
		return exists
	}

	if exists {
		return netFlow.graph.isAlternatingPathWithMatching(path, matching)
	}

	return false
}
