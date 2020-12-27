package main

import (
	"aoc-go/year2019"
	"aoc-go/year2020"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Figure out which year we're running
	getParts := year2020.GetParts
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "2020":
			getParts = year2020.GetParts
		case "2019":
			getParts = year2019.GetParts
		}
	}
	if len(os.Args) < 3 {
		// Run all puzzles for a particular year
		for i := 1; i <= 25; i++ {
			partone, parttwo, filename := getParts(i, false)
			fmt.Printf("* Day %d: 1: %s, 2: %s \n", i, partone(filename), parttwo(filename))
		}
		return
	}
	// Run one day from a particular year
	daystr := os.Args[2]
	test := false
	if len(os.Args) > 2 {
		test = true
	}
	day, err := strconv.Atoi(daystr)
	if err != nil || day <= 0 || day > 25 {
		log.Fatal("Day number must be between 1 and 25 inclusive and... must actually be a number.")
	}
	partone, parttwo, filename := getParts(day, test)
	fmt.Println("Part one:", partone(filename))
	fmt.Println("Part two:", parttwo(filename))
}
