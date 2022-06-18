package leetcode

func maximumWealth(accounts [][]int) int {
	max := 0

	for i := 0; i < len(accounts); i++ {

		count := 0

		for j := 0; j < len(accounts[i]); j++ {
			count += accounts[i][j]
		}

		if count > max {
			max = count
		}
	}

	return max
}
