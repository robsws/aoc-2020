package six

import (
	"aoc-go/files"
	"aoc-go/set"
	"fmt"
	"strings"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	orbits := parseOrbits(filename)
	total := 0
	for planet := range orbits {
		planet2 := planet
		for planet2 != "COM" {
			planet2 = orbits[planet2]
			total++
		}
	}
	return fmt.Sprint(total)
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	orbits := parseOrbits(filename)
	// Find the common planet on the path to COM from YOU and SAN
	pathToCom := set.MakeStringSet()
	youSteps := 0
	youStepsToPlanet := make(map[string]int)
	planet := "YOU"
	pathToCom.Add(planet)
	youStepsToPlanet[planet] = youSteps
	for planet != "COM" {
		planet = orbits[planet]
		pathToCom.Add(planet)
		youSteps++
		youStepsToPlanet[planet] = youSteps
	}
	planet = "SAN"
	santaSteps := 0
	for planet != "COM" {
		santaSteps++
		planet = orbits[planet]
		if pathToCom.Contains(planet) {
			break
		}
	}
	return fmt.Sprint(youStepsToPlanet[planet] + santaSteps - 2)
}

func parseOrbits(filename string) map[string]string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	orbits := make(map[string]string)
	for line := range fileStream {
		planets := strings.Split(line, ")")
		orbits[planets[1]] = planets[0]
	}
	return orbits
}
