package graph

func fordFulkerson(nf *NetworkFlow) (Edges, int) {
	result := make(Edges, 0, nf.graph.NEdges())
	totalFlow := 0
	for pathExists, path, _ := shortestPathFromSourceToSink(nf); pathExists; pathExists, path, _ = shortestPathFromSourceToSink(nf) {
		bottleneckCapacity := MAX_INT
		forAllEdgesInPath := func(fn func(int, int)) {
			for i, n := 1, len(path); i < n; i++ {
				from := path[i-1]
				to := path[i]
				fn(from, to)
			}
		}

		// Get bottleneck capacity.
		// This may be redundant since in my case the capacity is always 1.
		forAllEdgesInPath(func(from, to int) {
			if nf.net.arcs[from][to].residuum() < bottleneckCapacity {
				bottleneckCapacity = nf.net.arcs[from][to].residuum()
			}
		})

		forAllEdgesInPath(func(from, to int) {
			arc := nf.net.arcs[from][to]
			reverseArc := nf.net.arcs[to][from]
			// The arc in the residual network may or may not be reflected
			// by an edge in the original graph.
			if nil != arc.edge {
				arc.flow = arc.flow + bottleneckCapacity
				totalFlow += arc.flow
				result = append(result, arc.edge)
			} else {
				reverseArc.flow = reverseArc.flow - bottleneckCapacity
			}
		})
	}

	return result, totalFlow
}
