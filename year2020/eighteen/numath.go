package eighteen

import (
	"aoc-go/utils"
	"fmt"
	"strconv"
)

func evalLeftToRight(tokens []string) int {
	tokens = removeSpaces(tokens)
	tokens = resolveBrackets(tokens, evalLeftToRight)
	// Resolve this level of arithmetic
	result := 0
	op := ""
	for i := 0; i < len(tokens); i++ {
		// Resolve primitive value
		val, err := strconv.Atoi(tokens[i])
		if err == nil {
			switch op {
			case "+":
				result += val
			case "*":
				result *= val
			default:
				result = val
			}
			continue
		}
		// Resolve operator
		if tokens[i] == "+" || tokens[i] == "*" {
			op = tokens[i]
			continue
		}
	}
	return result
}

func evalAdditionFirst(tokens []string) int {
	tokens = removeSpaces(tokens)
	tokens = resolveBrackets(tokens, evalAdditionFirst)
	tokens = resolveAddition(tokens)
	// Resolve this level of arithmetic
	result := 1
	for i := 0; i < len(tokens); i++ {
		// Resolve primitive value
		val, err := strconv.Atoi(tokens[i])
		if err == nil {
			result *= val
			continue
		}
	}
	return result
}

func resolveBrackets(tokens []string, eval func([]string) int) []string {
	// Look for brackets
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(" {
			leftbrackets := 0
			var j int
			for j = i + 1; j < len(tokens); j++ {
				switch tokens[j] {
				case "(":
					leftbrackets++
				case ")":
					leftbrackets--
				}
				if leftbrackets < 0 {
					break
				}
			}
			val := eval(tokens[i+1 : j])
			newtokens := append(tokens[:i], fmt.Sprint(val))
			tokens = append(newtokens, tokens[j+1:]...)
		}
	}
	return tokens
}

func resolveAddition(tokens []string) []string {
	// Look for addition
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			lval := utils.MustAtoi(tokens[i-1])
			rval := utils.MustAtoi(tokens[i+1])
			val := lval + rval
			newtokens := append(tokens[:i-1], fmt.Sprint(val))
			tokens = append(newtokens, tokens[i+2:]...)
			i--
		}
	}
	return tokens
}

func removeSpaces(tokens []string) []string {
	newTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if token != " " {
			newTokens = append(newTokens, token)
		}
	}
	return newTokens
}
