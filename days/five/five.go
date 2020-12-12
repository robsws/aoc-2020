package five

import (
	"aoc-2020/files"
	"fmt"
	"sort"
)

// PartOne - find the highest seat ID on a boarding pass
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	passes := make([]boardingPass, 0)
	for line := range fileStream {
		passes = append(passes, parseBoardingPass(line))
	}
	sort.Sort(sort.Reverse(bySeatID(passes)))
	return fmt.Sprint(passes[0].CalcSeatID())
}

// PartTwo - find the missing boarding pass
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	passes := make([]boardingPass, 0)
	for line := range fileStream {
		passes = append(passes, parseBoardingPass(line))
	}
	sort.Sort(byLocation(passes))
	row := passes[0].row
	col := passes[0].col
	for _, pass := range passes {
		if pass.row != row || pass.col != col {
			mypass := boardingPass{loc: "", row: row, col: col}
			return fmt.Sprint(mypass.CalcSeatID())
		}
		col = (col + 1) % 8
		if col == 0 {
			row++
		}
	}
	return "No empty seat found."
}
