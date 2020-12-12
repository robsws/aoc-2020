package main

import (
	"aoc-2020/days"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		for i := 1; i <= 25; i++ {
			partone, parttwo, filename := days.GetParts(i, false)
			fmt.Printf("* Day %d: 1: %s, 2: %s \n", i, partone(filename), parttwo(filename))
		}
		return
	}
	daystr := os.Args[1]
	test := false
	if len(os.Args) > 2 {
		test = true
	}
	day, err := strconv.Atoi(daystr)
	if err != nil || day <= 0 || day > 25 {
		log.Fatal("Day number must be between 1 and 25 inclusive and... must actually be a number.")
	}
	partone, parttwo, filename := days.GetParts(day, test)
	fmt.Println("Part one:", partone(filename))
	fmt.Println("Part two:", parttwo(filename))
}
