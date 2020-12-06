package main

import (
	"aoc-2020/files"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	filename := os.Args[1]
	daysixparttwo(filename)
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
