package graph

func conditionalGeneralFold(G *Graph, T *StructurePriorityQueueProxy, halt chan bool, k int) (gPrime *Graph, kPrime int) {
	// TODO: Handle halting.
	// See the proof of Lemma 5.2 and 5.3 for examples.
	gPrime, kPrime1 := generalFold(G, halt, k)
	reduction := k - kPrime1
	kPrime = k - reduction
	var kPrime2 int
	// if there exists a strong 2-tuple ({ u , z }, 1 ) in T then
	if T.ContainsStrong2Tuple() {
		if reduction >= 1 {
			gPrime, kPrime2 = generalFold(gPrime, halt, kPrime1)
			reduction = kPrime1 - kPrime2
			kPrime -= reduction
			kPrime1 = kPrime2

			if reduction >= 1 {
				// if the repeated application of General_Fold reduces the parameter by at least 2 then apply it repeatedly;
				for reduction >= 1 {
					gPrime, kPrime2 = generalFold(gPrime, halt, kPrime2)
					reduction = kPrime1 - kPrime2
					kPrime -= reduction
					kPrime1 = kPrime2
				}

				// General-Fold is no longer applicable.
				return
			} else {
				// else if the application of General-Fold reduces
				// the parameter by 1 and (d ( u ) < 4)
				strong2Tuples := T.PopAllStrong2Tuples()
				var theTuple *structure
				for s2tInter := range strong2Tuples.Iter() {
					s2t := s2tInter.(*structure)
					gp := &goodPair{
						pair: s2t,
					}
					// According to preliminaries, the notation d(v) means the
					// degree of the vertex in G.
					if G.Degree(gp.U()) < 4 {
						theTuple = s2t
						break
					}
				}

				if theTuple != nil {
					// then apply it until it is no longer applicable;
					gPrime, kPrime2 = generalFold(gPrime, halt, kPrime1)
					reduction = kPrime1 - kPrime2
					kPrime -= reduction
					for reduction >= 1 {
						kPrime1 = kPrime2
						gPrime, kPrime2 = generalFold(gPrime, halt, kPrime2)
						reduction = kPrime1 - kPrime2
						kPrime -= reduction
					}
				}

				return

			}
		}

		panic("2-tuple exists but no reduction could be achieved - there must be a bug in generalFold.")
	} else {
		// else apply General-Fold until it is no longer applicable;
		for reduction >= 1 {
			gPrime, kPrime2 = generalFold(gPrime, halt, kPrime1)
			reduction = kPrime1 - kPrime2
			kPrime -= reduction
			kPrime1 = kPrime2
		}
	}

	return
}

// func conditionalStruction(G *Graph, T *StructurePriorityQueueProxy, halt chan bool, k int) (gPrime *Graph
