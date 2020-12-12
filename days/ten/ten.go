package ten

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"sort"
)

// PartOne - get number of 1 and 3 length diffs in joltage
func PartOne(filename string) string {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	adapters := make([]int, 0)
	for n := range fileStream {
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	var diffs [4]int
	diffs[adapters[0]]++
	for i := 0; i < len(adapters)-1; i++ {
		diffs[adapters[i+1]-adapters[i]]++
	}
	diffs[3]++
	return fmt.Sprint(diffs[1] * diffs[3])
}

// PartTwo - get number of adapter permutations
func PartTwo(filename string) string {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	adapters := make([]int, 0)
	adapters = append(adapters, 0)
	for n := range fileStream {
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	permutations := 1
	inarow := 0
	for i := 1; i < len(adapters)-1; i++ {
		if adapters[i+1]-adapters[i-1] <= 2 {
			inarow++
		} else {
			if inarow == 0 {
				continue
			}
			if inarow == 1 {
				permutations *= 2
			} else {
				/* For runs of adapters that are removable (1J away on both sides)
				   the number of permutations of those adapters is 2^n, where n is the
				   length of the run. However, some of those permutations are invalid,
				   because removing at least 3 in a row causes digits either side to be
				   more than 3J away from each other. */
				permutations *= utils.PowInt(2, inarow) - calcExceptions(inarow)
			}
			inarow = 0
		}
	}
	return fmt.Sprint(permutations)
}

func calcExceptions(inarow int) int {
	/* The amount of exceptions is the amount of permutations
	   that include a series of 3 or more removed adapters (a 'window').
	   The number of such permutations is the sum of all the
	   positions of a window of each size multiplied by the
	   amount of valid permutations of all the adapters around the
	   window in each position, which is 2 ^ number of adapters
	   that can be changed (adapters immediately left or right cannot
		be changed or they'd change the size of the window) */
	if inarow < 3 {
		return 0
	}
	exceptions := 1
	for window := 3; window < inarow; window++ {
		for pos := 0; pos < inarow-window; pos++ {
			leftChangeable := utils.PowInt(2, utils.MaxInt(0, pos-1))
			rightChangeable := utils.PowInt(2, utils.MaxInt(0, inarow-(pos+window)-1))
			exceptions += leftChangeable * rightChangeable
		}
	}
	return exceptions
}
