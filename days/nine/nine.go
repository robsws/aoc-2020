package nine

import (
	"aoc-2020/files"
	"fmt"
	"log"
	"sort"
)

const preamble = 25

// PartOne - find number that isn't sum of previous
func PartOne(filename string) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if ok {
		fmt.Println(nonsum)
	} else {
		fmt.Println("No non-summing number found.")
	}
}

// PartTwo - find encryption weakness of XMAS
func PartTwo(filename string) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if !ok {
		log.Fatal("No non-summing number found.")
	}
	start, end, ok := findContiguousSumRange(numbers, nonsum)
	if !ok {
		log.Fatalf("Non contiguous block found that sums to first nonsum value %d", nonsum)
	}
	sumrange := numbers[start : end+1]
	sort.Ints(sumrange)
	fmt.Println(sumrange[0] + sumrange[len(sumrange)-1])
}
