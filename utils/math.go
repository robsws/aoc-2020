package utils

import (
	"math"
)

const MAXINT = int(^uint(0) >> 1)

// PowInt - calculate power with integers
func PowInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// AbsInt - abs for ints
func AbsInt(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}

// MaxInt - max of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MinIntInList - get minimum int from a list and it's index
func MinIntInList(list []int) (int, int) {
	min := MAXINT
	minI := -1
	for i, n := range list {
		if n < min {
			min = n
			minI = i
		}
	}
	return minI, min
}

// MaxIntInList - get maximum int from a list and it's index
func MaxIntInList(list []int) (int, int) {
	max := -MAXINT
	maxI := -1
	for i, n := range list {
		if n > max {
			max = n
			maxI = i
		}
	}
	return maxI, max
}

// Vec2 - 2d vector
type Vec2 struct {
	X int
	Y int
}

// VecMul - multiply vector by scalar
func VecMul(v Vec2, scalar int) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

// VecAdd - add another vector to this vector
func VecAdd(v1 Vec2, v2 Vec2) Vec2 {
	return Vec2{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}
