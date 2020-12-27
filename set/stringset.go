package set

// StringSet - set of strings
type StringSet struct {
	m map[string]empty
}

// MakeStringSet - create an empty set
func MakeStringSet() StringSet {
	return StringSet{make(map[string]empty)}
}

// MakeStringSetFromSlice - create a set from a slice
func MakeStringSetFromSlice(slice []string) StringSet {
	s := MakeStringSet()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// Copy - copy this integer set
func (s StringSet) Copy() StringSet {
	var e empty
	ns := MakeStringSet()
	for k := range s.m {
		ns.m[k] = e
	}
	return ns
}

// Len - length of a set
func (s StringSet) Len() int {
	return len(s.m)
}

// Add - add item to set
func (s StringSet) Add(item string) {
	var e empty
	s.m[item] = e
}

// Remove - remove item from set
func (s StringSet) Remove(item string) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s StringSet) Union(other StringSet) {
	list := other.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Contains - check if item is in the set
func (s StringSet) Contains(item string) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s StringSet) ToSlice() []string {
	result := make([]string, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
