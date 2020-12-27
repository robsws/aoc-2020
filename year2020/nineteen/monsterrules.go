package nineteen

import (
	"aoc-go/utils"
	"fmt"
	"regexp"
	"strings"
)

type rule interface {
	resolve(rules map[int]rule) string
}

type primitiveRule struct {
	S string
}

func (r primitiveRule) resolve(rules map[int]rule) string {
	return r.S
}

type concatRule struct {
	IDs []int
}

func (r concatRule) resolve(rules map[int]rule) string {
	result := "(?:"
	for _, p := range r.IDs {
		result += rules[p].resolve(rules)
	}
	result += ")"
	return result
}

type decisionRule struct {
	Left  concatRule
	Right concatRule
}

func (r decisionRule) resolve(rules map[int]rule) string {
	return "(?:" + r.Left.resolve(rules) + "|" + r.Right.resolve(rules) + ")"
}

type repeatingRule struct {
	Left  int
	Right int
}

func (r repeatingRule) resolve(rules map[int]rule) string {
	if r.Right == -1 {
		return "(?:" + rules[r.Left].resolve(rules) + ")+"
	}
	// I'm not proud of this. xD
	sParts := [100]string{}
	for i := 0; i < 100; i++ {
		sParts[i] = "(?:" + rules[r.Left].resolve(rules) + "){" + fmt.Sprint(i+1) + "}(" + rules[r.Right].resolve(rules) + "){" + fmt.Sprint(i+1) + "}"
	}
	return "(?:" + strings.Join(sParts[:], "|") + ")"
}

func parseRules(fileStream chan string, altRules bool) (*regexp.Regexp, error) {
	// Gather rules
	rules := make(map[int]rule)
	ruleRe := regexp.MustCompile("^([0-9]+): (?:\"([a-z])\"|((?:[0-9]+ ?)+)(?: \\| ((?:[0-9]+ ?)+))?)$")
	for line := range fileStream {
		if line == "" {
			break
		}
		submatches := ruleRe.FindStringSubmatch(line)
		ruleID := utils.MustAtoi(submatches[1])
		if altRules {
			if ruleID == 8 {
				rules[ruleID] = repeatingRule{42, -1}
				continue
			} else if ruleID == 11 {
				rules[ruleID] = repeatingRule{42, 31}
				continue
			}
		}
		if submatches[2] != "" {
			rules[ruleID] = primitiveRule{submatches[2]}
			continue
		}
		leftConcatRule := parseConcatRule(submatches[3])
		if submatches[4] != "" {
			rightConcatRule := parseConcatRule(submatches[4])
			rules[ruleID] = decisionRule{leftConcatRule, rightConcatRule}
		} else {
			rules[ruleID] = leftConcatRule
		}
	}
	// Build regex
	reString := "^" + rules[0].resolve(rules) + "$"
	return regexp.Compile(reString)
}

func parseConcatRule(s string) concatRule {
	subIDStrs := strings.Split(s, " ")
	subIDs := make([]int, len(subIDStrs))
	for i, idStr := range subIDStrs {
		subIDs[i] = utils.MustAtoi(idStr)
	}
	return concatRule{subIDs}
}
