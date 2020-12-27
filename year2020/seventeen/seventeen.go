package seventeen

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"strings"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	activePoints := parseInputGrid(filename)
	for i := 0; i < 6; i++ {
		activePoints = simulateGrid(activePoints)
	}
	return fmt.Sprint(activePoints.Len())
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	activePoints := parseInputGrid4D(filename)
	for i := 0; i < 6; i++ {
		activePoints = simulateGrid4D(activePoints)
	}
	return fmt.Sprint(activePoints.Len())
}

func parseInputGrid(filename string) utils.Set {
	lines := files.GetLines(filename)
	activePoints := utils.MakeSet()
	for y, line := range lines {
		places := strings.Split(line, "")
		for z, c := range places {
			if c == "#" {
				activePoints.Add(makeCoordString(0, y, z))
			}
		}
	}
	return activePoints
}

func parseInputGrid4D(filename string) utils.Set {
	lines := files.GetLines(filename)
	activePoints := utils.MakeSet()
	for y, line := range lines {
		places := strings.Split(line, "")
		for z, c := range places {
			if c == "#" {
				activePoints.Add(makeCoordString4D(0, 0, y, z))
			}
		}
	}
	return activePoints
}

func simulateGrid(activePoints utils.Set) utils.Set {
	newActivePoints := utils.MakeSet()
	for _, point := range activePoints.ToSlice() {
		x, y, z := parseCoordString(point.(string))
		newPoints, active := updatePoint(x, y, z, activePoints, true)
		newActivePoints.Union(newPoints)
		if active {
			newActivePoints.Add(makeCoordString(x, y, z))
		}
	}
	return newActivePoints
}

func simulateGrid4D(activePoints utils.Set) utils.Set {
	newActivePoints := utils.MakeSet()
	for _, point := range activePoints.ToSlice() {
		w, x, y, z := parseCoordString4D(point.(string))
		newPoints, active := updatePoint4D(w, x, y, z, activePoints, true)
		newActivePoints.Union(newPoints)
		if active {
			newActivePoints.Add(makeCoordString4D(w, x, y, z))
		}
	}
	return newActivePoints
}

func updatePoint(x, y, z int, activePoints utils.Set, recurse bool) (utils.Set, bool) {
	activeSurrounding := 0
	newPoints := utils.MakeSet()
	for x1 := x - 1; x1 <= x+1; x1++ {
		for y1 := y - 1; y1 <= y+1; y1++ {
			for z1 := z - 1; z1 <= z+1; z1++ {
				if !(x == x1 && y == y1 && z == z1) {
					coord := makeCoordString(x1, y1, z1)
					if activePoints.Contains(coord) {
						activeSurrounding++
					}
					if recurse && !newPoints.Contains(coord) {
						_, active := updatePoint(x1, y1, z1, activePoints, false)
						if active {
							newPoints.Add(makeCoordString(x1, y1, z1))
						}
					}
				}
			}
		}
	}
	iAmActive := activePoints.Contains(makeCoordString(x, y, z))
	if iAmActive && !(activeSurrounding == 2 || activeSurrounding == 3) {
		iAmActive = false
	} else if !iAmActive && activeSurrounding == 3 {
		iAmActive = true
	}
	return newPoints, iAmActive
}

func updatePoint4D(w, x, y, z int, activePoints utils.Set, recurse bool) (utils.Set, bool) {
	activeSurrounding := 0
	newPoints := utils.MakeSet()
	for w1 := w - 1; w1 <= w+1; w1++ {
		for x1 := x - 1; x1 <= x+1; x1++ {
			for y1 := y - 1; y1 <= y+1; y1++ {
				for z1 := z - 1; z1 <= z+1; z1++ {
					if !(w == w1 && x == x1 && y == y1 && z == z1) {
						coord := makeCoordString4D(w1, x1, y1, z1)
						if activePoints.Contains(coord) {
							activeSurrounding++
						}
						if recurse && !newPoints.Contains(coord) {
							_, active := updatePoint4D(w1, x1, y1, z1, activePoints, false)
							if active {
								newPoints.Add(makeCoordString4D(w1, x1, y1, z1))
							}
						}
					}
				}
			}
		}
	}
	iAmActive := activePoints.Contains(makeCoordString4D(w, x, y, z))
	if iAmActive && !(activeSurrounding == 2 || activeSurrounding == 3) {
		iAmActive = false
	} else if !iAmActive && activeSurrounding == 3 {
		iAmActive = true
	}
	return newPoints, iAmActive
}

func makeCoordString(x, y, z int) string {
	return fmt.Sprint(x, ",", y, ",", z)
}

func makeCoordString4D(w, x, y, z int) string {
	return fmt.Sprint(w, ",", x, ",", y, ",", z)
}

func parseCoordString(coord string) (int, int, int) {
	parts := strings.Split(coord, ",")
	return utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1]), utils.MustAtoi(parts[2])
}

func parseCoordString4D(coord string) (int, int, int, int) {
	parts := strings.Split(coord, ",")
	return utils.MustAtoi(parts[0]), utils.MustAtoi(parts[1]), utils.MustAtoi(parts[2]), utils.MustAtoi(parts[3])
}

func printActivePoints(activePoints utils.Set) {
	for _, point := range activePoints.ToSlice() {
		fmt.Println(point)
	}
}
