package two

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
)

// PartOne - get number of valid passports with simple rules
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	re := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	valids := 0
	for line := range fileStream {
		submatches := re.FindStringSubmatch(line)
		min := utils.MustAtoi(submatches[1])
		max := utils.MustAtoi(submatches[2])
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
	return fmt.Sprint(valids)
}

// PartTwo - get number of valid passports with complex rules
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	re := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)$")
	valids := 0
	for line := range fileStream {
		submatches := re.FindStringSubmatch(line)
		loc1 := utils.MustAtoi(submatches[1])
		loc2 := utils.MustAtoi(submatches[2])
		loc1--
		loc2--
		rulechar := []rune(submatches[3])[0]
		password := submatches[4]
		if (rune(password[loc1]) == rulechar || rune(password[loc2]) == rulechar) && !(rune(password[loc1]) == rulechar && rune(password[loc2]) == rulechar) {
			valids++
		}
	}
	return fmt.Sprint(valids)
}
