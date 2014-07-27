package graph

func getVertices(g *Graph) []Vertex {
	result := make([]Vertex, 0)
	n := len(g.Vertices)
	for v := range g.Vertices {
		result = append(result, v)
	}

	for _, v := range result {
		result = append(result, Vertex(int(v)+n))
	}

	return result
}

func addBipartiteEdges(g *Graph) {
	before := len(g.Edges)
	after := 2 * before
	border := before + 1
	// Invariant: F = {(A_v,B_v)|(v,u) \in E or (u,v) \in E}

	for i := before; i < after; i++ {
		orig := g.Edges[i-before]
		g.AddEdge(
			Vertex(int(orig.from)+border),
			Vertex(int(orig.to)+border))
	}

}

// func makeBipartite(g *Graph) *Graph {
// 	before := len(g.Edges)
// 	edges := before * 2
// 	border := before + 1
// 	result := mkGraphWithVertices(len(g.Vertices) * 2)
// 	for i := 0; i < before; i++ {
// 		orig := g.Edges[i]
// 		result.AddEdge(orig.from, orig.to)
// 	}
// //
// 	return result
// }
