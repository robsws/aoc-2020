package fifteen

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"strings"
)

// PartOne - elf number game round 2020
func PartOne(filename string) string {
	return fmt.Sprint(elfNumberGame(filename, 2020))
}

// PartTwo - elf number game round 30 million
func PartTwo(filename string) string {
	return fmt.Sprint(elfNumberGame(filename, 30000000))
}

func elfNumberGame(filename string, round int) int {
	inputs := strings.Split(files.GetLines(filename)[0], ",")
	lastSpoken := make(map[int]int)
	roundResult := -1
	for i := 0; i < round; i++ {
		if i < len(inputs) {
			n := utils.MustAtoi(inputs[i])
			lastSpoken[n] = i
			roundResult = n
			continue
		}
		nLastSpoken, ok := lastSpoken[roundResult]
		lastSpoken[roundResult] = i - 1
		if !ok {
			roundResult = 0
		} else {
			roundResult = i - 1 - nLastSpoken
		}
	}
	return roundResult
}
