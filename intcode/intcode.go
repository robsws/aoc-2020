package intcode

import (
	"aoc-go/utils"
	"fmt"
	"log"
	"strings"
)

const opAdd = 1
const opMultiply = 2
const opInput = 3
const opOutput = 4
const opJumpIfTrue = 5
const opJumpIfFalse = 6
const opLessThan = 7
const opEquals = 8
const opTerminate = 99

var paramCount = map[int]int{
	opAdd:         3,
	opMultiply:    3,
	opInput:       1,
	opOutput:      1,
	opJumpIfTrue:  2,
	opJumpIfFalse: 2,
	opLessThan:    3,
	opEquals:      3,
	opTerminate:   0,
}

const positionMode = 0
const immediateMode = 1

// Program - encapsulates an intcode program
type Program struct {
	memory     []int
	counter    int
	terminated bool
	Input      chan int
	Output     chan int
	Debug      bool
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
	return Program{initMemory, 0, false, make(chan int), make(chan int), false}
}

// Run - run an intcode program to completion
func (program *Program) Run() {
	for !program.terminated {
		program.Step()
	}
	close(program.Output)
}

// Step - run one execution cycle of an intcode program
func (program *Program) Step() {
	opCode, params := program.processCommand()
	switch opCode {
	case opAdd:
		program.add(*params[0], *params[1], params[2])
	case opMultiply:
		program.multiply(*params[0], *params[1], params[2])
	case opInput:
		program.input(params[0])
	case opOutput:
		program.output(*params[0])
	case opJumpIfTrue:
		program.jumpIfTrue(*params[0], *params[1])
	case opJumpIfFalse:
		program.jumpIfFalse(*params[0], *params[1])
	case opLessThan:
		program.lessThan(*params[0], *params[1], params[2])
	case opEquals:
		program.equals(*params[0], *params[1], params[2])
	case opTerminate:
		program.terminated = true
	default:
		log.Fatal("Invalid opcode ", program.memory[program.counter])
	}
}

func (program *Program) add(a int, b int, addr *int) {
	*addr = a + b
	program.counter += 4
}

func (program *Program) multiply(a int, b int, addr *int) {
	*addr = a * b
	program.counter += 4
}

func (program *Program) input(addr *int) {
	*addr = <-program.Input
	program.counter += 2
}

func (program *Program) output(value int) {
	program.Output <- value
	program.counter += 2
}

func (program *Program) jumpIfTrue(value int, addr int) {
	if value != 0 {
		program.counter = addr
	} else {
		program.counter += 3
	}
}

func (program *Program) jumpIfFalse(value int, addr int) {
	if value == 0 {
		program.counter = addr
	} else {
		program.counter += 3
	}
}

func (program *Program) lessThan(a int, b int, addr *int) {
	if a < b {
		*addr = 1
	} else {
		*addr = 0
	}
	program.counter += 4
}

func (program *Program) equals(a int, b int, addr *int) {
	if a == b {
		*addr = 1
	} else {
		*addr = 0
	}
	program.counter += 4
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
	return Program{newMemory, program.counter, program.terminated, make(chan int), make(chan int), program.Debug}
}

func (program *Program) processCommand() (opCode int, params []*int) {
	instruction := program.memory[program.counter]
	opCode, paramModes := parseInstruction(instruction)
	params = make([]*int, len(paramModes))
	for i, mode := range paramModes {
		paramValue := program.memory[program.counter+i+1]
		switch mode {
		case immediateMode:
			params[i] = &paramValue
		case positionMode:
			params[i] = &program.memory[paramValue]
		}
	}
	if program.Debug {
		program.debugCommand(opCode, paramModes, program.memory[program.counter+1:program.counter+1+paramCount[opCode]])
	}
	return
}

func parseInstruction(instruction int) (opCode int, paramModes []int) {
	opCode = instruction % 100
	opParamCount, ok := paramCount[opCode]
	if !ok {
		log.Fatal("Parameter count could not be obtained for ", opCode)
	}
	paramModes = make([]int, opParamCount)
	instruction = (instruction - opCode) / 100
	for i := 0; i < opParamCount; i++ {
		paramModes[i] = instruction % 10
		instruction = (instruction - paramModes[i]) / 10
	}
	return
}

func (program *Program) debugCommand(opCode int, paramModes []int, rawParams []int) {
	opCodeNames := map[int]string{
		opAdd:         "Add",
		opMultiply:    "Multiply",
		opInput:       "Input",
		opOutput:      "Output",
		opJumpIfTrue:  "Jump-If-True",
		opJumpIfFalse: "Jump-If-False",
		opLessThan:    "Less-Than",
		opEquals:      "Equals",
		opTerminate:   "Terminate",
	}
	commandStr := opCodeNames[opCode]
	for i := 0; i < len(rawParams); i++ {
		if paramModes[i] == positionMode {
			commandStr += " pos{" + fmt.Sprint(rawParams[i]) + "}(" + fmt.Sprint(program.memory[rawParams[i]]) + ")"
		} else if paramModes[i] == immediateMode {
			commandStr += " " + fmt.Sprint(rawParams[i])
		} else {
			commandStr += " unk{" + fmt.Sprint(rawParams[i]) + "}"
		}
	}
	fmt.Println(commandStr)
}
