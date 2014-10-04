package graph

import (
	"fmt"
	"testing"
)

func TestMkExampleGraph(t *testing.T) {
	for i := 0; i < 15; i++ {
		WriteExampleGraph(fmt.Sprintf("../examples/example_%d", i+1), (i+1)*10)
	}
}
