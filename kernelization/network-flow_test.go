package kernelization

import (
	"testing"

	"github.com/elkorn/vertex-cover-kernelization/graph"
	"github.com/stretchr/testify/assert"
)

func TestMkToNetworkFlow(t *testing.T) {
	g := graph.MkGraph(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	source := graph.Vertex(7)
	sink := graph.Vertex(8)

	result := mkNetworkFlow(g)

	// 1. we need source and sink nodes.
	assert.True(t, result.graph.HasVertex(7))
	assert.True(t, result.graph.HasVertex(8))
	assert.Equal(t, source, result.source)
	assert.Equal(t, sink, result.sink)
	// 2. Source has to be connected to every vertex in set A.
	for i := 1; i <= 3; i++ {
		assert.True(t, result.graph.HasEdge(source, graph.Vertex(i)), "Source has to be connected to all edges in set A.")
	}

	// 3. Sink has to be connected to every vertex in set B.
	for i := 4; i <= 6; i++ {
		assert.True(t, result.graph.HasEdge(graph.Vertex(i), sink), "Sink has to be connected to all edges in set B.")
	}
}

func TestNet(t *testing.T) {
	g := graph.MkGraph(4)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)

	netFlow := mkNetworkFlow(g)
	assert.Equal(t, netFlow.graph.NVertices(), len(netFlow.net.arcs), "The flow net should be an NxN matrix.")
	netFlow.graph.ForAllEdges(func(edge *graph.Edge, _ chan<- bool) {
		arc := netFlow.net.arcs[edge.From-1][edge.To-1]
		assert.NotNil(t, arc, "Each edge in the network flow must be represented by an Arc.")
		assert.Equal(t, 1, arc.capacity, "Every arc must have an initial capacity of 1.")
		assert.Equal(t, 0, arc.flow, "Every arc must have an initial flow of 0.")
	})

	assert.Equal(t, 2, netFlow.net.length[0], "The Net structure must contain information about the number of arcs going out of a specified vertex.")
	assert.Equal(t, 2, netFlow.net.length[1], "The Net structure must contain information about the number of arcs going out of a specified vertex.")
}

// func TestNetworkFlowProteins(t *testing.T) {
// 	g := ScanGraph("../examples/sh2/sh2-3.dim")

// 	InVerboseContext(func() {
// 		kPrime := networkFlowKernelization(g, 246)
// 		log.Println(kPrime)
// 	})

// }
