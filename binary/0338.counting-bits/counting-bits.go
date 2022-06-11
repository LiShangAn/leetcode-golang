package leetcode

func countBits(n int) []int {

	result := make([]int, n+1)

	/*
	 0	0			0
	 1	1			1
	 2	10		1
	 3	11		2
	 4	100		1
	 5	101		2
	 6	110		2
	 7	111		3
	 8	1000	1
	*/
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + i&1
	}

	return result
}
