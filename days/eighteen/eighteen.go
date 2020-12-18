package eighteen

import (
	"aoc-2020/files"
	"fmt"
	"strings"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	sum := 0
	for line := range fileStream {
		tokens := strings.Split(line, "")
		result := evalLeftToRight(tokens)
		sum += result
	}
	return fmt.Sprint(sum)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	sum := 0
	for line := range fileStream {
		tokens := strings.Split(line, "")
		result := evalAdditionFirst(tokens)
		sum += result
	}
	return fmt.Sprint(sum)
}
