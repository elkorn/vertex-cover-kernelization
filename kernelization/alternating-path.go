package kernelization

import "github.com/deckarep/golang-set"

// TODO: Refactor to use shortest-path.
func forAllCoordPairsInPath(path *IntStack, fn func(int, int, int, int, chan<- bool)) {
	done := make(chan bool, 1)

	iter := path.Iter()
	prevFrom, prevTo := <-iter, <-iter
	n := path.s.count

	for i := 2; i < n; i++ {
		curFrom, curTo := <-iter, <-iter
		fn(prevFrom, prevTo, curFrom, curTo, done)
		select {
		case <-done:
			return
		default:
		}

		prevFrom, prevTo = curFrom, curTo
	}

}

func (self *Graph) isAlternatingPathWithMatching(path *IntStack, matching mapset.Set) (result bool) {
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
		utility.Debug("Previous: %v, contains: %v", previous, containsPrevious)
		utility.Debug("Current: %v, contains: %v", current, containsCurrent)
		if containsPrevious == containsCurrent {
			result = false
			done <- true
		}
	})

	return result
}

func isAlternatingPathWithMatching(path []*Edge, matching mapset.Set) bool {
	// P is an alternating path if:
	//  - P is a path in G
	//  - for each subsequent pair of edges from P one of them
	//    belongs to M and the other one does not.

	containsPrevious := matching.Contains(path[0])
	for i, n := 1, len(path); i < n; i++ {
		containsCurrent := matching.Contains(path[i])
		if containsPrevious == containsCurrent {
			return false
		}

		containsPrevious = containsCurrent
	}

	return true
	// forAllCoordPairsInPath(path, func(prevFrom, prevTo, curFrom, curTo int, done chan<- bool) {

	// 	previous := MkEdgeValFromInts(prevFrom, prevTo)
	// 	current := MkEdgeValFromInts(curFrom, curTo)
	// 	containsPrevious := matching.Contains(previous)
	// 	containsCurrent := matching.Contains(current)
	// 	utility.Debug("Previous: %v, contains: %v", previous, containsPrevious)
	// 	utility.Debug("Current: %v, contains: %v", current, containsCurrent)
	// 	if containsPrevious == containsCurrent {
	// 		result = false
	// 		done <- true
	// 	}
	// })

	// return result
}

func reachableFromWithMatching(toReach Vertices, reachFrom mapset.Set, netFlow *NetworkFlow, matching mapset.Set) mapset.Set {
	result := mapset.NewSet()
	for from := range reachFrom.Iter() {
		for _, to := range toReach {
			if result.Contains(to) {
				continue
			}

			if (from.(Vertex)).isReachableWithMatchingThroughAlternatingPath(to, netFlow, matching) {
				result.Add(to)
			}
		}
	}

	return result
}

func (from Vertex) isReachableWithMatchingThroughAlternatingPath(to Vertex, netFlow *NetworkFlow, matching mapset.Set) bool {
	exists, path, _ := shortestPath(netFlow.net, from, to)
	if exists {
		if path.s.count <= 2 {
			// TODO: I'm not sure whether in case where only a single edge is in the path,
			// the mere fact of it existing is enough - does it have to belong/not belong to M?
			return true
		}

		return netFlow.graph.isAlternatingPathWithMatching(path, matching)
	}

	return false
}
