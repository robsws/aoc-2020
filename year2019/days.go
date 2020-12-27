package year2019

import (
	"aoc-go/year2019/eight"
	"aoc-go/year2019/eighteen"
	"aoc-go/year2019/eleven"
	"aoc-go/year2019/fifteen"
	"aoc-go/year2019/five"
	"aoc-go/year2019/four"
	"aoc-go/year2019/fourteen"
	"aoc-go/year2019/nine"
	"aoc-go/year2019/nineteen"
	"aoc-go/year2019/one"
	"aoc-go/year2019/seven"
	"aoc-go/year2019/seventeen"
	"aoc-go/year2019/six"
	"aoc-go/year2019/sixteen"
	"aoc-go/year2019/ten"
	"aoc-go/year2019/thirteen"
	"aoc-go/year2019/three"
	"aoc-go/year2019/twelve"
	"aoc-go/year2019/twenty"
	"aoc-go/year2019/twentyfive"
	"aoc-go/year2019/twentyfour"
	"aoc-go/year2019/twentyone"
	"aoc-go/year2019/twentythree"
	"aoc-go/year2019/twentytwo"
	"aoc-go/year2019/two"
	"path/filepath"
)

// GetParts - get the functions for part one and two for given day
func GetParts(day int, test bool) (func(string) string, func(string) string, string) {
	filename := "input.txt"
	if test {
		filename = "input_test.txt"
	}
	switch day {
	case 1:
		return one.PartOne, one.PartTwo, filepath.Join("year2019/one/", filename)
	case 2:
		return two.PartOne, two.PartTwo, filepath.Join("year2019/two/", filename)
	case 3:
		return three.PartOne, three.PartTwo, filepath.Join("year2019/three/", filename)
	case 4:
		return four.PartOne, four.PartTwo, filepath.Join("year2019/four/", filename)
	case 5:
		return five.PartOne, five.PartTwo, filepath.Join("year2019/five/", filename)
	case 6:
		return six.PartOne, six.PartTwo, filepath.Join("year2019/six/", filename)
	case 7:
		return seven.PartOne, seven.PartTwo, filepath.Join("year2019/seven/", filename)
	case 8:
		return eight.PartOne, eight.PartTwo, filepath.Join("year2019/eight/", filename)
	case 9:
		return nine.PartOne, nine.PartTwo, filepath.Join("year2019/nine/", filename)
	case 10:
		return ten.PartOne, ten.PartTwo, filepath.Join("year2019/ten/", filename)
	case 11:
		return eleven.PartOne, eleven.PartTwo, filepath.Join("year2019/eleven/", filename)
	case 12:
		return twelve.PartOne, twelve.PartTwo, filepath.Join("year2019/twelve/", filename)
	case 13:
		return thirteen.PartOne, thirteen.PartTwo, filepath.Join("year2019/thirteen/", filename)
	case 14:
		return fourteen.PartOne, fourteen.PartTwo, filepath.Join("year2019/fourteen/", filename)
	case 15:
		return fifteen.PartOne, fifteen.PartTwo, filepath.Join("year2019/fifteen/", filename)
	case 16:
		return sixteen.PartOne, sixteen.PartTwo, filepath.Join("year2019/sixteen/", filename)
	case 17:
		return seventeen.PartOne, seventeen.PartTwo, filepath.Join("year2019/seventeen/", filename)
	case 18:
		return eighteen.PartOne, eighteen.PartTwo, filepath.Join("year2019/eighteen/", filename)
	case 19:
		return nineteen.PartOne, nineteen.PartTwo, filepath.Join("year2019/nineteen/", filename)
	case 20:
		return twenty.PartOne, twenty.PartTwo, filepath.Join("year2019/twenty/", filename)
	case 21:
		return twentyone.PartOne, twentyone.PartTwo, filepath.Join("year2019/twentyone/", filename)
	case 22:
		return twentytwo.PartOne, twentytwo.PartTwo, filepath.Join("year2019/twentytwo/", filename)
	case 23:
		return twentythree.PartOne, twentythree.PartTwo, filepath.Join("year2019/twentythree/", filename)
	case 24:
		return twentyfour.PartOne, twentyfour.PartTwo, filepath.Join("year2019/twentyfour/", filename)
	default:
		return twentyfive.PartOne, twentyfive.PartTwo, filepath.Join("year2019/twentyfive/", filename)
	}
}
