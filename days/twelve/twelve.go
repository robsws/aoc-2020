package twelve

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"regexp"
)

// PartOne - get manhattan distance travelled by ferry
func PartOne(filename string) string {
	return fmt.Sprint(simulateFerry(filename, false))
}

// PartTwo - get manhattan distance travelled by ferry with waypoint
func PartTwo(filename string) string {
	return fmt.Sprint(simulateFerry(filename, true))
}

func simulateFerry(filename string, usesWaypoint bool) int {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	ferry := makeFerry(usesWaypoint)
	re := regexp.MustCompile("([A-Z])([0-9]+)")
	for line := range fileStream {
		parts := re.FindStringSubmatch(line)
		command := rune(parts[1][0])
		value := utils.MustAtoi(parts[2])
		ferry.TakeCommand(command, value)
	}
	return ferry.ManhattanDistance()
}
