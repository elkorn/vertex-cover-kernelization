package kernelization

import (
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

func fordFulkerson(nf *NetworkFlow) (graph.Edges, int) {
	result := make(graph.Edges, 0, nf.graph.NEdges())
	totalFlow := 0
	for pathExists, path, _ := shortestPathFromSourceToSink(nf); pathExists; pathExists, path, _ = shortestPathFromSourceToSink(nf) {
		bottleneckCapacity := utility.MAX_INT
		forAllEdgesInPath := func(fn func(int, int)) {
			iter := path.Iter()
			from, ok := <-iter
			for ok {
				to, ok := <-iter
				utility.Debug("%v -> %v (ok: %v)", from, to, ok)
				if !ok {
					break
				}
				fn(from, to)
				from = to
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
