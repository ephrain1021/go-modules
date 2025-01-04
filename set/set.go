package set

// Set is a collection that holds unique string elements.
type Set struct {
	elements map[string]bool
}

// New creates and initializes a new Set.
func New() *Set {
	return &Set{
		elements: make(map[string]bool),
	}
}

// Add inserts an element into the Set.
func (s *Set) Add(value string) {
	s.elements[value] = true
}

// Remove deletes an element from the Set.
func (s *Set) Remove(value string) {
	delete(s.elements, value)
}

// Contains checks if the Set contains a specific element.
func (s *Set) Contains(value string) bool {
	return s.elements[value]
}

// Size returns the number of elements in the Set.
func (s *Set) Size() int {
	return len(s.elements)
}

// List returns all elements in the Set as a slice.
func (s *Set) List() []string {
	keys := make([]string, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

// Union returns a new Set that contains all elements from the current Set and another Set.
func (s *Set) Union(other *Set) *Set {
	unionSet := New()
	for key := range s.elements {
		unionSet.Add(key)
	}
	for key := range other.elements {
		unionSet.Add(key)
	}
	return unionSet
}

// Intersection returns a new Set that contains elements present in both Sets.
func (s *Set) Intersection(other *Set) *Set {
	intersectionSet := New()
	for key := range s.elements {
		if other.Contains(key) {
			intersectionSet.Add(key)
		}
	}
	return intersectionSet
}

// Difference returns a new Set that contains elements present in the current Set but not in the other Set.
func (s *Set) Difference(other *Set) *Set {
	differenceSet := New()
	for key := range s.elements {
		if !other.Contains(key) {
			differenceSet.Add(key)
		}
	}
	return differenceSet
}
