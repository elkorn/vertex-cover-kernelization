package graph

func removeAt(source Edges, position int) Edges {
	return append(source[:position], source[position+1:]...)
}
