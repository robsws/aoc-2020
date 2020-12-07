package utils

type empty struct{}

// Set - implement a mathematical set using a map
type Set struct {
	m map[string]interface{}
	e empty
}

// MakeSet - create an empty set
func MakeSet() Set {
	var e empty
	s := Set{make(map[string]interface{}), e}
	return s
}

// Len - length of a set
func (s *Set) Len() int {
	return len(Keys(s.m))
}

// Add - add item to set
func (s *Set) Add(item string) {
	s.m[item] = s.e
}

// Union - add another set to this set
func (s *Set) Union(t Set) {
	list := t.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// ToSlice - get elements as slice
func (s *Set) ToSlice() []string {
	return Keys(s.m)
}
