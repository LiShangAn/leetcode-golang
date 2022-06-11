package leetcode

func maxSubArray(nums []int) int {

	result := nums[0]

	res := make([]int, len(nums))
	res[0] = 0

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			res[i] = res[i] + nums[i]
		} else {
			res[i] = res[i-1] + nums[i]
		}

		result = max(res[i], result)
		if res[i] < 0 {
			res[i] = 0
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
