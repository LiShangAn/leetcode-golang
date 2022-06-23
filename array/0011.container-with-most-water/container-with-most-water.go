package leetcode

import "fmt"

func maxArea(height []int) int {

	l := 0
	r := len(height) - 1

	result := 0

	for l <= r {

		min := getMin(height[l], height[r])
		area := min * (r - l)

		fmt.Printf(":l: %d, %d r: %d, %d    area: %d\n", l, height[l], r, height[r], area)

		if area > result {
			result = area
		}

		if height[r] > height[l] {
			l++
		} else if height[l] >= height[r] {
			r--
		}
	}
	return result
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
