package leetcode

func hammingDistance(x int, y int) int {

	tmp := x ^ y

	count := 0

	for tmp >= 1 {
		count += tmp & 1
		tmp >>= 1
	}

	return count
}
