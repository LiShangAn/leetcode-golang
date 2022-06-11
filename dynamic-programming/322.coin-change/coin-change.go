package leetcode

import "fmt"

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	fmt.Println(dp)

	// init slice
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 1; i <= amount; i++ {
		for c := 0; c < len(coins); c++ {
			if coins[c] <= i {
				dp[i] = min(dp[i-coins[c]]+1, dp[i])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
