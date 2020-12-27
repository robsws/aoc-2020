package one

import (
	"aoc-go/files"
	"fmt"
)

// PartOne - find two numbers that sum to 2020
func PartOne(filename string) string {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			if p+q == 2020 {
				return fmt.Sprint(p * q)
			}
		}
		numbers = append(numbers, p)
	}
	return "No sum to 2020 found."
}

// PartTwo - find three numbers that sum to 2020
func PartTwo(filename string) string {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			for _, r := range numbers {
				if p+q+r == 2020 {
					return fmt.Sprint(p * q * r)
				}
			}
		}
		numbers = append(numbers, p)
	}
	return "No sum to 2020 found."
}
