package fourteen

import (
	"aoc-2020/files"
	"aoc-2020/utils"
	"fmt"
	"regexp"
)

// PartOne - docking decoder v1
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	maskRe := regexp.MustCompile("^mask = ([X01]+)$")
	memRe := regexp.MustCompile("^mem\\[([0-9]+)\\] = ([0-9]+)$")
	mem := make(map[int]int)
	mask := [36]rune{}
	for line := range fileStream {
		maskMatch := maskRe.FindStringSubmatch(line)
		if maskMatch != nil {
			for i, s := range maskMatch[1] {
				mask[i] = rune(s)
			}
			continue
		}
		memMatch := memRe.FindStringSubmatch(line)
		if memMatch != nil {
			addr := utils.MustAtoi(memMatch[1])
			val := utils.MustAtoi(memMatch[2])
			mem[addr] = applyMask(val, mask)
		}
	}
	total := 0
	for _, val := range mem {
		total += val
	}
	return fmt.Sprint(total)
}

// PartTwo - docking decoder v2
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	maskRe := regexp.MustCompile("^mask = ([X01]+)$")
	memRe := regexp.MustCompile("^mem\\[([0-9]+)\\] = ([0-9]+)$")
	mem := make(map[int]int)
	mask := [36]rune{}
	for line := range fileStream {
		maskMatch := maskRe.FindStringSubmatch(line)
		if maskMatch != nil {
			for i, s := range maskMatch[1] {
				mask[i] = rune(s)
			}
			continue
		}
		memMatch := memRe.FindStringSubmatch(line)
		if memMatch != nil {
			inputAddr := utils.MustAtoi(memMatch[1])
			val := utils.MustAtoi(memMatch[2])
			addrs := applyV2Mask(inputAddr, mask)
			for _, addr := range addrs {
				mem[addr] = val
			}
		}
	}
	total := 0
	for _, val := range mem {
		total += val
	}
	return fmt.Sprint(total)
}
