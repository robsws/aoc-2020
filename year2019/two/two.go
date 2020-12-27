package two

import (
	"aoc-go/files"
	"aoc-go/intcode"
	"fmt"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	csv := files.GetLines(filename)[0]
	program := intcode.ParseProgram(csv)
	program.Set(1, 12)
	program.Set(2, 2)
	program.Run()
	return fmt.Sprint(program.Get(0))
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	csv := files.GetLines(filename)[0]
	template := intcode.ParseProgram(csv)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := template.Copy()
			program.Set(1, noun)
			program.Set(2, verb)
			program.Run()
			if program.Get(0) == 19690720 {
				return fmt.Sprint(100*noun + verb)
			}
		}
	}
	return "No argument combination found that results in target."
}
