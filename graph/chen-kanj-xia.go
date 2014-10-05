package graph

const REDUCING_EMPTY_RESULT = -128

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

func conditionalStruction(G *Graph, T *StructurePriorityQueueProxy, halt chan bool, k int) (gPrime *Graph, kPrime int) {
	reduction := 0
	// if there exists a strong 2-tuple { u , z} in T then
	if T.ContainsStrong2Tuple() {
		s2t, _ := T.Pop()
		gp := &goodPair{
			pair: s2t,
		}

		// if there exists w ∈ { u , z} such that d (w) = 3 and the Struction is applicable to w then apply it;
		if nu, nuSet := G.getNeighborsWithSet(gp.U()); G.Degree(gp.U()) == 3 && gp.U().isStructionApplicable(G, nuSet) {
			gPrime, reduction = structionWithGivenNeighbors(G, gp.U(), nu, nuSet)
		} else if nz, nzSet := G.getNeighborsWithSet(gp.Z()); G.Degree(gp.Z()) == 3 && gp.Z().isStructionApplicable(G, nzSet) {
			gPrime, reduction = structionWithGivenNeighbors(G, gp.Z(), nz, nzSet)
		}

		kPrime = k - reduction

		return
	} else {
		// else if there exists a vertex u ∈ G where d ( u ) = 3 or d ( u ) = 4 and such that the Struction is applicable to u
		G.ForAllVertices(func(u Vertex, done chan<- bool) {
			deg := G.Degree(u)
			if deg == 3 || deg == 4 {
				nv, nvSet := G.getNeighborsWithSet(u)
				if u.isStructionApplicable(G, nvSet) {
					// then apply it;
					gPrime, reduction = structionWithGivenNeighbors(G, u, nv, nvSet)
					done <- true
					return
				}
			}
		})

		kPrime = k - reduction
		return
	}
}

func reducing(G *Graph, T *StructurePriorityQueueProxy, halt chan bool, k int) int {
	// NOTE: If Reducing is not applicable, it might as well return 0,
	// since no reduction could be achieved.

	// a. for each tuple ( S , q ) ∈ T do
	// a.1. if | S | < q then reject;
	// a.2. for every vertex u ∈ S do T = T ∪ {( S − { u }, S
	// q − 1 )} ;
	// a.3. if S is not an independent set then T = T ∪ ( ( u ,v)∈ E , u ,v∈ S {( S − { u , v}, q − 1 )}) ;
	// a.4. if there exists v ∈ G such that | N (v) ∩ S | ≥ | S | − q + 1 then return (1 + VC ( G − v, T , k − 1 ) ); exit;
	// b. if Conditional_General_Fold(G) or Conditional_Struction(G) in the given order is applicable then
	// apply it; exit;
	// c. if there are vertices u and v in G such that v dominates u then return (1 + VC ( G − v, T , k − 1 ) ); exit;

}
