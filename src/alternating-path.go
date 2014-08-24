package graph

import "github.com/deckarep/golang-set"

func (self *Graph) isAlternatingPathWithMatching(path []int, matching mapset.Set) bool {
	// This version of the method compares by pointers!
	previous := self.getEdgeByCoordinates(path[0], path[1])
	// P is an alternating path if:
	//  - P is a path in G
	//  - for each subsequent pair of edges from P one of them
	//    belongs to M and the other one does not.

	for i := 2; i < len(path); i++ {
		current := self.getEdgeByCoordinates(path[i-1], path[i])
		containsPrevious := matching.Contains(previous)
		containsCurrent := matching.Contains(current)
		Debug("Previous: %v, contains: %v", previous, containsPrevious)
		Debug("Current: %v, contains: %v", current, containsCurrent)
		if containsPrevious == containsCurrent {
			return false
		}

		// TODO this may cause havoc - are pointers getting switched in sources?
		previous = current
	}

	return true
}

func isAlternatingPathWithMatching(path []int, matching mapset.Set) bool {
	// This version of the method compares by values!
	previous := MkEdgeValFromInts(path[0], path[1])
	// P is an alternating path if:
	//  - P is a path in G
	//  - for each subsequent pair of edges from P one of them
	//    belongs to M and the other one does not.

	for i := 2; i < len(path); i++ {
		current := MkEdgeValFromInts(path[i-1], path[i])
		containsPrevious := matching.Contains(previous)
		containsCurrent := matching.Contains(current)
		Debug("Previous: %v, contains: %v", previous, containsPrevious)
		Debug("Current: %v, contains: %v", current, containsCurrent)
		if containsPrevious == containsCurrent {
			return false
		}

		previous = current
	}

	return true
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
