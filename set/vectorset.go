package set

import "aoc-go/vector"

// Vec2Set - set of vectors
type Vec2Set struct {
	m map[vector.Vec2]empty
}

// MakeVec2Set - create an empty set
func MakeVec2Set() Vec2Set {
	return Vec2Set{make(map[vector.Vec2]empty)}
}

// MakeVec2SetFromSlice - create a set from a slice
func MakeVec2SetFromSlice(slice []vector.Vec2) Vec2Set {
	s := MakeVec2Set()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// Copy - copy this integer set
func (s *Vec2Set) Copy() Vec2Set {
	var e empty
	ns := MakeVec2Set()
	for k := range s.m {
		ns.m[k] = e
	}
	return ns
}

// Len - length of a set
func (s Vec2Set) Len() int {
	return len(s.m)
}

// Add - add item to set
func (s *Vec2Set) Add(item vector.Vec2) {
	var e empty
	s.m[item] = e
}

// Remove - remove item from set
func (s *Vec2Set) Remove(item vector.Vec2) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s *Vec2Set) Union(other Vec2Set) {
	list := other.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Intersection - get items in both sets
func (s *Vec2Set) Intersection(other Vec2Set) {
	list := s.ToSlice()
	for _, item := range list {
		if !other.Contains(item) {
			s.Remove(item)
		}
	}
}

// Contains - check if item is in the set
func (s Vec2Set) Contains(item vector.Vec2) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s Vec2Set) ToSlice() []vector.Vec2 {
	result := make([]vector.Vec2, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
