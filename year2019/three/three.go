package three

import (
	"aoc-go/files"
	"aoc-go/utils"
	"aoc-go/vector"
	"fmt"
	"regexp"
)

// PartOne - Find the closest wire crossing
func PartOne(filename string) string {
	lines := files.GetLines(filename)
	wire1 := parseInstructions(lines[0])
	wire2 := parseInstructions(lines[1])
	path1 := followInstructions(wire1)
	path2 := followInstructions(wire2)
	crossovers := findCrossovers(path1, path2)
	_, distance := findClosest(crossovers)
	return fmt.Sprint(distance)
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	lines := files.GetLines(filename)
	wire1 := parseInstructions(lines[0])
	wire2 := parseInstructions(lines[1])
	path1 := followInstructions(wire1)
	path2 := followInstructions(wire2)
	crossovers := findCrossovers(path1, path2)
	speed := findQuickest(crossovers)
	return fmt.Sprint(speed)
}

type instruction struct {
	direction string
	distance  int
}

func parseInstructions(line string) []instruction {
	re := regexp.MustCompile("([RDUL])([0-9]+)")
	matches := re.FindAllStringSubmatch(line, -1)
	instructions := make([]instruction, 0)
	for _, match := range matches {
		distance := utils.MustAtoi(match[2])
		instructions = append(instructions, instruction{match[1], distance})
	}
	return instructions
}

func followInstructions(instructions []instruction) trip {
	journey := trip{make(map[vector.Vec2]int), 0, vector.Vec2{X: 0, Y: 0}}
	for _, i := range instructions {
		switch i.direction {
		case "U":
			journey.recordPathAlongLine(vector.Vec2{X: 0, Y: -1}, i.distance)
		case "D":
			journey.recordPathAlongLine(vector.Vec2{X: 0, Y: 1}, i.distance)
		case "L":
			journey.recordPathAlongLine(vector.Vec2{X: -1, Y: 0}, i.distance)
		case "R":
			journey.recordPathAlongLine(vector.Vec2{X: 1, Y: 0}, i.distance)
		}
	}
	return journey
}

type trip struct {
	distanceTravelledToPoint map[vector.Vec2]int
	distanceTravelled        int
	position                 vector.Vec2
}

func (t *trip) recordPathAlongLine(direction vector.Vec2, distance int) {
	for i := 0; i < distance; i++ {
		t.position = t.position.Add(direction)
		t.distanceTravelled++
		_, alreadyVisited := t.distanceTravelledToPoint[t.position]
		if !alreadyVisited {
			t.distanceTravelledToPoint[t.position] = t.distanceTravelled
		}
	}
}

type crossover struct {
	point         vector.Vec2
	combinedSteps int
}

func findCrossovers(path1 trip, path2 trip) []crossover {
	crossovers := make([]crossover, 0)
	for point, dist1 := range path1.distanceTravelledToPoint {
		dist2, visited := path2.distanceTravelledToPoint[point]
		if visited {
			crossovers = append(crossovers, crossover{point, dist1 + dist2})
		}
	}
	return crossovers
}

func findClosest(crossovers []crossover) (point vector.Vec2, distance int) {
	closestDistance := utils.MAXINT
	var closestPoint vector.Vec2
	for _, c := range crossovers {
		distance := utils.AbsInt(c.point.X) + utils.AbsInt(c.point.Y)
		if distance < closestDistance {
			closestDistance = distance
			closestPoint = point
		}
	}
	return closestPoint, closestDistance
}

func findQuickest(crossovers []crossover) int {
	bestSpeed := utils.MAXINT
	for _, c := range crossovers {
		if c.combinedSteps < bestSpeed {
			bestSpeed = c.combinedSteps
		}
	}
	return bestSpeed
}
