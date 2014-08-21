package graph

import "testing"

func TestMkExampleGraph(t *testing.T) {
	inVerboseContext(func() {
		Debug("%v", string(MkExampleGraph(10)))
	})
}
