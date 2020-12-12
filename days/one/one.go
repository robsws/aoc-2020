package one

import (
	"aoc-2020/files"
	"fmt"
)

// PartOne - find two numbers that sum to 2020
func PartOne(filename string) {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			if p+q == 2020 {
				fmt.Println(p * q)
				return
			}
		}
		numbers = append(numbers, p)
	}
}

// PartTwo - find three numbers that sum to 2020
func PartTwo(filename string) {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			for _, r := range numbers {
				if p+q+r == 2020 {
					fmt.Println(p * q * r)
					return
				}
			}
		}
		numbers = append(numbers, p)
	}
}
