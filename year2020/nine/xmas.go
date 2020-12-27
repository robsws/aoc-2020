package nine

import "aoc-go/set"

func findNonSumNumber(numbers []int, preamble int) (int, bool) {
	for i := 0; i < len(numbers)-preamble; i++ {
		sums := set.MakeIntSet()
		for j := 0; j < preamble; j++ {
			for k := 0; k < preamble; k++ {
				if k != j {
					sums.Add(numbers[i+k] + numbers[i+j])
				}
			}
		}
		if !sums.Contains(numbers[i+preamble]) {
			return numbers[i+preamble], true
		}
	}
	return -1, false
}

func findContiguousSumRange(numbers []int, n int) (int, int, bool) {
	for i := 0; i < len(numbers)-1; i++ {
		total := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			total += numbers[j]
			if total == n {
				return i, j, true
			} else if total > n {
				break
			}
		}
	}
	return -1, -1, false
}
