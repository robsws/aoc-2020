package vector

// Vec2 - 2d vector
type Vec2 struct {
	X int
	Y int
}

// Mul - multiply vector by scalar
func (v Vec2) Mul(scalar int) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

// Add - add another vector to this vector
func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}
