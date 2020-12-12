package eleven

import (
	"aoc-2020/files"
	"fmt"
)

// PartOne - simulate airport seating using adjacent algorithm and calculate convergent occupied seats
func PartOne(filename string) {
	lines := files.GetLines(filename)
	height := len(lines)
	width := len(lines[0])
	grid := initGrid(width, height)
	// Initialise the grid with the input
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = rune(lines[y][x])
		}
	}
	// Simulate rounds until grid doesn't change
	prevgrid := initGrid(width, height)
	for !gridsAreEqual(grid, prevgrid) {
		prevgrid = grid
		grid = simulateSeats(grid, false)
	}
	// Count occupied seats
	fmt.Println(countTotalOccupied(grid))
}

// PartTwo - simulate airport seating using visible algorithm and calculate convergent occupied seats
func PartTwo(filename string) {
	lines := files.GetLines(filename)
	height := len(lines)
	width := len(lines[0])
	grid := initGrid(width, height)
	// Initialise the grid with the input
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			grid[x][y] = rune(lines[y][x])
		}
	}
	// Simulate rounds until grid doesn't change
	prevgrid := initGrid(width, height)
	for !gridsAreEqual(grid, prevgrid) {
		prevgrid = grid
		grid = simulateSeats(grid, true)
	}
	// Count occupied seats
	fmt.Println(countTotalOccupied(grid))
}
