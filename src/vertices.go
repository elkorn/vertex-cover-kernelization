package graph

import "github.com/deckarep/golang-set"

type Vertices []Vertex

func (toReach Vertices) reachableFromWithMatching(reachFrom mapset.Set, net Net, matching mapset.Set) mapset.Set {
	result := mapset.NewSet()
	for from := range reachFrom.Iter() {
		for _, to := range toReach {
			if result.Contains(to) {
				continue
			}

			if (from.(Vertex)).isReachableWithMatchingThroughAlternatingPath(to, net, matching) {
				result.Add(to)
			}
		}
	}

	return result
}
