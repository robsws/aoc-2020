package vector

// Vec3 - 3d vector
type Vec3 struct {
	X int
	Y int
	Z int
}

// Mul - multiply vector by scalar
func (v Vec3) Mul(scalar int) Vec3 {
	return Vec3{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

// Add - add another vector to this vector
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z}
}
