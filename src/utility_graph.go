package graph

import "errors"

func getOtherVertex(v Vertex, edge *Edge) Vertex {
	if edge.from != v {
		return edge.from
	}

	if edge.to != v {
		return edge.to
	}

	panic(errors.New("An edge with the same vertex as both endpoints may not exist."))
}

func (self *Graph) getEdgeByCoordinates(from, to int) *Edge {
	result := self.neighbors[from][to]
	if nil == result {
		return self.neighbors[to][from]
	}

	return result
}

func mkGraph1() *Graph {
	/*
		   1o---o2
			|\ /|
			| o5|
			|/ \|
		   4o---o3
	*/
	g := MkGraph(5)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	return g
}

func mkGraph2() *Graph {
	/*
		   1o--------o2
			|\      /|
			|5o----o6|
			| |    | |
			|8o----o7|
			|/      \|
		   4o--------o3
	*/
	g := MkGraph(8)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(1, 5)
	g.AddEdge(2, 6)
	g.AddEdge(3, 7)
	g.AddEdge(4, 8)
	g.AddEdge(5, 6)
	g.AddEdge(6, 7)
	g.AddEdge(7, 8)
	g.AddEdge(8, 5)
	return g
}

func mkGraph3() *Graph {
	/*
	           1
	          / \
	     3---+   +---2
	    / \         / \
	   7   6       5   4
	*/
	g := MkGraph(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6)
	g.AddEdge(3, 7)

	return g
}

func mkGraph4() *Graph {
	/*
	           1
	          / \
	     3---+   +---2
	    / \         / \
	   7---6       5---4
	*/

	g := mkGraph3()

	g.AddEdge(6, 7)
	g.AddEdge(4, 5)

	return g
}

func mkGraph5() *Graph {
	/*
		  1   6
		 / \ / \
		3   2   7
		   / \
		  5---4
	*/

	g := MkGraph(7)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(4, 5)
	g.AddEdge(6, 7)
	g.AddEdge(6, 2)

	return g
}

func mkGraph6() *Graph {
	/*
			3        6
			 \      /
		   2--4----5--7
			 /      \
			1        8
	*/
	g := MkGraph(8)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(5, 7)
	g.AddEdge(5, 8)
	return g
}

func mkPetersenGraph() *Graph {
	g := MkGraph(10)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 1)
	g.AddEdge(1, 6)
	g.AddEdge(6, 9)
	g.AddEdge(9, 7)
	g.AddEdge(7, 10)
	g.AddEdge(10, 8)
	g.AddEdge(8, 6)
	g.AddEdge(7, 2)
	g.AddEdge(8, 3)
	g.AddEdge(9, 4)
	g.AddEdge(10, 5)
	return g
}
