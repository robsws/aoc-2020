package five

import (
	"aoc-go/files"
	"aoc-go/intcode"
	"fmt"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	lines := files.GetLines(filename)
	program := intcode.ParseProgram(lines[0])
	go program.Run()
	program.Input <- 1
	result := <-program.Output
	return fmt.Sprint(result)
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	return "Not yet implemented"
}
