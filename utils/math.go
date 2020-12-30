package utils

import (
	"math"
)

// MAXINT - maximum integer value
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

// IntPermutations - get all permutations of an int slice
func IntPermutations(list []int) [][]int {
	permutations := make([][]int, 0)
	if len(list) == 1 {
		permutations = append(permutations, list)
		return permutations
	}
	for _, elem := range list {
		// Make a smaller list of all other elements
		subList := make([]int, 0, len(list)-1)
		for _, other := range list {
			if other != elem {
				subList = append(subList, other)
			}
		}
		// Recurse on the smaller list
		subPermutations := IntPermutations(subList)
		// Bolt on fixed elem to the beginning of each one
		for _, subPermutation := range subPermutations {
			newPermutation := make([]int, 0, len(list))
			newPermutation = append(newPermutation, elem)
			newPermutation = append(newPermutation, subPermutation...)
			permutations = append(permutations, newPermutation)
		}
	}
	return permutations
}
