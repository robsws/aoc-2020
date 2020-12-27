package nine

import (
	"aoc-go/files"
	"fmt"
	"sort"
)

// PartOne - find number that isn't sum of previous
func PartOne(filename string, test bool) string {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	preamble := 25
	if test {
		preamble = 5
	}
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if ok {
		return fmt.Sprint(nonsum)
	}
	return "No non-summing number found."
}

// PartTwo - find encryption weakness of XMAS
func PartTwo(filename string, test bool) string {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	preamble := 25
	if test {
		preamble = 5
	}
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if !ok {
		return "No non-summing number found."
	}
	start, end, ok := findContiguousSumRange(numbers, nonsum)
	if !ok {
		return fmt.Sprintf("Non contiguous block found that sums to first nonsum value %d", nonsum)
	}
	sumrange := numbers[start : end+1]
	sort.Ints(sumrange)
	return fmt.Sprint(sumrange[0] + sumrange[len(sumrange)-1])
}
