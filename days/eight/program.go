package eight

func runProgram(program []instruction) (int, bool) {
	visited := make([]int, len(program))
	pc := 0
	acc := 0
	terminated := false
	for {
		if pc == len(program) {
			terminated = true
			break
		}
		i := program[pc]
		visited[pc]++
		if visited[pc] == 2 {
			terminated = false
			break
		}
		switch i.Command {
		case "acc":
			acc += i.Value
		case "jmp":
			pc += i.Value
			continue
		}
		pc++
	}
	return acc, terminated
}

type instruction struct {
	Command string
	Value   int
}
