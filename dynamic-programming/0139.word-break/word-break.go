package leetcode

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)

	dp[0] = true
	n := len(s)
	for i := 0; i < len(s); i++ {

		if dp[i] {
			for _, word := range wordDict {

				if (n-i >= len(word)) && s[i:i+len(word)] == word {

					fmt.Println(s[i : i+len(word)])

					dp[i+len(word)] = true
				}
			}
		}
	}

	fmt.Println(dp)

	return dp[len(s)]
}
