package seven

import (
	"aoc-go/set"
)

func getContainers(containedBy map[string][]string, colour string, top bool) set.StringSet {
	containers := containedBy[colour]
	colours := set.MakeStringSet()
	for _, container := range containers {
		colours.Union(getContainers(containedBy, container, false))
	}
	if !top {
		colours.Add(colour)
	}
	return colours
}

type container struct {
	Colour string
	Amount int
}

func countContained(containerOf map[string][]container, c container, top bool) int {
	total := 0
	for _, innerContainer := range containerOf[c.Colour] {
		total += c.Amount * countContained(containerOf, innerContainer, false)
	}
	if !top {
		total += c.Amount
	}
	return total
}
