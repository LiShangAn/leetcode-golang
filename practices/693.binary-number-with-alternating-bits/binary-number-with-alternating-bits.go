package leetcode

func hasAlternatingBits(n int) bool {

	previous := n & 1
	n >>= 1

	for n >= 1 {
		if n&1 == previous {
			return false
		}
		previous = n & 1
		n >>= 1
	}

	return true
}
