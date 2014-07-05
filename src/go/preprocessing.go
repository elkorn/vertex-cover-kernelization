package graph

import (
	"errors"
	"fmt"
)

func RemoveVerticesOfDegree(g *Graph, degree int) error {
	for vertex := range g.Vertices {
		vDegree, err := g.Degree(vertex)
		if nil != err {
			return errors.New(fmt.Sprintf("Vertex %v does not exist in the graph.", vertex))
		}

		if vDegree == degree {
			g.RemoveVertex(vertex)
		}
	}

	return nil
}
