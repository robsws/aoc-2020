package four

import (
	"aoc-2020/files"
	"fmt"
	"regexp"
)

// PartOne - count valid passports with basic rules
func PartOne(filename string) {
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

// PartTwo - count valid passports with complex rules
func PartTwo(filename string) {
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
