package graph

import "log"

func shortestPathFromSourceToSink(nf *NetworkFlow) (bool, []int, []int) {
	return shortestPath(nf.net, nf.source, nf.sink)
}

// TODO add bool parameter allowing to not compute the path itself
func shortestPath(net Net, from, to Vertex) (bool, []int, []int) {
	n := len(net.arcs)
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	distance := make([]int, n)
	queue := MkQueue(n)
	mark := func(v Vertex) {
		vi := v.toInt()
		marked[vi] = true
		queue.Push(vi)
	}
	pathTo := func(v Vertex) []int {
		vi := v.toInt()
		si := from.toInt()

		inVerboseContext(func() {
			Debug("vi: %v, marked: %v", vi, marked)
		})

		if !marked[vi] {
			return nil
		}

		path := MkStack()

		for x := vi; x != si; x = edgeTo[x] {
			path.Push(x)
		}

		path.Push(si)
		return path.Values() // it can be done much better performance-wise.
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
					log.Println("Marking reverse arc %v", w+1)
					// Dealing with a reverse arc, existing only in the residual net.
					mark(MkVertex(w))
				} else if !arc.edge.isDeleted {
					mark(arc.edge.to)
				}
			}
		}
	}

	log.Println("Path from", from, "to", to, "in", n)
	return marked[to.toInt()], pathTo(to), distance
}
