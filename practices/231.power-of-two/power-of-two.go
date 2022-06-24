package leetcode

func isPowerOfTwo(n int) bool {

	tmp := 1

	for tmp < n {
		tmp <<= 1
	}

	return tmp == n
}
