package graph

import "github.com/deckarep/golang-set"

const REDUCING_EMPTY_RESULT = -128

type ChenKanjXiaVC struct {
	G *Graph
	// This will contain the tuples being decomposed by Reducing.
	// The paper states that they are supposed to reside in T, but nowhere does
	// it state what priority should they have.
	// TODO: Find a case that will render this solution impossible - such as one
	// requiring to maintain a specific order of the tuples, particularly one
	// where tuples are not being invalidated between 2 recursive calls.
	// (because a.4, b and c is not applicable)
	tuples mapset.Set
	T      *StructurePriorityQueueProxy
	k      int
	halt   chan bool
}

func (self *ChenKanjXiaVC) conditionalGeneralFold() (wasApplicable bool) {
	gPrime, kPrime1 := generalFold(self.G, self.halt, self.k)
	reduction := self.k - kPrime1
	self.k -= reduction

	var kPrime2 int
	// if there exists a strong 2-tuple ({ u , z }, 1 ) in T then
	if self.T.ContainsStrong2Tuple() {
		if reduction >= 1 {
			wasApplicable = true
			_, kPrime2 = generalFold(self.G, self.halt, kPrime1)
			reduction = kPrime1 - kPrime2
			self.k -= reduction

			if reduction >= 1 {
				// if the repeated application of self.General_Fold reduces the parameter by at least 2 then apply it repeatedly;
				for reduction >= 1 {
					kPrime1 = kPrime2
					_, kPrime2 = generalFold(self.G, self.halt, kPrime2)
					reduction = kPrime1 - kPrime2
					self.k -= reduction
				}

				// General-Fold is no longer applicable.
				return
			} else {
				// else if the application of self.General-Fold reduces
				// the parameter by 1 and (d ( u ) < 4)
				strong2Tuples := self.T.PopAllStrong2Tuples()
				var theTuple *structure
				for s2tInter := range strong2Tuples.Iter() {
					s2t := s2tInter.(*structure)
					gp := &goodPair{
						pair: s2t,
					}
					// According to preliminaries, the notation d(v) means the
					// degree of the vertex in self.G.
					if self.G.Degree(gp.U()) < 4 {
						theTuple = s2t
						break
					}
				}

				if theTuple != nil {
					// then apply it until it is no longer applicable;
					gPrime, kPrime2 = generalFold(gPrime, self.halt, kPrime1)
					reduction = kPrime1 - kPrime2
					self.k -= reduction
					for reduction >= 1 {
						kPrime1 = kPrime2
						gPrime, kPrime2 = generalFold(gPrime, self.halt, kPrime2)
						reduction = kPrime1 - kPrime2
						self.k -= reduction
					}
				}

				return

			}
		}

		panic("2-tuple exists but no reduction could be achieved - there must be a bug in generalFold.")
	} else {
		// else apply General-Fold until it is no longer applicable;
		for reduction >= 1 {
			wasApplicable = true
			gPrime, kPrime2 = generalFold(gPrime, self.halt, kPrime1)
			reduction = kPrime1 - kPrime2
			self.k -= reduction
			kPrime1 = kPrime2
		}
	}

	return
}

func (self *ChenKanjXiaVC) conditionalStruction() (wasApplicable bool) {
	reduction := 0
	// if there exists a strong 2-tuple { u , z} in T then
	if self.T.ContainsStrong2Tuple() {
		s2t, _ := self.T.Pop()
		gp := &goodPair{
			pair: s2t,
		}

		// if there exists w ∈ { u , z} such that d (w) = 3 and the Struction is applicable to w then apply it;
		nu, nuSet := self.G.getNeighborsWithSet(gp.U())
		if self.G.Degree(gp.U()) == 3 &&
			gp.U().isStructionApplicable(self.G, nuSet) {
			_, reduction = structionWithGivenNeighbors(self.G, gp.U(), nu, nuSet)
			wasApplicable = reduction > 0
		} else {
			nz, nzSet := self.G.getNeighborsWithSet(gp.Z())
			if self.G.Degree(gp.Z()) == 3 &&
				gp.Z().isStructionApplicable(self.G, nzSet) {
				_, reduction = structionWithGivenNeighbors(self.G, gp.Z(), nz, nzSet)
				wasApplicable = reduction > 0
			}

			self.k -= reduction
		}

		return
	} else {
		// else if there exists a vertex u ∈ self.G where d ( u ) = 3 or d ( u ) = 4 and such that the Struction is applicable to u
		self.G.ForAllVertices(func(u Vertex, done chan<- bool) {
			deg := self.G.Degree(u)
			if deg == 3 || deg == 4 {
				nv, nvSet := self.G.getNeighborsWithSet(u)
				if u.isStructionApplicable(self.G, nvSet) {
					// then apply it;
					_, reduction = structionWithGivenNeighbors(self.G, u, nv, nvSet)
					wasApplicable = reduction > 0
					done <- true
				}
			}
		})

		self.k -= reduction
	}

	return
}

func (self *ChenKanjXiaVC) updateTuplesByInclusion(includedVertices ...Vertex) {
	for _, includedVertex := range includedVertices {
		for t := range self.tuples.Iter() {
			tuple := t.(*structure)
			// If a vertex u ∈ S is removed from the graph by including it
			// in the cover, the vertex is removed from S and q is unchanged.
			if tuple.q > 0 && tuple.S.Contains(includedVertex) {
				tuple.S.Remove(includedVertex)
			}
		}
	}
}

func (self *ChenKanjXiaVC) updateTuplesByExclusion(excludedVertices ...Vertex) {
	toRemove := mapset.NewSet()

	for _, excludedVertex := range excludedVertices {
		for t := range self.tuples.Iter() {
			tuple := t.(*structure)
			// If one of the vertices in S is removed and is excluded from the
			// cover, then the tuple is modified by removing the vertex from S
			// and decrementing q by 1.
			if tuple.S.Contains(excludedVertex) {
				tuple.q -= 1
				if tuple.q == 0 {
					// If q = 0 then the tuple S will be removed because the
					// information represented by ( S , q ) is satisfied by
					// any minimum vertex cover.
					toRemove.Add(tuple)
				} else {
					tuple.S.Remove(excludedVertex)
				}
			}
		}
	}

	self.tuples = self.tuples.Difference(toRemove)
}

func (self *ChenKanjXiaVC) invalidateTuples() {
	self.tuples.Clear()
}

func (self *ChenKanjXiaVC) reducing() int {
	// NOTE: If Reducing is not applicable, it might as well return self,
	// since no reduction could be achieved.
	rejectedTuples := mapset.NewSet()
	// a. for each tuple ( S , q ) ∈ T do
	for s := range self.tuples.Iter() {
		tuple := s.(*structure)
		// a.1. if | S | < q then reject;
		if tuple.S.Cardinality() < tuple.q {
			rejectedTuples.Add(tuple)
			continue
		}
		// a.2. for every vertex u ∈ S do
		for u := range tuple.S.Iter() {
			if tuple.q == 1 {
				// most likely won't happen
				rejectedTuples.Add(tuple)
				continue
			}

			// T = T ∪ {( S − { u }, q − 1 )};
			S := tuple.S.Clone()
			S.Remove(u.(Vertex))
			newTuple := MkStructureWithSet(tuple.q-1, S)
			if newTuple.computePriority(self.G) < 3 {
				// It's a 2-tuple.
				// When the algorithm branches on a vertex in a 2-tuple, this
				// vertex is picked as follows.
				// If there is a vertex w ∈ S = { u , z } such that w has a
				// neighbor u where u is almost-dominated by the vertex in
				// S − {w} , then the algorithm will branch on the vertex in
				// S − {w} (that is, if there is a vertex in S with a neighbor
				// that is almost-dominated by the other vertex in S, then the
				// algorithm will pick the other vertex in S).
				// Otherwise, it will pick a vertex in S arbitrarily and
				// branch on it.
				// We will always assume that the vertex in the 2-tuple S={u, z}
				// that the algorithm branches on is z.
				// The algorithm can be made oblivious to this choice
				// by ordering the vertices in a 2-tuple as described above
				// whenever the 2-tuple is created.
				verts := newTuple.S.Iter()
				v1, v2 := (<-verts).(Vertex), (<-verts).(Vertex)
				newTuple.S.Remove(v1)
				newTuple.S.Remove(v2)
				newTuple.S.Add(v2)
				newTuple.S.Add(v1)
				// TODO: Verify that the vertex order is retained.
			}

			self.T.Push(newTuple, self.G)
		}

		// a.3. if S is not an independent set then
		isIndependent, edges := isIndependentSet(tuple.S, self.G)
		if !isIndependent {
			// T = T ∪ (\forall(u ,v)∈ E , u ,v∈ S {( S − { u , v}, q − 1 )}) ;
			S := tuple.S.Clone()
			for _, edge := range edges {
				S.Remove(edge.from)
				S.Remove(edge.to)
			}

			self.T.Push(MkStructureWithSet(tuple.q-1, S), self.G)
		}

		// Some vertices might be determined to be in a minimum vertex cover by
		// step a.4 of Reducing.
		// a.4. if there exists v ∈ self.G such that |N(v) ∩ S|≥|S|−q+1 then
		// return (1 + VC (self.G−v, T, k−1)); exit;
		// This condition means that v has to be adjacent to a vertex in S -
		// if it is, it is included in the cover.
		var includedVertex Vertex
		self.G.ForAllVertices(func(v Vertex, done chan<- bool) {
			_, Nv := self.G.getNeighborsWithSet(v)
			if Nv.Intersect(tuple.S).Cardinality() >= tuple.S.Cardinality()-tuple.q+1 {
				// v can be included in the cover, a new instance of the problem
				// has to be run.
				includedVertex = v
				done <- true
			}
		})

		newProblemInstance := self.Copy()
		newProblemInstance.updateTuplesByInclusion(includedVertex)
		newProblemInstance.k--
		return newProblemInstance.VC()
	}

	// If we maintain existing tuples, then the constraints imposed by the newly
	// generated tuples may conflict with those imposed by existing ones.
	// To overcome this, and since the algorithm only processes 2-tuples, when
	// the subroutine Reducing finishes processing the tuples in step a, we will
	// maintain only one 2-tuple and invalidate the rest.
	// Therefore, if 2-tuples exist after step a of Reducing, we will pick any
	// strong 2-tuple in case a strong 2-tuple exists and invalidate the rest,
	// or we will pick any 2-tuple and invalidate
	// the rest, otherwise.

	// TODO: Look for operations based on Lemma 5.1. and see if all vertices in G have to actually be checked, or just a neighborhodd.

	// To be on the safe side with regard to conflicting constraints imposed by
	// tuples and CS/CGF, when we decide to apply the S or the GF operations, we
	// will invalidate all the constraints imposed by the tuples.
	// That is, we will basically remove all the tuples.

	// b. if Conditional_General_Fold(G) or Conditional_Struction(G) in the self.given order is applicable then
	// apply it; exit;
	if self.conditionalGeneralFold() || self.conditionalStruction() {
		self.invalidateTuples()
		return self.k
	}

	// c. if there are vertices u and v in self.G such that v dominates u then return (1 + VC ( self.G − v, T , k − 1 ) ); exit;
	self.G.ForAllVertices(func(v Vertex, done chan<- bool) {

	})

	return self.k
}

func (self *ChenKanjXiaVC) Copy() *ChenKanjXiaVC {
	return &ChenKanjXiaVC{
		G:      self.G.Copy(),
		tuples: self.tuples.Clone(),
		k:      self.k,
		halt:   self.halt,
	}
}

func (self *ChenKanjXiaVC) VC() int {
	panic( /*errors.New(*/ "Not implemented." /*)*/)
}
