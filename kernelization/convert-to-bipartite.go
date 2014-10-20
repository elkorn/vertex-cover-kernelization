package kernelization

import "github.com/elkorn/vertex-cover-kernelization/graph"

func getVertices(g *graph.Graph) graph.Vertices {
	result := make(graph.Vertices, 0)
	n := len(g.Vertices)
	for _, v := range g.Vertices {
		result = append(result, v)
	}

	for _, v := range result {
		result = append(result, graph.MkVertex(v.ToInt()+n))
	}

	return result
}

func addBipartiteEdges(g *graph.Graph, original *graph.Graph, border int) {
	original.ForAllEdges(func(edge *graph.Edge, done chan<- bool) {
		// Invariant: F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
		g.AddEdge(edge.From, graph.MkVertex(edge.To.ToInt()+border))
	})
}

func makeBipartite(g *graph.Graph) *graph.Graph {
	/*
		Convert G(V,E) to a bipartite graph H=(U,F) with the following properties:
		A = V
		B = V
		U = A \sum B
		F = {(A_v,B_u)|(v,u) \in E or (u,v) \in E}
	*/

	border := g.CurrentVertexIndex
	result := graph.MkGraphRememberingDeletedVertices(border*2, g.IsVertexDeleted) // remember deleted graph.vertices
	addBipartiteEdges(result, g, border)
	return result
}

func makeBipartiteForNetworkFlow(g *graph.Graph) *graph.Graph {
	border := len(g.Vertices)
	result := graph.MkGraphRememberingDeletedVertices(border*2+2, g.IsVertexDeleted) // remember deleted graph.vertices
	addBipartiteEdges(result, g, border)
	return result
}
