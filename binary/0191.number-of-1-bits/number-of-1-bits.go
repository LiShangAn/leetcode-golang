package leetcode

import "fmt"

/*
	5 > 101

	method1: use mod
		101%2 = 1
		100%2 = 0

	method2" use &
*/

func hammingWeight1(num uint32) int {
	count := 0

	for num > 0 {
		fmt.Printf("%b\n", num)
		if num%2 == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}

func hammingWeight2(num uint32) int {

	count := 0

	for num > 0 {
		fmt.Printf("%b\n", num)
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}
