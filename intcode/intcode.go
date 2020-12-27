package intcode

import (
	"aoc-go/utils"
	"log"
	"strings"
)

const opAdd = 1
const opMultiply = 2
const opTerminate = 99

// Program - encapsulates an intcode program
type Program struct {
	memory     []int
	counter    int
	Terminated bool
}

// ParseProgram - parse an intcode program from a comma separated list of ints
func ParseProgram(csv string) Program {
	initMemoryStr := strings.Split(csv, ",")
	initMemory := make([]int, len(initMemoryStr))
	for i, intStr := range initMemoryStr {
		initMemory[i] = utils.MustAtoi(intStr)
	}
	return MakeProgram(initMemory)
}

// MakeProgram - initialize a new intcode program given initial state of memory
func MakeProgram(initMemory []int) Program {
	return Program{initMemory, 0, false}
}

// Run - run an intcode program to completion
func (program *Program) Run() {
	for !program.Terminated {
		program.Step()
	}
}

// Step - run one execution cycle of an intcode program
func (program *Program) Step() {
	switch program.memory[program.counter] {
	case opAdd:
		operandAddr1 := program.memory[program.counter+1]
		operandAddr2 := program.memory[program.counter+2]
		resultAddr := program.memory[program.counter+3]
		program.add(operandAddr1, operandAddr2, resultAddr)
		program.counter += 4
	case opMultiply:
		operandAddr1 := program.memory[program.counter+1]
		operandAddr2 := program.memory[program.counter+2]
		resultAddr := program.memory[program.counter+3]
		program.multiply(operandAddr1, operandAddr2, resultAddr)
		program.counter += 4
	case opTerminate:
		program.Terminated = true
	default:
		log.Fatal("Invalid opcode ", program.memory[program.counter])
	}
}

func (program *Program) add(aAddr int, bAddr int, resultAddr int) {
	program.memory[resultAddr] = program.memory[aAddr] + program.memory[bAddr]
}

func (program *Program) multiply(aAddr int, bAddr int, resultAddr int) {
	program.memory[resultAddr] = program.memory[aAddr] * program.memory[bAddr]
}

// Get - get a value from the program's memory
func (program *Program) Get(addr int) int {
	return program.memory[addr]
}

// Set - set a value in memory
func (program *Program) Set(addr int, value int) {
	program.memory[addr] = value
}

// Copy - copy an intcode program
func (program *Program) Copy() Program {
	newMemory := make([]int, len(program.memory))
	for i, val := range program.memory {
		newMemory[i] = val
	}
	return Program{newMemory, program.counter, program.Terminated}
}
