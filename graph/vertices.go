package graph

import "github.com/deckarep/golang-set"

type Vertices []Vertex

func (toReach Vertices) reachableFromWithMatching(reachFrom mapset.Set, netFlow *NetworkFlow, matching mapset.Set) mapset.Set {
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
