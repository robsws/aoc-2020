package seven

import (
	"aoc-go/files"
	"aoc-go/intcode"
	"aoc-go/utils"
	"fmt"
)

// PartOne - max thruster output signal
func PartOne(filename string) string {
	phases := []int{0, 1, 2, 3, 4}
	return fmt.Sprint(getMaxThrusterSignal(filename, phases, runAmplifiers))
}

// PartTwo - max looped thruster output signal
func PartTwo(filename string) string {
	phases := []int{5, 6, 7, 8, 9}
	return fmt.Sprint(getMaxThrusterSignal(filename, phases, runLoopedAmplifiers))
}

func getMaxThrusterSignal(filename string, phases []int, runAmpsFunc func(intcode.Program, []int) int) int {
	lines := files.GetLines(filename)
	program := intcode.ParseProgram(lines[0])
	phasePermutations := utils.IntPermutations(phases)
	maxSignal := -1
	for _, phasePermutation := range phasePermutations {
		signal := runAmpsFunc(program, phasePermutation)
		if signal > maxSignal {
			maxSignal = signal
		}
	}
	return maxSignal
}

func runAmplifiers(program intcode.Program, phases []int) int {
	var programs [5]intcode.Program
	// Initialise programs
	for i := 0; i < 5; i++ {
		programs[i] = program.Copy()
		go programs[i].Run()
		programs[i].Input <- phases[i]
	}
	// Feed I/O through the amplifier chain
	programs[0].Input <- 0
	for i := 1; i < 5; i++ {
		programs[i].Input <- <-programs[i-1].Output
	}
	return <-programs[4].Output
}

func runLoopedAmplifiers(program intcode.Program, phases []int) int {
	var programs [5]intcode.Program
	// Initialise programs
	for i := 0; i < 5; i++ {
		programs[i] = program.Copy()
		go programs[i].Run()
		programs[i].Input <- phases[i]
	}
	// Feed I/O through the looped amplifier chain
	firstRound := true
	for !programs[0].Finished {
		if firstRound {
			programs[0].Input <- 0
			firstRound = false
		} else {
			programs[0].Input <- <-programs[4].Output
		}
		for i := 1; i < 5; i++ {
			programs[i].Input <- <-programs[i-1].Output
		}
	}
	return <-programs[4].Output
}
