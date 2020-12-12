package days

import (
	"aoc-2020/days/eight"
	"aoc-2020/days/eighteen"
	"aoc-2020/days/eleven"
	"aoc-2020/days/fifteen"
	"aoc-2020/days/five"
	"aoc-2020/days/four"
	"aoc-2020/days/fourteen"
	"aoc-2020/days/nine"
	"aoc-2020/days/nineteen"
	"aoc-2020/days/one"
	"aoc-2020/days/seven"
	"aoc-2020/days/seventeen"
	"aoc-2020/days/six"
	"aoc-2020/days/sixteen"
	"aoc-2020/days/ten"
	"aoc-2020/days/thirteen"
	"aoc-2020/days/three"
	"aoc-2020/days/twelve"
	"aoc-2020/days/twenty"
	"aoc-2020/days/twentyfive"
	"aoc-2020/days/twentyfour"
	"aoc-2020/days/twentyone"
	"aoc-2020/days/twentythree"
	"aoc-2020/days/twentytwo"
	"aoc-2020/days/two"
	"path/filepath"
)

// GetParts - get the functions for part one and two for given day
func GetParts(day int, test bool) (func(string), func(string), string) {
	filename := "input.txt"
	if test {
		filename = "input_test.txt"
	}
	switch day {
	case 1:
		return one.PartOne, one.PartTwo, filepath.Join("days/one/", filename)
	case 2:
		return two.PartOne, two.PartTwo, filepath.Join("days/two/", filename)
	case 3:
		return three.PartOne, three.PartTwo, filepath.Join("days/three/", filename)
	case 4:
		return four.PartOne, four.PartTwo, filepath.Join("days/four/", filename)
	case 5:
		return five.PartOne, five.PartTwo, filepath.Join("days/five/", filename)
	case 6:
		return six.PartOne, six.PartTwo, filepath.Join("days/six/", filename)
	case 7:
		return seven.PartOne, seven.PartTwo, filepath.Join("days/seven/", filename)
	case 8:
		return eight.PartOne, eight.PartTwo, filepath.Join("days/eight/", filename)
	case 9:
		return nine.PartOne, nine.PartTwo, filepath.Join("days/nine/", filename)
	case 10:
		return ten.PartOne, ten.PartTwo, filepath.Join("days/ten/", filename)
	case 11:
		return eleven.PartOne, eleven.PartTwo, filepath.Join("days/eleven/", filename)
	case 12:
		return twelve.PartOne, twelve.PartTwo, filepath.Join("days/twelve/", filename)
	case 13:
		return thirteen.PartOne, thirteen.PartTwo, filepath.Join("days/thirteen/", filename)
	case 14:
		return fourteen.PartOne, fourteen.PartTwo, filepath.Join("days/fourteen/", filename)
	case 15:
		return fifteen.PartOne, fifteen.PartTwo, filepath.Join("days/fifteen/", filename)
	case 16:
		return sixteen.PartOne, sixteen.PartTwo, filepath.Join("days/sixteen/", filename)
	case 17:
		return seventeen.PartOne, seventeen.PartTwo, filepath.Join("days/seventeen/", filename)
	case 18:
		return eighteen.PartOne, eighteen.PartTwo, filepath.Join("days/eighteen/", filename)
	case 19:
		return nineteen.PartOne, nineteen.PartTwo, filepath.Join("days/nineteen/", filename)
	case 20:
		return twenty.PartOne, twenty.PartTwo, filepath.Join("days/twenty/", filename)
	case 21:
		return twentyone.PartOne, twentyone.PartTwo, filepath.Join("days/twentyone/", filename)
	case 22:
		return twentytwo.PartOne, twentytwo.PartTwo, filepath.Join("days/twentytwo/", filename)
	case 23:
		return twentythree.PartOne, twentythree.PartTwo, filepath.Join("days/twentythree/", filename)
	case 24:
		return twentyfour.PartOne, twentyfour.PartTwo, filepath.Join("days/twentyfour/", filename)
	default:
		return twentyfive.PartOne, twentyfive.PartTwo, filepath.Join("days/twentyfive/", filename)
	}
}
