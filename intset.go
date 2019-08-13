package main

type IntSet []int

// Contains returns true if x is in the set, otherwise false
func (s *IntSet) Contains(x int) bool {
	for _, v := range *s {
		if v == x {
			return true
		}
	}
	return false
}

// Insert adds x to the set if it is not already present
func (s *IntSet) Insert(x int) {
	if s.Contains(x) {
		return
	}
	*s = append(*s, x)
}
