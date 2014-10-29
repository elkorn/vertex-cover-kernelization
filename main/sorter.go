package main

type fileSorter struct {
	files []dataFileDescriptor
}

// Len is part of sort.Interface.
func (s *fileSorter) Len() int {
	return len(s.files)
}

// Swap is part of sort.Interface.
func (s *fileSorter) Swap(i, j int) {
	s.files[i], s.files[j] = s.files[j], s.files[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *fileSorter) Less(i, j int) bool {
	if s.files[i].vertices == s.files[j].vertices {
		return s.files[i].degreeDistribution < s.files[j].degreeDistribution
	}

	return s.files[i].vertices < s.files[j].vertices
}
