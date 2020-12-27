package twentyfive

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	lines := files.GetLines(filename)
	cardPub := utils.MustAtoi(lines[0])
	doorPub := utils.MustAtoi(lines[1])
	cardLoop := 0
	doorLoop := 0
	subject := 7
	value := 1
	for i := 1; i < 1000000; i++ {
		value = (value * subject) % 20201227
		if value == cardPub {
			cardLoop = i
		}
		if value == doorPub {
			doorLoop = i
		}
		if cardLoop > 0 && doorLoop > 0 {
			break
		}
	}
	subject = cardPub
	value = 1
	for i := 0; i < doorLoop; i++ {
		value = (value * subject) % 20201227
	}
	return fmt.Sprint(value)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	return "Merry Christmas 🎄"
}
