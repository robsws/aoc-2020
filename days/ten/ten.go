package ten

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"sort"
)

// PartOne - get number of 1 and 3 length diffs in joltage
func PartOne(filename string) {
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
	fmt.Println(diffs[1] * diffs[3])
}

// PartTwo - get number of adapter permutations
func PartTwo(filename string) {
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
				permutations *= utils.PowInt(2, inarow) - (utils.PowInt(2, inarow-2) - 1)
			}
			inarow = 0
		}
	}
	fmt.Println(permutations)
}
