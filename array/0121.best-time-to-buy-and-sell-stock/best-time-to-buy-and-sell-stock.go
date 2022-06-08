package leetcode

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	max := 0

	/*
		[]int{7, 			1, 			5, 			3, 			6, 			4}

		i:	  0	1	2 6 4 5
		min:  7 1	1	1	1 1
		max:	0 0 4 4 5 5
		p[i]: 7	1 5 3 6 4

	*/
	for i := 0; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}
