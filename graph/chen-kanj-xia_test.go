package graph

// TODO: Test this after fixing findCrown. It's slow and crashes.
// Also, come up with better test scenarios.
//

// func TestConditionalGeneralFold(t *testing.T) {
// 	g := ScanGraph("../examples/sh2/sh2-3.dim.sh")
// 	T := identifyStructures(g, 246)
// 	instance := &ChenKanjXiaVC{
// 		G:      g,
// 		T:      T,
// 		tuples: mapset.NewSet(),
// 		k:      246,
// 		halt:   nil,
// 	}

// 	InVerboseContext(func() {
// 	instance.conditionalGeneralFold()
// 	})
// }
