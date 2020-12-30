package eight

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"strings"
)

// PartOne - Not yet implemented
func PartOne(filename string) string {
	layers := parseImage(filename, 25, 6)
	digitCounts := findDigitCounts(layers)
	// find layer with fewest zeros
	minZeros := utils.MAXINT
	minZeroLayer := -1
	for layer, counts := range digitCounts {
		if counts[0] < minZeros {
			minZeros = counts[0]
			minZeroLayer = layer
		}
	}
	// 1s x 2s
	return fmt.Sprint(digitCounts[minZeroLayer][1] * digitCounts[minZeroLayer][2])
}

// PartTwo - Not yet implemented
func PartTwo(filename string) string {
	layers := parseImage(filename, 25, 6)
	image := displayImage(layers)
	result := fmt.Sprintln()
	for _, row := range image {
		result += fmt.Sprintln(row)
	}
	return result
}

func parseImage(filename string, width int, height int) [][][]int {
	line := files.GetLines(filename)[0]
	data := strings.Split(line, "")
	layerSize := width * height
	layers := make([][][]int, 0)
	for i, pixel := range data {
		pixelNo := i % layerSize
		layerNo := int(i / layerSize)
		if pixelNo == 0 {
			layers = append(layers, make([][]int, height))
			for y := 0; y < height; y++ {
				layers[layerNo][y] = make([]int, width)
			}
		}
		x := pixelNo % width
		y := int(pixelNo / width)
		layers[layerNo][y][x] = utils.MustAtoi(pixel)
	}
	return layers
}

func findDigitCounts(layers [][][]int) []map[int]int {
	digitCount := make([]map[int]int, len(layers))
	for l := range layers {
		digitCount[l] = make(map[int]int)
		for y := range layers[l] {
			for x := range layers[l][y] {
				digitCount[l][layers[l][y][x]]++
			}
		}
	}
	return digitCount
}

func displayImage(layers [][][]int) []string {
	// init image
	image := make([][]int, len(layers[0]))
	for y := 0; y < len(layers[0]); y++ {
		image[y] = make([]int, len(layers[0][0]))
	}
	// overlay layers onto image, starting with last
	for l := len(layers) - 1; l >= 0; l-- {
		for y := range layers[l] {
			for x := range layers[l][y] {
				pixel := layers[l][y][x]
				if pixel != 2 {
					image[y][x] = pixel
				}
			}
		}
	}
	// serialize into string
	lines := make([]string, 0)
	for y := range image {
		line := ""
		for _, pixel := range image[y] {
			if pixel == 0 {
				line += " "
			} else {
				line += "#"
			}
		}
		lines = append(lines, line)
	}
	return lines
}
