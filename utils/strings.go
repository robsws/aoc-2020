package utils

import (
	"log"
	"strconv"
)

// MustAtoi - convert string to int and error out if it doesn't work
func MustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// MustAtoiAll - convert all strings in list to ints
func MustAtoiAll(list []string) []int {
	intList := make([]int, len(list))
	for i, s := range list {
		intList[i] = MustAtoi(s)
	}
	return intList
}
