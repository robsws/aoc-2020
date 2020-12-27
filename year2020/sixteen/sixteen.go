package sixteen

import (
	"aoc-go/files"
	"aoc-go/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	rules := readRules(fileStream)
	readMyTicket(fileStream)
	tickets := readNearbyTickets(fileStream)
	errorRate := determineErrorRate(tickets, rules)
	return fmt.Sprint(errorRate)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	rules := readRules(fileStream)
	myticket := readMyTicket(fileStream)
	tickets := readNearbyTickets(fileStream)
	fieldPositions := determineFieldPositions(tickets, rules)
	departureRe := regexp.MustCompile("^departure")
	answer := 1
	for field, position := range fieldPositions {
		if departureRe.MatchString(field) {
			answer *= myticket[position]
		}
	}
	return fmt.Sprint(answer)
}

func readRules(fileStream chan string) map[string]utils.Set {
	ruleRe := regexp.MustCompile("([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	rules := make(map[string]utils.Set)
	for line := range fileStream {
		if line == "" {
			break
		}
		submatches := ruleRe.FindStringSubmatch(line)
		rule := submatches[1]
		min1 := utils.MustAtoi(submatches[2])
		max1 := utils.MustAtoi(submatches[3])
		min2 := utils.MustAtoi(submatches[4])
		max2 := utils.MustAtoi(submatches[5])
		rules[rule] = utils.MakeSet()
		for i := min1; i <= max1; i++ {
			rules[rule].Add(i)
		}
		for i := min2; i <= max2; i++ {
			rules[rule].Add(i)
		}
	}
	return rules
}

func readMyTicket(fileStream chan string) []int {
	<-fileStream // Consume the header
	nums := parseTicket(<-fileStream)
	<-fileStream // Consume space
	return nums
}

func readNearbyTickets(fileStream chan string) [][]int {
	<-fileStream // Consume the header
	tickets := make([][]int, 0)
	for line := range fileStream {
		nums := parseTicket(line)
		tickets = append(tickets, nums)
	}
	return tickets
}

func parseTicket(line string) []int {
	numstrs := strings.Split(line, ",")
	nums := make([]int, len(numstrs))
	for i, s := range numstrs {
		nums[i] = utils.MustAtoi(s)
	}
	return nums
}

func determineErrorRate(tickets [][]int, rules map[string]utils.Set) int {
	errorRate := 0
	for _, ticket := range tickets {
		for _, n := range ticket {
			valid := false
			for _, validValues := range rules {
				if validValues.Contains(n) {
					valid = true
					break
				}
			}
			if !valid {
				errorRate += n
			}
		}
	}
	return errorRate
}

func determineFieldPositions(tickets [][]int, rules map[string]utils.Set) map[string]int {
	invalidPositions := make(map[string]utils.Set)
	for field := range rules {
		invalidPositions[field] = utils.MakeSet()
	}
	// Find invalid positions for each field based on nearby tickets
	for _, ticket := range tickets {
		// check ticket is valid
		valid := true
		for _, value := range ticket {
			valueValid := false
			for _, validValues := range rules {
				if validValues.Contains(value) {
					valueValid = true
					break
				}
			}
			if !valueValid {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		// record any invalid positions for fields based on rules
		for position, value := range ticket {
			for field, values := range rules {
				if !values.Contains(value) {
					invalidPositions[field].Add(position)
				}
			}
		}
	}
	// Invert to get valid positions
	validPositions := make(map[string]utils.Set)
	for field, invalidPos := range invalidPositions {
		validPositions[field] = utils.MakeSet()
		for i := 0; i < len(rules); i++ {
			if !invalidPos.Contains(i) {
				validPositions[field].Add(i)
			}
		}
	}
	// Loop through all permutations of valid positions until
	// one is valid for all fields
	pairs := make(pairList, 0)
	for k, v := range validPositions {
		pairs = append(pairs, pair{k, v.Len()})
	}
	sort.Sort(pairs)
	permutation, ok := findValidPermutation(pairs, validPositions, make([]string, 0), utils.MakeSet())
	if !ok {
		return nil
	}
	return permutation
}

func findValidPermutation(pairs pairList, validPositions map[string]utils.Set, fixedFields []string, fixedPositions utils.Set) (map[string]int, bool) {
	if len(fixedFields) == len(validPositions) {
		return make(map[string]int), true
	}
	for _, k := range pairs {
		field := k.Key
		positions := validPositions[field]
		if !utils.StringSliceContains(fixedFields, field) {
			for _, position := range positions.ToSlice() {
				if !fixedPositions.Contains(position) {
					newFixedFields := make([]string, len(fixedFields), len(fixedFields)+1)
					for i, f := range fixedFields {
						newFixedFields[i] = f
					}
					newFixedPositions := utils.MakeSet()
					for _, p := range fixedPositions.ToSlice() {
						newFixedPositions.Add(p)
					}
					newFixedFields = append(newFixedFields, field)
					newFixedPositions.Add(position)
					newPermutation, ok := findValidPermutation(pairs, validPositions, newFixedFields, newFixedPositions)
					if ok {
						newPermutation[field] = position.(int)
						return newPermutation, true
					}
				}
			}
		}
	}
	return nil, false
}

type pair struct {
	Key   string
	Value int
}
type pairList []pair

func (l pairList) Len() int           { return len(l) }
func (l pairList) Less(i, j int) bool { return l[i].Value < l[j].Value }
func (l pairList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
