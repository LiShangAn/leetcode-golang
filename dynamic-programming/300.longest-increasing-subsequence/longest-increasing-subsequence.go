package leetcode

func lengthOfLIS(nums []int) int {

	dp := make([]int, len(nums))
	result := 0

	for i := 0; i < len(nums); i++ {
		// fmt.Println("------------------------")
		// fmt.Printf("i: %d, nums[%d]: %d\n", i, i, nums[i])

		for j := 0; j < i; j++ {
			// fmt.Printf("j: %d, nums[%d]: %d\n", j, j, nums[j])
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
				// fmt.Println(dp[i])
			}
		}

		dp[i] = dp[i] + 1
		// fmt.Println(dp[i])
		result = max(result, dp[i])
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
