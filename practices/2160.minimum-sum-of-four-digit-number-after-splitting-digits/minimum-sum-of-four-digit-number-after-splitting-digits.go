package leetcode

import "sort"

func minimumSum(num int) int {
	result := 0
	// 2 9 3 4
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})
	// 9 4 3 2
	for i, digit := range digits {
		if i < 2 {
			result = result + digit
		} else {
			result = result + 10*digit
		}
	}
	return result
}
