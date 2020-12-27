package set

// IntSet - set of integers
type IntSet struct {
	m map[int]empty
}

// MakeIntSet - create an empty set
func MakeIntSet() IntSet {
	return IntSet{make(map[int]empty)}
}

// MakeIntSetFromSlice - create a set from a slice
func MakeIntSetFromSlice(slice []int) IntSet {
	s := MakeIntSet()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// Copy - copy this integer set
func (s IntSet) Copy() IntSet {
	var e empty
	ns := MakeIntSet()
	for k := range s.m {
		ns.m[k] = e
	}
	return ns
}

// Len - length of a set
func (s IntSet) Len() int {
	return len(s.m)
}

// Add - add item to set
func (s IntSet) Add(item int) {
	var e empty
	s.m[item] = e
}

// Remove - remove item from set
func (s IntSet) Remove(item int) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s IntSet) Union(other IntSet) {
	list := other.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Contains - check if item is in the set
func (s IntSet) Contains(item int) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s IntSet) ToSlice() []int {
	result := make([]int, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
