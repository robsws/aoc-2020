package utils

type empty struct{}

// Set - implement a mathematical set using a map
type Set struct {
	m map[interface{}]interface{}
	e empty
}

// MakeSet - create an empty set
func MakeSet() Set {
	var e empty
	s := Set{make(map[interface{}]interface{}), e}
	return s
}

// MakeSetFromSlice - create an empty set
func MakeSetFromSlice(slice []interface{}) Set {
	s := MakeSet()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// CopySet - copy a set
func CopySet(s Set) Set {
	ns := Set{make(map[interface{}]interface{}), s.e}
	for k := range s.m {
		ns.m[k] = ns.e
	}
	return ns
}

// Len - length of a set
func (s Set) Len() int {
	return len(Keys(s.m))
}

// Add - add item to set
func (s Set) Add(item interface{}) {
	s.m[item] = s.e
}

// Remove - remove item from set
func (s Set) Remove(item interface{}) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s Set) Union(t Set) {
	list := t.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Contains - check if item is in the set
func (s Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s Set) ToSlice() []interface{} {
	return Keys(s.m)
}
