package main

import "testing"

func TestContains(t *testing.T) {
	s := IntSet{}
	if s.Contains(0) {
		t.Errorf("The empty set should not contain 0")
	}
	xs := []int{1, 2, 3}
	s = IntSet(xs)
	for _, x := range xs {
		if !s.Contains(x) {
			t.Errorf("Expected set to contain %d", x)
		}
	}
	if s.Contains(0) {
		t.Errorf("The set {1,2,3} should not contain 0")
	}
}

func TestInsert(t *testing.T) {
	s := IntSet{}
	xs := []int{1, 2, 3}
	for _, x := range xs {
		if s.Contains(x) {
			t.Errorf("%d has not yet been added to the set", x)
		}
		s.Insert(x)
		if !s.Contains(x) {
			t.Errorf("Expected the set to contain %d", x)
		}
	}
	if len(s) != len(xs) {
		t.Errorf("The set should contain %d elements", len(xs))
	}
	s.Insert(1)
	if len(s) != len(xs) {
		t.Errorf("Inserting an element twice should not affect size of the set")
	}
}
