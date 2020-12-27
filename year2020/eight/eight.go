package eight

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
)

// PartOne - run program and output accumulator
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	instructionRe := regexp.MustCompile("^([a-z]{3}) ((?:\\+|-)[0-9]+)$")
	program := make([]instruction, 0)
	for line := range fileStream {
		submatches := instructionRe.FindStringSubmatch(line)
		valueI := utils.MustAtoi(submatches[2])
		program = append(program, instruction{Command: submatches[1], Value: valueI})
	}
	acc, _ := runProgram(program)
	return fmt.Sprint(acc)
}

// PartTwo - amend program so that it terminates, and output accumulator
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	instructionRe := regexp.MustCompile("^([a-z]{3}) ((?:\\+|-)[0-9]+)$")
	program := make([]instruction, 0)
	for line := range fileStream {
		submatches := instructionRe.FindStringSubmatch(line)
		valueI := utils.MustAtoi(submatches[2])
		program = append(program, instruction{Command: submatches[1], Value: valueI})
	}
	for i, op := range program {
		if op.Command == "nop" || op.Command == "jmp" {
			newprog := make([]instruction, len(program))
			copy(newprog, program)
			newop := "jmp"
			if op.Command == "jmp" {
				newop = "nop"
			}
			newprog[i].Command = newop
			acc, ok := runProgram(newprog)
			if ok {
				return fmt.Sprint(acc)
			}
		}
	}
	return "No terminating program found"
}
