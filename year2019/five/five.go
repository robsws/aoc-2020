package five

import (
	"aoc-go/files"
	"aoc-go/intcode"
	"fmt"
)

// PartOne - intcode
func PartOne(filename string) string {
	lines := files.GetLines(filename)
	program := intcode.ParseProgram(lines[0])
	go program.Run()
	program.Input <- 1
	result := 0
	for out := range program.Output {
		result = out
	}
	return fmt.Sprint(result)
}

// PartTwo - intcode
func PartTwo(filename string) string {
	lines := files.GetLines(filename)
	program := intcode.ParseProgram(lines[0])
	go program.Run()
	program.Input <- 5
	result := <-program.Output
	return fmt.Sprint(result)
}
