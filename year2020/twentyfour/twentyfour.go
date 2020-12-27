package twentyfour

import (
	"aoc-go/files"
	"aoc-go/set"
	"fmt"
	"regexp"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	coords := parseTiles(filename)
	return fmt.Sprint(coords.Len())
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	coords := parseTiles(filename)
	for i := 0; i < 100; i++ {
		coords = doRound(coords)
	}
	return fmt.Sprint(coords.Len())
}

func parseDirections(line string) [3]int {
	re := regexp.MustCompile("(se|sw|nw|ne|w|e)")
	directions := re.FindAllString(line, -1)
	coord := [3]int{}
	for _, d := range directions {
		switch d {
		case "w":
			coord[0]--
			coord[1]++
		case "e":
			coord[0]++
			coord[1]--
		case "sw":
			coord[0]--
			coord[2]++
		case "ne":
			coord[0]++
			coord[2]--
		case "se":
			coord[1]--
			coord[2]++
		case "nw":
			coord[1]++
			coord[2]--
		}
	}
	return coord
}

func parseTiles(filename string) set.CoordSet {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	coords := set.MakeCoordSet()
	for line := range fileStream {
		coord := parseDirections(line)
		if coords.Contains(coord) {
			coords.Remove(coord)
		} else {
			coords.Add(coord)
		}
	}
	return coords
}

func doRound(coords set.CoordSet) set.CoordSet {
	newCoords := set.MakeCoordSet()
	for _, coord := range coords.ToSlice() {
		checkAndUpdateSurrounding(coord, coords, &newCoords, true)
	}
	return newCoords
}

func checkAndUpdateSurrounding(coord [3]int, coords set.CoordSet, newCoords *set.CoordSet, recurse bool) {
	blackTiles := 0
	offsets := [6][3]int{
		{0, 1, -1},
		{1, 0, -1},
		{1, -1, 0},
		{0, -1, 1},
		{-1, 0, 1},
		{-1, 1, 0},
	}
	for _, offset := range offsets {
		surroundingTile := [3]int{
			coord[0] + offset[0],
			coord[1] + offset[1],
			coord[2] + offset[2],
		}
		if coords.Contains(surroundingTile) {
			blackTiles++
		} else if recurse {
			checkAndUpdateSurrounding(surroundingTile, coords, newCoords, false)
		}
	}
	isBlack := coords.Contains(coord)
	if isBlack {
		if blackTiles > 0 && blackTiles < 3 {
			newCoords.Add(coord)
		}
	} else {
		if blackTiles == 2 {
			newCoords.Add(coord)
		}
	}
}
