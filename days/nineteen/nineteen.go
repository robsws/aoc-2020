package nineteen

import (
	"aoc-2020/files"
	"fmt"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	regex, err := parseRules(fileStream, false)
	if err != nil {
		return fmt.Sprint("Could not parse rules. ", err)
	}
	valid := 0
	for line := range fileStream {
		if regex.MatchString(line) {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	regex, err := parseRules(fileStream, true)
	if err != nil {
		return fmt.Sprint("Could not parse rules. ", err)
	}
	valid := 0
	for line := range fileStream {
		if regex.MatchString(line) {
			valid++
		}
	}
	return fmt.Sprint(valid)
}
