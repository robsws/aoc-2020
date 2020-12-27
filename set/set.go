package set

type Set interface {
	Add(item interface{})
	Remove(item interface{}) interface{}
	Copy() Set
	Len() int
	Union(other Set)
	Contains(item interface{}) bool
	ToSlice() []interface{}
}

type empty struct{}
