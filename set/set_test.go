package set

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := New()

	// Add elements
	s.Add("apple")
	s.Add("banana")

	// Test size
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}

	// Test Contains
	if !s.Contains("apple") {
		t.Errorf("expected set to contain 'apple'")
	}
	if !s.Contains("banana") {
		t.Errorf("expected set to contain 'banana'")
	}
}

func TestRemove(t *testing.T) {
	s := New()

	// Add and remove elements
	s.Add("apple")
	s.Remove("apple")

	// Test size
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	// Test Contains
	if s.Contains("apple") {
		t.Errorf("did not expect set to contain 'apple'")
	}
}

func TestContains(t *testing.T) {
	s := New()

	// Add elements
	s.Add("apple")

	// Test Contains
	if !s.Contains("apple") {
		t.Errorf("expected set to contain 'apple'")
	}
	if s.Contains("banana") {
		t.Errorf("did not expect set to contain 'banana'")
	}
}

func TestSize(t *testing.T) {
	s := New()

	// Test size of empty set
	if s.Size() != 0 {
		t.Errorf("expected size 0, got %d", s.Size())
	}

	// Add elements and test size
	s.Add("apple")
	s.Add("banana")
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}
}

func TestList(t *testing.T) {
	s := New()

	// Add elements
	s.Add("apple")
	s.Add("banana")

	// Test List
	list := s.List()
	if len(list) != 2 {
		t.Errorf("expected list length 2, got %d", len(list))
	}

	// Check contents
	foundApple := false
	foundBanana := false
	for _, item := range list {
		if item == "apple" {
			foundApple = true
		}
		if item == "banana" {
			foundBanana = true
		}
	}
	if !foundApple || !foundBanana {
		t.Errorf("expected list to contain 'apple' and 'banana', got %v", list)
	}
}

func TestUnion(t *testing.T) {
	s1 := New()
	s2 := New()

	// Add elements
	s1.Add("apple")
	s1.Add("banana")
	s2.Add("banana")
	s2.Add("cherry")

	// Test Union
	union := s1.Union(s2)
	if union.Size() != 3 {
		t.Errorf("expected union size 3, got %d", union.Size())
	}

	// Check contents
	if !union.Contains("apple") || !union.Contains("banana") || !union.Contains("cherry") {
		t.Errorf("expected union to contain 'apple', 'banana', and 'cherry'")
	}
}

func TestIntersection(t *testing.T) {
	s1 := New()
	s2 := New()

	// Add elements
	s1.Add("apple")
	s1.Add("banana")
	s2.Add("banana")
	s2.Add("cherry")

	// Test Intersection
	intersection := s1.Intersection(s2)
	if intersection.Size() != 1 {
		t.Errorf("expected intersection size 1, got %d", intersection.Size())
	}

	// Check contents
	if !intersection.Contains("banana") {
		t.Errorf("expected intersection to contain 'banana'")
	}
}

func TestDifference(t *testing.T) {
	s1 := New()
	s2 := New()

	// Add elements
	s1.Add("apple")
	s1.Add("banana")
	s2.Add("banana")
	s2.Add("cherry")

	// Test Difference
	difference := s1.Difference(s2)
	if difference.Size() != 1 {
		t.Errorf("expected difference size 1, got %d", difference.Size())
	}

	// Check contents
	if !difference.Contains("apple") {
		t.Errorf("expected difference to contain 'apple'")
	}
}
