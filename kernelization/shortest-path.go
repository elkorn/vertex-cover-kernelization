package kernelization

import (
	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/elkorn/vertex-cover-kernelization/utility"
)

func shortestPathFromSourceToSink(nf *NetworkFlow) (bool, *graph.IntStack, []int) {
	return shortestPath(nf.net, nf.source, nf.sink)
}

// TODO: migrate to the new logic.
func shortestPath(net Net, from, to graph.Vertex) (bool, *graph.IntStack, []int) {
	n := len(net.arcs)
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	distance := make([]int, n)
	queue := graph.MkIntQueue(n)
	mark := func(v graph.Vertex) {
		vi := v.ToInt()
		marked[vi] = true
		queue.Push(vi)
	}
	pathTo := func(v graph.Vertex) *graph.IntStack {
		vi := v.ToInt()
		si := from.ToInt()

		if !marked[vi] {
			return nil
		}

		pathLength := distance[v.ToInt()]
		path := graph.MkIntStack(pathLength + 1)

		for x := vi; x != si; x = edgeTo[x] {
			path.Push(x)
		}

		path.Push(si)
		return path
	}

	mark(from)
	distance[from.ToInt()] = 0

	for !queue.Empty() {
		v := queue.Pop()
		for w, arc := range net.arcs[v] {
			if nil == arc {
				continue
			}

			utility.Debug("[%v->%v] marked: %v, residuum: %v", v, w, marked[w], arc.residuum())
			if !marked[w] && arc.residuum() > 0 {
				edgeTo[w] = v // Note the last edge on the shortest path.
				distance[w] = distance[v] + 1
				if nil == arc.edge {
					// Dealing with a reverse arc, existing only in the residual net.
					mark(graph.MkVertex(w))
				} else if !arc.edge.IsDeleted() {
					mark(arc.edge.To)
				}
			}
		}
	}

	utility.Debug("Path exists: %v", marked[to.ToInt()])
	return marked[to.ToInt()], pathTo(to), distance
}

// TODO: merge with forAllCoordPairsInPath.
func forEachCoordsInPath(from, to graph.Vertex, edgeTo []int, g *graph.Graph, fn func(int, int, chan<- bool)) {
	done := make(chan bool, 1)
	vi := to.ToInt()
	si := from.ToInt()
	for x := vi; x != si; x = edgeTo[x] {
		fn(x, edgeTo[x], done)
		select {
		case <-done:
			return
		default:
		}
	}
}

func ShortestPathInGraph(g *graph.Graph, from, to graph.Vertex) []*graph.Edge {
	utility.Debug("Looking from path from %v to %v", from, to)
	exists, edgeTo, distance := shortestPathInGraph(g, from, to)
	if !exists {
		return nil
	}

	length := distance[to.ToInt()]
	path := make([]*graph.Edge, length)
	index := length - 1

	forEachCoordsInPath(from, to, edgeTo, g, func(coordFrom, coordTo int, done chan<- bool) {
		path[index] = g.GetEdgeByCoordinates(coordFrom, coordTo)
		index--
	})

	return path
}

func ShortestDistanceInGraph(g *graph.Graph, from, to graph.Vertex) int {
	_, _, distance := shortestPathInGraph(g, from, to)
	return distance[to.ToInt()]
}

func shortestPathInGraph(g *graph.Graph, from, to graph.Vertex) (bool, []int, []int) {
	n := g.CurrentVertexIndex
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	distance := make([]int, n)
	queue := graph.MkIntQueue(n)
	mark := func(v graph.Vertex) {
		vi := v.ToInt()
		marked[vi] = true
		queue.Push(vi)
	}

	mark(from)
	distance[from.ToInt()] = 0

	for !queue.Empty() {
		vi := queue.Pop()
		v := graph.MkVertex(vi)
		g.ForAllNeighbors(v, func(edge *graph.Edge, done chan<- bool) {
			w := graph.GetOtherVertex(v, edge)
			wi := w.ToInt()
			utility.Debug("[%v->%v] marked: %v", v, w, marked[wi])
			if !marked[wi] {
				edgeTo[wi] = vi // Note the last edge on the shortest path.
				distance[wi] = distance[vi] + 1
				mark(w)
			}
		})
	}

	return marked[to.ToInt()], edgeTo, distance
}
