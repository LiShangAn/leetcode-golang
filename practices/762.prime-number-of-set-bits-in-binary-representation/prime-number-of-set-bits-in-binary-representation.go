package leetcode

import "fmt"

func countPrimeSetBits(left int, right int) int {

	result := 0

	for num := left; num <= right; num++ {
		count := countBit(num)

		fmt.Println(count)
		fmt.Println(isPrime(count))

		if isPrime(count) {
			result++
		}
	}
	return result
}

func countBit(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
