package utils

import (
	"math"
)

// PowInt - calculate power with integers
func PowInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
