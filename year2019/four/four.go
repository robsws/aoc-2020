package four

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	lower, upper := parseRange(filename)
	valids := 0
	for password := lower; password <= upper; password++ {
		if validatePassword(password) {
			valids++
		}
	}
	return fmt.Sprint(valids)
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	lower, upper := parseRange(filename)
	valids := 0
	for password := lower; password <= upper; password++ {
		if validatePasswordV2(password) {
			valids++
		}
	}
	return fmt.Sprint(valids)
}

func parseRange(filename string) (lower int, upper int) {
	lines := files.GetLines(filename)
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	submatches := re.FindStringSubmatch(lines[0])
	lower = utils.MustAtoi(submatches[1])
	upper = utils.MustAtoi(submatches[2])
	return
}

func validatePassword(password int) bool {
	passStr := strconv.Itoa(password)
	chars := strings.Split(passStr, "")
	prevChar := "/"
	twoAdjacentSame := false
	for _, char := range chars {
		ordering := strings.Compare(prevChar, char)
		switch ordering {
		case 1:
			return false
		case 0:
			twoAdjacentSame = true
		}
		prevChar = char
	}
	return twoAdjacentSame
}

func validatePasswordV2(password int) bool {
	passStr := strconv.Itoa(password)
	chars := strings.Split(passStr, "")
	prevChar := "/"
	inSimilarString := false
	twoAdjacentSame := false
	twoPairFound := false
	for _, char := range chars {
		ordering := strings.Compare(prevChar, char)
		switch ordering {
		case 1:
			return false
		case 0:
			if !inSimilarString {
				twoAdjacentSame = true
				inSimilarString = true
			} else {
				twoAdjacentSame = false
			}
		case -1:
			if twoAdjacentSame {
				twoPairFound = true
			}
			inSimilarString = false
		}
		prevChar = char
	}
	return twoAdjacentSame || twoPairFound
}
