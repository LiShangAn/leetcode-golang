package leetcode

func maxProduct(nums []int) int {

	minValue := nums[0]
	maxValue := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		tmp := minValue
		minValue = min(min(minValue*nums[i], maxValue*nums[i]), nums[i])
		maxValue = max(max(tmp*nums[i], maxValue*nums[i]), nums[i])

		result = max(maxValue, result)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
