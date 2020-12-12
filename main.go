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
		log.Fatal("Not enough args")
	}
	daystr := os.Args[1]
	test := false
	if len(os.Args) > 2 {
		test = true
	}
	day, _ := strconv.Atoi(daystr)
	partone, parttwo, filename := days.GetParts(day, test)
	fmt.Println("Part one:")
	partone(filename)
	fmt.Println("Part two:")
	parttwo(filename)
}
