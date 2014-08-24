package graph

type Crown struct {
	I Vertices // An independent subset of G
	H Vertices // Head of the crown: N(I) - neighbors of I
	// Also, there must exist a matching M on the edges connecting I and H
	// such that all elements of H are matched.
}

func (self *Crown) Width() int {
	return len(self.H)
}
