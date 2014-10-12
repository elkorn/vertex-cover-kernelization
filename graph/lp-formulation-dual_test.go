package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDualReduction(t *testing.T) {
	g := mkGraph1()
	InVerboseContext(func() {
		nt := mkNtDualReduction(g, 10)
		err := nt.solve()
		assert.Nil(t, err)
	})
}
