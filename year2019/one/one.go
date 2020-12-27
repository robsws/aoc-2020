package one

import (
	"aoc-go/files"
	"fmt"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	total := 0
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	for mass := range fileStream {
		total += calcFuel(mass)
	}
	return fmt.Sprint(total)
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	total := 0
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	for mass := range fileStream {
		extraFuel := calcFuel(mass)
		for extraFuel > 0 {
			total += extraFuel
			extraFuel = calcFuel(extraFuel)
		}
	}
	return fmt.Sprint(total)
}

func calcFuel(mass int) int {
	return int(mass/3) - 2
}
