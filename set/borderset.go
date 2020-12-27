package set

// BorderSet - set of bool rows
type BorderSet struct {
	m map[[10]bool]empty
}

// MakeBorderSet - create an empty set
func MakeBorderSet() BorderSet {
	return BorderSet{make(map[[10]bool]empty)}
}

// MakeBorderSetFromSlice - create a set from a slice
func MakeBorderSetFromSlice(slice [][10]bool) BorderSet {
	s := MakeBorderSet()
	for _, elem := range slice {
		s.Add(elem)
	}
	return s
}

// Copy - copy this integer set
func (s BorderSet) Copy() BorderSet {
	var e empty
	ns := MakeBorderSet()
	for k := range s.m {
		ns.m[k] = e
	}
	return ns
}

// Len - length of a set
func (s BorderSet) Len() int {
	return len(s.m)
}

// Add - add item to set
func (s BorderSet) Add(item [10]bool) {
	var e empty
	s.m[item] = e
}

// Remove - remove item from set
func (s BorderSet) Remove(item [10]bool) {
	delete(s.m, item)
}

// Union - add another set to this set
func (s BorderSet) Union(other BorderSet) {
	list := other.ToSlice()
	for _, item := range list {
		s.Add(item)
	}
}

// Contains - check if item is in the set
func (s BorderSet) Contains(item [10]bool) bool {
	_, ok := s.m[item]
	return ok
}

// ToSlice - get elements as slice
func (s BorderSet) ToSlice() [][10]bool {
	result := make([][10]bool, 0)
	for k := range s.m {
		result = append(result, k)
	}
	return result
}
