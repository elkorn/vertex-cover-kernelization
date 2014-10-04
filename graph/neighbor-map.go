package graph

type NeighborMap []Neighbors

func MkNeighborMap(size int) NeighborMap {
	result := make(NeighborMap, size)
	for i, _ := range result {
		result[i] = make(Neighbors, 0, size)
	}

	return result
}

func (self *NeighborMap) AddNeighborOfVertex(v, n Vertex) {
	index := v.toInt()
	if (*self)[index] == nil {
		(*self)[index] = Neighbors{n}
	} else {
		if !contains((*self)[index], n) {
			(*self)[index] = append((*self)[index], n)
		}
	}
}

func (self *NeighborMap) ForAll(fn func(Vertex, Neighbors, chan<- bool)) {
	done := make(chan bool, 1)
	for idx, neighbors := range *self {
		fn(MkVertex(idx), neighbors, done)
		select {
		case <-done:
			return
		default:
		}
	}
}
