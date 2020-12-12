package three

import (
	"aoc-2020/files"
	"fmt"
)

// PartOne - get number of trees hit
func PartOne(filename string) string {
	resultStream := make(chan int)
	go countTrees(filename, 3, 1, resultStream)
	trees := <-resultStream
	return fmt.Sprint(trees)
}

// PartTwo - get number of tree hit for various directions
func PartTwo(filename string) string {
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
	return fmt.Sprint(trees)
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
