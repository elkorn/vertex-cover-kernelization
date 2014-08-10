package graph

func shortestPathFromSourceToSink(nf *NetworkFlow) (bool, Container) {
	n := len(nf.graph.Vertices)
	marked := make([]bool, n) // Is there a known shortest path to a vertex?
	edgeTo := make([]int, n)  // The last vertex on the known path to a vertex.
	queue := MkQueue(n)
	mark := func(v Vertex) {
		vi := v.toInt()
		marked[vi] = true
		queue.Push(vi)
	}
	pathTo := func(v Vertex) *Stack {
		vi := v.toInt()
		si := nf.source.toInt()

		if !marked[vi] {
			return nil
		}

		path := MkStack()

		for x := vi; x != si; x = edgeTo[x] {
			path.Push(x)
		}

		path.Push(si)
		return path
	}

	mark(nf.source)

	for !queue.Empty() {
		v := queue.Pop()
		for w, arc := range nf.net.arcs[v] {
			if nil == arc {
				continue
			}

			Debug("[%v->%v] marked: %v, residuum: %v", v, w, marked[w], arc.residuum())
			if !marked[w] && arc.residuum() > 0 {
				edgeTo[w] = v // Note the last edge on the shortest path.
				if nil == arc.edge {
					// Dealing with a reverse arc, existing only in the residual net.
					// This case is treated explicitly only for clarity.
					mark(Vertex(w + 1))
				} else {
					mark(arc.edge.to)
				}
			}
		}
	}

	return marked[nf.sink.toInt()], pathTo(nf.sink)

}
