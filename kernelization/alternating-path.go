package kernelization

import (
	"github.com/deckarep/golang-set"
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

// TODO: Refactor to use shortest-path.
func forAllCoordPairsInPath(path *graph.IntStack, fn func(int, int, int, int, chan<- bool)) {
	done := make(chan bool, 1)

	iter := path.Iter()
	prevFrom, prevTo := <-iter, <-iter
	n := path.Size()

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

func isAlternatingPathWithMatching(self *graph.Graph, path *graph.IntStack, matching mapset.Set) (result bool) {
	// This version of the method compares by pointers!
	result = true
	forAllCoordPairsInPath(path, func(prevFrom, prevTo, curFrom, curTo int, done chan<- bool) {
		previous := self.GetEdgeByCoordinates(prevFrom, prevTo)
		// P is an alternating path if:
		//  - P is a path in G
		//  - for each subsequent pair of edges from P one of them
		//    belongs to M and the other one does not.
		current := self.GetEdgeByCoordinates(curFrom, curTo)
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

func reachableFromWithMatching(toReach graph.Vertices, reachFrom mapset.Set, netFlow *NetworkFlow, matching mapset.Set) mapset.Set {
	result := mapset.NewThreadUnsafeSet()
	for from := range reachFrom.Iter() {
		for _, to := range toReach {
			if result.Contains(to) {
				continue
			}

			if isReachableWithMatchingThroughAlternatingPath(from.(graph.Vertex), to, netFlow, matching) {
				result.Add(to)
			}
		}
	}

	return result
}

func isReachableWithMatchingThroughAlternatingPath(from, to graph.Vertex, netFlow *NetworkFlow, matching mapset.Set) bool {
	exists, path, _ := shortestPath(netFlow.net, from, to)
	if exists {
		if path.Size() <= 2 {
			// TODO: I'm not sure whether in case where only a single edge is in the path,
			// the mere fact of it existing is enough - does it have to belong/not belong to M?
			return true
		}

		return isAlternatingPathWithMatching(netFlow.graph, path, matching)
	}

	return false
}
