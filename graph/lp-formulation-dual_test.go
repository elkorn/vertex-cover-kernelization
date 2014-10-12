package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDualReduction(t *testing.T) {
	g := mkGraph1()
	InVerboseContext(func() {
		nt := mklpDualFormulation(g, 10)
		err := nt.solve()
		assert.Nil(t, err)
	})
}
