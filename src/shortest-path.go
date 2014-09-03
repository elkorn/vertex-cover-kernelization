package graph

func shortestPathFromSourceToSink(nf *NetworkFlow) (bool, []int, []int) {
	return shortestPath(nf.net, nf.source, nf.sink)
}

// TODO: migrate to the new logic.
func shortestPath(net Net, from, to Vertex) (bool, []int, []int) {
	n := len(net.arcs)
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	distance := make([]int, n)
	queue := MkIntQueue(n)
	mark := func(v Vertex) {
		vi := v.toInt()
		marked[vi] = true
		queue.Push(vi)
	}
	pathTo := func(v Vertex) []int {
		vi := v.toInt()
		si := from.toInt()

		Debug("vi: %v, marked: %v", vi, marked)

		if !marked[vi] {
			return nil
		}

		path := MkIntStack(distance[v.toInt()])

		for x := vi; x != si; x = edgeTo[x] {
			path.Push(x)
		}

		path.Push(si)
		// TODO: introduce Queue.Iter() to get rid of O(N) here.
		return path.Values()
	}

	mark(from)
	distance[from.toInt()] = 0

	for !queue.Empty() {
		v := queue.Pop()
		for w, arc := range net.arcs[v] {
			if nil == arc {
				continue
			}

			Debug("[%v->%v] marked: %v, residuum: %v", v, w, marked[w], arc.residuum())
			if !marked[w] && arc.residuum() > 0 {
				edgeTo[w] = v // Note the last edge on the shortest path.
				distance[w] = distance[v] + 1
				if nil == arc.edge {
					// Dealing with a reverse arc, existing only in the residual net.
					mark(MkVertex(w))
				} else if !arc.edge.isDeleted {
					mark(arc.edge.to)
				}
			}
		}
	}

	return marked[to.toInt()], pathTo(to), distance
}

// TODO: merge with forAllCoordPairsInPath.
func forEachCoordsInPath(from, to Vertex, edgeTo []int, g *Graph, fn func(int, int, chan<- bool)) {
	done := make(chan bool, 1)
	vi := to.toInt()
	si := from.toInt()
	for x := vi; x != si; x = edgeTo[x] {
		fn(x, edgeTo[x], done)
		select {
		case <-done:
			return
		default:
		}
	}
}

func ShortestPathInGraph(g *Graph, from, to Vertex) []*Edge {
	Debug("Looking from path from %v to %v", from, to)
	exists, edgeTo, distance := shortestPathInGraph(g, from, to)
	if !exists {
		return nil
	}

	length := distance[to.toInt()]
	path := make([]*Edge, length)
	index := length - 1

	forEachCoordsInPath(from, to, edgeTo, g, func(coordFrom, coordTo int, done chan<- bool) {
		path[index] = g.getEdgeByCoordinates(coordFrom, coordTo)
		index--
	})

	return path
}

func ShortestDistanceInGraph(g *Graph, from, to Vertex) int {
	_, _, distance := shortestPathInGraph(g, from, to)
	return distance[to.toInt()]
}

func shortestPathInGraph(g *Graph, from, to Vertex) (bool, []int, []int) {
	n := g.currentVertexIndex
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	distance := make([]int, n)
	queue := MkIntQueue(n)
	mark := func(v Vertex) {
		vi := v.toInt()
		marked[vi] = true
		queue.Push(vi)
	}

	mark(from)
	distance[from.toInt()] = 0

	for !queue.Empty() {
		vi := queue.Pop()
		v := MkVertex(vi)
		g.ForAllNeighbors(v, func(edge *Edge, index int, done chan<- bool) {
			w := getOtherVertex(v, edge)
			wi := w.toInt()
			Debug("[%v->%v] marked: %v", v, w, marked[wi])
			if !marked[wi] {
				edgeTo[wi] = vi // Note the last edge on the shortest path.
				distance[wi] = distance[vi] + 1
				mark(w)
			}
		})
	}

	return marked[to.toInt()], edgeTo, distance
}
