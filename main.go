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
	daythreepartone(filename)
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

func daythreepartone(filename string) {
	resultStream := make(chan int)
	go countTrees(filename, 3, 1, resultStream)
	trees := <-resultStream
	fmt.Println(trees)
}

func daythreeparttwo(filename string) {
	resultStream := make(chan int)
	go countTrees(filename, 1, 1, resultStream)
	go countTrees(filename, 3, 1, resultStream)
	go countTrees(filename, 5, 1, resultStream)
	go countTrees(filename, 7, 1, resultStream)
	go countTrees(filename, 1, 2, resultStream)
	trees := 1
	for i := 0; i < 5; i++ {
		trees *= <-resultStream
	}
	fmt.Println(trees)
}

func countTrees(filename string, dx int, dy int, resultStream chan int) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	x := 0
	y := 0
	trees := 0
	for line := range fileStream {
		if y%dy == 0 {
			if line[x] == '#' {
				trees++
			}
			x = (x + dx) % len(line)
		}
		y++
	}
	resultStream <- trees
}
