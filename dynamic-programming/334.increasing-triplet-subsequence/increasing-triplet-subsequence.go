package leetcode

import (
	"math"
)

func increasingTriplet(nums []int) bool {

	leftMin := make([]int, len(nums))
	rightMax := make([]int, len(nums))

	leftMin[0] = math.MaxInt
	rightMax[len(nums)-1] = math.MinInt

	for i := 1; i < len(nums); i++ {
		leftMin[i] = min(leftMin[i-1], nums[i-1])
	}

	for i := len(nums) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i+1])
	}

	// fmt.Println(leftMin)
	// fmt.Println(rightMax)

	for i := 0; i < len(nums); i++ {
		if leftMin[i] < nums[i] && nums[i] < rightMax[i] {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
