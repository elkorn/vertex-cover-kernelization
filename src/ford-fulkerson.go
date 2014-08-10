package graph

func fordFulkerson(nf *NetworkFlow) (Edges, int) {
	result := Edges{}
	totalFlow := 0
	for pathExists, path := shortestPathFromSourceToSink(nf); pathExists; pathExists, path = shortestPathFromSourceToSink(nf) {
		values := path.Values()
		cfp := MAX_INT
		forAllEdgesInPath := func(fn func(int, int)) {
			for i, n := 1, len(values); i < n; i++ {
				from := values[i-1]
				to := values[i]
				fn(from, to)
			}
		}
		// Get capacity, this may be redundant since in my case
		// the capacity is always 1.
		forAllEdgesInPath(func(from, to int) {
			if nf.net.arcs[from][to].residuum() < cfp {
				cfp = nf.net.arcs[from][to].residuum()
			}
		})

		forAllEdgesInPath(func(from, to int) {
			arc := nf.net.arcs[from][to]
			reverseArc := nf.net.arcs[to][from]
			// This is a point that I have not accounted for earlier.
			// The arc in the residual network may or may not be reflected
			// by an edge in the original graph.
			if nil != arc.edge {
				arc.flow = arc.flow + cfp
				totalFlow += arc.flow
				result = append(result, arc.edge)
			} else {
				reverseArc.flow = reverseArc.flow - cfp
			}
		})
	}

	return result, totalFlow
}
