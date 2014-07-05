package graph

func extend(slice []Edge, element Edge) []Edge {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]Edge, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func append(slice []Edge, items ...Edge) []Edge {
	for _, item := range items {
		slice = extend(slice, item)
	}

	return slice
}
