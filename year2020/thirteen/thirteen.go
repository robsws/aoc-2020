package thirteen

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"strings"
)

// PartOne - find best bus to catch to the airport
func PartOne(filename string) string {
	lines := files.GetLines(filename)
	time := utils.MustAtoi(lines[0])
	busStrs := strings.Split(lines[1], ",")
	buses := make([]int, 0)
	for _, s := range busStrs {
		if s != "x" {
			buses = append(buses, utils.MustAtoi(s))
		}
	}
	waitTimes := make([]int, len(buses))
	for i, bus := range buses {
		timeSinceBus := time % bus
		if timeSinceBus == 0 {
			waitTimes[i] = 0
		} else {
			waitTimes[i] = bus - timeSinceBus
		}
	}
	bestBusI, bestWaitTime := utils.MinIntInList(waitTimes)
	return fmt.Sprint(buses[bestBusI] * bestWaitTime)
}

// PartTwo - work out when buses arrive at times listed
func PartTwo(filename string) string {
	getBusAlignmentTime([]int{3, 4, 5}, []int{0, 3, 4})
	lines := files.GetLines(filename)
	busStrs := strings.Split(lines[1], ",")
	buses := make([]int, 0)
	busOffsets := make([]int, 0)
	for i, s := range busStrs {
		if s != "x" {
			bus := utils.MustAtoi(s)
			buses = append(buses, bus)
			if i == 0 {
				busOffsets = append(busOffsets, 0)
			} else {
				busOffsets = append(busOffsets, bus-i)
			}
		}
	}
	return fmt.Sprint(getBusAlignmentTime(buses, busOffsets))
}

func getBusAlignmentTime(buses []int, busOffsets []int) int {
	// Chinese remainder theorem solution
	product := 1
	for _, bus := range buses {
		product *= bus
	}
	quotients := make([]int, len(buses))
	for i, bus := range buses {
		quotients[i] = product / bus
	}
	bezouts := make([]int, len(buses))
	for i, bus := range buses {
		_, _, bezout := extendedEuclid(bus, quotients[i])
		bezouts[i] = bezout
	}
	answer := 0
	for i, bezout := range bezouts {
		answer += busOffsets[i] * quotients[i] * bezout
	}
	// Reduce to the smallest possible greater than 0
	for answer < 0 {
		answer += product
	}
	for answer > product {
		answer -= product
	}
	return answer
}

func extendedEuclid(a, b int) (int, int, int) {
	remainders := [3]int{a, b, -1}
	bezoutA := [3]int{1, 0, -1}
	bezoutB := [3]int{0, 1, -1}
	for remainders[2] != 0 {
		quotient := int(remainders[0] / remainders[1])
		remainders[2] = remainders[0] % remainders[1]
		bezoutA[2] = bezoutA[0] - quotient*bezoutA[1]
		bezoutB[2] = bezoutB[0] - quotient*bezoutB[1]
		remainders[0] = remainders[1]
		remainders[1] = remainders[2]
		bezoutA[0] = bezoutA[1]
		bezoutA[1] = bezoutA[2]
		bezoutB[0] = bezoutB[1]
		bezoutB[1] = bezoutB[2]
	}
	return remainders[0], bezoutA[0], bezoutB[0]
}
