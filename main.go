package main

import (
	"aoc-2020/files"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := os.Args[1]
	daytwoparttwo(filename)
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
