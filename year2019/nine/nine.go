package nine

import (
	"aoc-go/files"
	"aoc-go/intcode"
	"fmt"
)

// PartOne - complete intcode check
func PartOne(filename string) string {
	code := files.GetLines(filename)[0]
	program := intcode.ParseProgram(code)
	go program.Run()
	program.Input <- 1
	return fmt.Sprint(<-program.Output)
}

// PartTwo - boost sensors
func PartTwo(filename string) string {
	code := files.GetLines(filename)[0]
	program := intcode.ParseProgram(code)
	go program.Run()
	program.Input <- 2
	return fmt.Sprint(<-program.Output)
}
