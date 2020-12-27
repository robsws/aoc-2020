package fourteen

import "aoc-go/utils"

func applyMask(n int, mask [36]rune) int {
	nbin := intToBinary36(n)
	for i, bit := range mask {
		if bit == '0' || bit == '1' {
			nbin[i] = bit
		}
	}
	return binary36ToInt(nbin)
}

func applyV2Mask(n int, mask [36]rune) []int {
	results := make([]int, 0)
	variables := make([]int, 0)
	nbin := intToBinary36(n)
	for i, bit := range mask {
		switch bit {
		case '1':
			nbin[i] = '1'
		case 'X':
			nbin[i] = '0'
			variables = append(variables, utils.PowInt(2, 35-i))
		}
	}
	result := binary36ToInt(nbin)
	results = append(results, result)
	// compute possibilities given floating digits
	for _, v := range variables {
		for _, r := range results {
			results = append(results, r+v)
		}
	}
	return results
}

func intToBinary36(n int) [36]rune {
	result := [36]rune{}
	for i := 0; i < 36; i++ {
		pow := utils.PowInt(2, 35-i)
		if n-pow >= 0 {
			result[i] = '1'
			n -= pow
		} else {
			result[i] = '0'
		}
	}
	return result
}

func binary36ToInt(nbin [36]rune) int {
	result := 0
	for i, bit := range nbin {
		if bit == '1' {
			result += utils.PowInt(2, 35-i)
		}
	}
	return result
}
