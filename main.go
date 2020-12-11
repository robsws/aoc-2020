package main

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	filename := os.Args[1]
	dayelevenparttwo(filename)
}

func dayonepartone(filename string) {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			if p+q == 2020 {
				fmt.Println(p * q)
				return
			}
		}
		numbers = append(numbers, p)
	}
}

func dayoneparttwo(filename string) {
	numberStream := make(chan int)
	go files.StreamInts(filename, numberStream)
	numbers := make([]int, 0)
	for p := range numberStream {
		for _, q := range numbers {
			for _, r := range numbers {
				if p+q+r == 2020 {
					fmt.Println(p * q * r)
					return
				}
			}
		}
		numbers = append(numbers, p)
	}
}

func daytwopartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	re := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	valids := 0
	for line := range fileStream {
		submatches := re.FindStringSubmatch(line)
		min, _ := strconv.Atoi(submatches[1])
		max, _ := strconv.Atoi(submatches[2])
		rulechar := []rune(submatches[3])[0]
		password := submatches[4]
		charcount := 0
		for _, char := range password {
			if char == rulechar {
				charcount++
			}
		}
		if charcount >= min && charcount <= max {
			valids++
		}
	}
	fmt.Println(valids)
}

func daytwoparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	re := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	valids := 0
	for line := range fileStream {
		submatches := re.FindStringSubmatch(line)
		loc1, _ := strconv.Atoi(submatches[1])
		loc2, _ := strconv.Atoi(submatches[2])
		loc1--
		loc2--
		rulechar := []rune(submatches[3])[0]
		password := submatches[4]
		if (rune(password[loc1]) == rulechar || rune(password[loc2]) == rulechar) && !(rune(password[loc1]) == rulechar && rune(password[loc2]) == rulechar) {
			valids++
		}
	}
	fmt.Println(valids)
}

func daythreepartone(filename string) {
	resultStream := make(chan int)
	go countTrees(filename, 3, 1, resultStream)
	trees := <-resultStream
	fmt.Println(trees)
}

func daythreeparttwo(filename string) {
	resultStream := make(chan int)
	go countTrees(filename, 1, 1, resultStream)
	go countTrees(filename, 3, 1, resultStream)
	go countTrees(filename, 5, 1, resultStream)
	go countTrees(filename, 7, 1, resultStream)
	go countTrees(filename, 1, 2, resultStream)
	trees := 1
	for i := 0; i < 5; i++ {
		trees *= <-resultStream
	}
	fmt.Println(trees)
}

func countTrees(filename string, dx int, dy int, resultStream chan int) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	x := 0
	y := 0
	trees := 0
	for line := range fileStream {
		if y%dy == 0 {
			if line[x] == '#' {
				trees++
			}
			x = (x + dx) % len(line)
		}
		y++
	}
	resultStream <- trees
}

func dayfourpartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	valids := 0
	re := regexp.MustCompile("([a-z]{3}):[^ ]+")
	currentFields := make(map[string]bool)
	for line := range fileStream {
		if line == "" {
			if validatePassportFields(currentFields) {
				valids++
			}
			currentFields = make(map[string]bool)
		} else {
			fields := re.FindAllStringSubmatch(line, -1)
			for _, field := range fields {
				currentFields[field[1]] = true
			}
		}
	}
	if validatePassportFields(currentFields) {
		valids++
	}
	fmt.Println(valids)
}

func validatePassportFields(fields map[string]bool) bool {
	requiredFields := [7]string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
	valid := true
	for _, key := range requiredFields {
		_, exists := fields[key]
		if !exists {
			valid = false
			break
		}
	}
	return valid
}

func dayfourparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	valids := 0
	re := regexp.MustCompile("([a-z]{3}):([^ ]+)")
	currentFields := make(map[string]string)
	for line := range fileStream {
		if line == "" {
			if validatePassport(currentFields) {
				valids++
			}
			currentFields = make(map[string]string)
		} else {
			fields := re.FindAllStringSubmatch(line, -1)
			for _, field := range fields {
				currentFields[field[1]] = field[2]
			}
		}
	}
	if validatePassport(currentFields) {
		valids++
	}
	fmt.Println(valids)
}

func validatePassport(fields map[string]string) bool {
	requiredFields := map[string]func(string) bool{
		"byr": func(v string) bool { return validateYear(v, 1920, 2002) },
		"ecl": validateEyeColour,
		"eyr": func(v string) bool { return validateYear(v, 2020, 2030) },
		"hcl": validateHairColour,
		"hgt": validateHeight,
		"iyr": func(v string) bool { return validateYear(v, 2010, 2020) },
		"pid": validatePassportNumber,
	}
	valid := true
	for key, validator := range requiredFields {
		_, exists := fields[key]
		if !exists {
			valid = false
		} else {
			valid = validator(fields[key])
		}
		if !valid {
			break
		}
	}
	return valid
}

func validateYear(year string, min, max int) bool {
	yearI, _ := strconv.Atoi(year)
	return yearI <= max && yearI >= min
}

func validateHeight(h string) bool {
	re := regexp.MustCompile("([0-9]+)(cm|in)")
	if !re.MatchString(h) {
		return false
	}
	submatches := re.FindStringSubmatch(h)
	height, _ := strconv.Atoi(submatches[1])
	if submatches[2] == "cm" {
		return height >= 150 && height <= 193
	}
	return height >= 59 && height <= 76
}

func validateHairColour(c string) bool {
	re := regexp.MustCompile("^#[0-9a-f]{6}$")
	return re.MatchString(c)
}

func validateEyeColour(c string) bool {
	re := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
	return re.MatchString(c)
}

func validatePassportNumber(n string) bool {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(n)
}

func dayfivepartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	passes := make([]boardingPass, 0)
	for line := range fileStream {
		passes = append(passes, parseBoardingPass(line))
	}
	sort.Sort(sort.Reverse(bySeatID(passes)))
	fmt.Println(passes[0].CalcSeatID())
}

func dayfiveparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	passes := make([]boardingPass, 0)
	for line := range fileStream {
		passes = append(passes, parseBoardingPass(line))
	}
	sort.Sort(byLocation(passes))
	row := passes[0].row
	col := passes[0].col
	for _, pass := range passes {
		if pass.row != row || pass.col != col {
			mypass := boardingPass{loc: "", row: row, col: col}
			fmt.Println(mypass.CalcSeatID())
			break
		}
		col = (col + 1) % 8
		if col == 0 {
			row++
		}
	}
}

type boardingPass struct {
	loc string
	row int
	col int
}

func (a boardingPass) CalcSeatID() int {
	return a.row*8 + a.col
}

type bySeatID []boardingPass

func (l bySeatID) Len() int           { return len(l) }
func (l bySeatID) Less(i, j int) bool { return l[i].CalcSeatID() < l[j].CalcSeatID() }
func (l bySeatID) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

type byLocation []boardingPass

func (l byLocation) Len() int { return len(l) }
func (l byLocation) Less(i, j int) bool {
	return l[i].row < l[j].row || (l[i].row == l[j].row && l[i].col < l[j].col)
}
func (l byLocation) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func parseBoardingPass(loc string) boardingPass {
	row := parseRowLocation(loc[:7])
	col := parseColumnLocation(loc[7:])
	return boardingPass{loc: loc, row: row, col: col}
}

func parseRowLocation(rowstr string) int {
	if len(rowstr) != 7 {
		log.Fatalf("Row location (%s) string must be 7 digits", rowstr)
	}
	row := 0
	poweroftwo := 1
	for i := 6; i >= 0; i-- {
		if rowstr[i] == 'B' {
			row += poweroftwo
		}
		poweroftwo *= 2
	}
	return row
}

func parseColumnLocation(colstr string) int {
	if len(colstr) != 3 {
		log.Fatalf("Column location (%s) string must be 3 digits", colstr)
	}
	col := 0
	poweroftwo := 1
	for i := 2; i >= 0; i-- {
		if colstr[i] == 'R' {
			col += poweroftwo
		}
		poweroftwo *= 2
	}
	return col
}

func daysixpartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	total := 0
	group := make(map[rune]int)
	for line := range fileStream {
		if line == "" {
			for _, n := range group {
				total += n
			}
			group = make(map[rune]int)
			continue
		}
		for i := 0; i < len(line); i++ {
			group[rune(line[i])] = 1
		}
	}
	for _, n := range group {
		total += n
	}
	fmt.Println(total)
}

func daysixparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	total := 0
	group := make(map[rune]int)
	groupsize := 0
	for line := range fileStream {
		if line == "" {
			for _, n := range group {
				if n == groupsize {
					total++
				}
			}
			group = make(map[rune]int)
			groupsize = 0
			continue
		}
		groupsize++
		for i := 0; i < len(line); i++ {
			group[rune(line[i])]++
		}
	}
	for _, n := range group {
		if n == groupsize {
			total++
		}
	}
	fmt.Println(total)
}

func daysevenpartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	containerRe := regexp.MustCompile("^[a-z]+ [a-z]+")
	containedRe := regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+)")
	containedBy := make(map[string][]string)
	for line := range fileStream {
		container := containerRe.FindString(line)
		contained := containedRe.FindAllStringSubmatch(line, -1)
		for _, submatch := range contained {
			containedBy[submatch[2]] = append(containedBy[submatch[2]], container)
		}
	}
	goldContainers := getContainers(containedBy, "shiny gold", true)
	fmt.Println(goldContainers.Len())
}

func getContainers(containedBy map[string][]string, colour string, top bool) utils.Set {
	containers := containedBy[colour]
	colours := utils.MakeSet()
	for _, container := range containers {
		colours.Union(getContainers(containedBy, container, false))
	}
	if !top {
		colours.Add(colour)
	}
	return colours
}

func daysevenparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	containerRe := regexp.MustCompile("^[a-z]+ [a-z]+")
	containedRe := regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+)")
	containerOf := make(map[string][]container)
	for line := range fileStream {
		containerColour := containerRe.FindString(line)
		containedColours := containedRe.FindAllStringSubmatch(line, -1)
		for _, submatch := range containedColours {
			amount, _ := strconv.Atoi(submatch[1])
			containerOf[containerColour] = append(containerOf[containerColour], container{Colour: submatch[2], Amount: amount})
		}
	}
	goldContains := countContained(containerOf, container{Colour: "shiny gold", Amount: 1}, true)
	fmt.Println(goldContains)
}

type container struct {
	Colour string
	Amount int
}

func countContained(containerOf map[string][]container, c container, top bool) int {
	total := 0
	for _, innerContainer := range containerOf[c.Colour] {
		total += c.Amount * countContained(containerOf, innerContainer, false)
	}
	if !top {
		total += c.Amount
	}
	return total
}

func dayeightpartone(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	instructionRe := regexp.MustCompile("^([a-z]{3}) ((?:\\+|-)[0-9]+)$")
	program := make([]instruction, 0)
	for line := range fileStream {
		submatches := instructionRe.FindStringSubmatch(line)
		valueI, _ := strconv.Atoi(submatches[2])
		program = append(program, instruction{Command: submatches[1], Value: valueI})
	}
	acc, _ := runProgram(program)
	fmt.Println(acc)
}

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

func dayeightparttwo(filename string) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	instructionRe := regexp.MustCompile("^([a-z]{3}) ((?:\\+|-)[0-9]+)$")
	program := make([]instruction, 0)
	for line := range fileStream {
		submatches := instructionRe.FindStringSubmatch(line)
		valueI, _ := strconv.Atoi(submatches[2])
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
				fmt.Println(acc)
				break
			}
		}
	}
}

func dayninepartone(filename string, preamble int) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if ok {
		fmt.Println(nonsum)
	} else {
		fmt.Println("No non-summing number found.")
	}
}

func findNonSumNumber(numbers []int, preamble int) (int, bool) {
	for i := 0; i < len(numbers)-preamble; i++ {
		sums := utils.MakeSet()
		for j := 0; j < preamble; j++ {
			for k := 0; k < preamble; k++ {
				if k != j {
					sums.Add(numbers[i+k] + numbers[i+j])
				}
			}
		}
		if !sums.Contains(numbers[i+preamble]) {
			return numbers[i+preamble], true
		}
	}
	return -1, false
}

func daynineparttwo(filename string, preamble int) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	numbers := make([]int, 0)
	for n := range fileStream {
		numbers = append(numbers, n)
	}
	nonsum, ok := findNonSumNumber(numbers, preamble)
	if !ok {
		log.Fatal("No non-summing number found.")
	}
	start, end, ok := findContiguousSumRange(numbers, nonsum)
	if !ok {
		log.Fatalf("Non contiguous block found that sums to first nonsum value %d", nonsum)
	}
	sumrange := numbers[start : end+1]
	sort.Ints(sumrange)
	fmt.Println(sumrange[0] + sumrange[len(sumrange)-1])
}

func findContiguousSumRange(numbers []int, n int) (int, int, bool) {
	for i := 0; i < len(numbers)-1; i++ {
		total := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			total += numbers[j]
			if total == n {
				return i, j, true
			} else if total > n {
				break
			}
		}
	}
	return -1, -1, false
}

func daytenpartone(filename string) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	adapters := make([]int, 0)
	for n := range fileStream {
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	var diffs [4]int
	diffs[adapters[0]]++
	for i := 0; i < len(adapters)-1; i++ {
		diffs[adapters[i+1]-adapters[i]]++
	}
	diffs[3]++
	fmt.Println(diffs[1] * diffs[3])
}

func daytenparttwo(filename string) {
	fileStream := make(chan int)
	go files.StreamInts(filename, fileStream)
	adapters := make([]int, 0)
	adapters = append(adapters, 0)
	for n := range fileStream {
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	permutations := 1
	inarow := 0
	for i := 1; i < len(adapters)-1; i++ {
		if adapters[i+1]-adapters[i-1] <= 2 {
			inarow++
		} else {
			if inarow == 0 {
				continue
			}
			if inarow == 1 {
				permutations *= 2
			} else {
				permutations *= utils.PowInt(2, inarow) - (utils.PowInt(2, inarow-2) - 1)
			}
			inarow = 0
		}
	}
	fmt.Println(permutations)
}

func dayelevenpartone(filename string) {
	lines := files.GetLines(filename)
	height := len(lines)
	width := len(lines[0])
	grid := initGrid(width, height)
	// Initialise the grid with the input
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = rune(lines[y][x])
		}
	}
	// Simulate rounds until grid doesn't change
	prevgrid := initGrid(width, height)
	for !gridsAreEqual(grid, prevgrid) {
		prevgrid = grid
		grid = simulateSeats(grid, false)
	}
	// Count occupied seats
	fmt.Println(countTotalOccupied(grid))
}

func dayelevenparttwo(filename string) {
	lines := files.GetLines(filename)
	height := len(lines)
	width := len(lines[0])
	grid := initGrid(width, height)
	// Initialise the grid with the input
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = rune(lines[y][x])
		}
	}
	// Simulate rounds until grid doesn't change
	prevgrid := initGrid(width, height)
	for !gridsAreEqual(grid, prevgrid) {
		prevgrid = grid
		grid = simulateSeats(grid, true)
	}
	// Count occupied seats
	fmt.Println(countTotalOccupied(grid))
}

func initGrid(width, height int) [][]rune {
	grid := make([][]rune, width)
	for i := 0; i < width; i++ {
		grid[i] = make([]rune, height)
		for j := 0; j < height; j++ {
			grid[i][j] = 'x'
		}
	}
	return grid
}

func gridsAreEqual(grid1, grid2 [][]rune) bool {
	width := len(grid1)
	height := len(grid1[0])
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid1[x][y] != grid2[x][y] {
				return false
			}
		}
	}
	return true
}

func simulateSeats(grid [][]rune, visible bool) [][]rune {
	width := len(grid)
	height := len(grid[0])
	newgrid := initGrid(width, height)
	// For every cell in the grid
	for x, col := range grid {
		for y, seat := range col {
			// Skip if floor cell
			if seat == '.' {
				newgrid[x][y] = '.'
				continue
			}
			var occupied int
			if visible {
				occupied = countVisibleSeatsOccupied(grid, x, y)
			} else {
				occupied = countSurroundingSeatsOccupied(grid, x, y)
			}
			if seat == 'L' && occupied == 0 {
				newgrid[x][y] = '#'
				continue
			}
			maxOccupied := 4
			if visible {
				maxOccupied = 5
			}
			if seat == '#' && occupied >= maxOccupied {
				newgrid[x][y] = 'L'
				continue
			}
			newgrid[x][y] = grid[x][y]
		}
	}
	return newgrid
}

func countSurroundingSeatsOccupied(grid [][]rune, x int, y int) int {
	// Count all occupied seats surrounding position
	width := len(grid)
	height := len(grid[0])
	occupied := 0
	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= width {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= height {
				continue
			}
			if !(i == x && j == y) && grid[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func countVisibleSeatsOccupied(grid [][]rune, x int, y int) int {
	// Count all occupied seats surrounding position
	width := len(grid)
	height := len(grid[0])
	occupied := 0
	for xOffset := -1; xOffset <= 1; xOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			i, j := x+xOffset, y+yOffset
			end := false
			if i < 0 || i >= width || j < 0 || j >= height {
				end = true
			}
			for !end && grid[i][j] == '.' {
				i += xOffset
				j += yOffset
				if i < 0 || i >= width || j < 0 || j >= height {
					end = true
				}
			}
			if end {
				continue
			}
			if !(i == x && j == y) && grid[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func countTotalOccupied(grid [][]rune) int {
	width := len(grid)
	height := len(grid[0])
	total := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[x][y] == '#' {
				total++
			}
		}
	}
	return total
}
