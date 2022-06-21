package leetcode

func subtractProductAndSum(n int) int {

	product := 1
	sum := 0

	for n > 0 {

		lastNum := n % 10

		product *= lastNum
		sum += lastNum
		n /= 10
	}

	return product - sum

}
