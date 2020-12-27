package set

// CoordSet - set of integers
type CoordSet struct {
	m map[[3]int]empty
}

// MakeCoordSet - create an empty set
func MakeCoordSet() CoordSet {
	return CoordSet{make(map[[3]int]empty)}
}

// MakeCoordSetFromSlice - create a set from a slice
func MakeCoordSetFromSlice(slice [][3]int) CoordSet {
	s := MakeCoordSet()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// Copy - copy this integer set
func (s CoordSet) Copy() CoordSet {
	var e empty
	ns := MakeCoordSet()
	for k := range s.m {
		ns.m[k] = e
	}
	return ns
}

// Len - length of a set
func (s CoordSet) Len() int {
	return len(s.m)
}

// Add - add item to set
func (s CoordSet) Add(item [3]int) {
	var e empty
	s.m[item] = e
}

// Remove - remove item from set
func (s CoordSet) Remove(item [3]int) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s CoordSet) Union(other CoordSet) {
	list := other.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Contains - check if item is in the set
func (s CoordSet) Contains(item [3]int) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s CoordSet) ToSlice() [][3]int {
	result := make([][3]int, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
