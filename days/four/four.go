package four

import (
	"aoc-2020/files"
	"fmt"
	"regexp"
)

// PartOne - count valid passports with basic rules
func PartOne(filename string) string {
	return fmt.Sprint(countValidPassports(filename, validatePassportFields))
}

// PartTwo - count valid passports with complex rules
func PartTwo(filename string) string {
	return fmt.Sprint(countValidPassports(filename, validatePassport))
}

func countValidPassports(filename string, checkValid func(map[string]string) bool) int {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	valids := 0
	re := regexp.MustCompile("([a-z]{3}):([^ ]+)")
	currentFields := make(map[string]string)
	for line := range fileStream {
		if line == "" {
			if checkValid(currentFields) {
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
	if checkValid(currentFields) {
		valids++
	}
	return valids
}
