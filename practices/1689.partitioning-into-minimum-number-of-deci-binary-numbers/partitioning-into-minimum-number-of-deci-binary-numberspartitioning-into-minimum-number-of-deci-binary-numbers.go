package leetcode

import "strconv"

func minPartitions1(n string) int {

	result := 0
	for i := 0; i < len(n); i++ {
		result = max(result, int(n[i]-'0'))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minPartitions2(n string) int {

	result := 0
	for i := 0; i < len(n); i++ {
		num, _ := strconv.Atoi(string(n[i]))

		if num > result {
			result = num
		}
	}

	return result
}
