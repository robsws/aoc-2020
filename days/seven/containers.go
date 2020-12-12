package seven

import "aoc-2020/utils"

func getContainers(containedBy map[string][]string, colour string, top bool) utils.Set {
	containers := containedBy[colour]
	colours := utils.MakeSet()
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
