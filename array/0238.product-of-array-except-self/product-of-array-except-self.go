package leetcode

import "fmt"

func productExceptSelf(nums []int) []int {

	// answ 24,12,8,6
	// nums 1, 2, 3, 4

	// left to right
	// res >1, 1, 2, 6

	// right to left
	// res <24, 12,8,6

	res := make([]int, len(nums))
	res[0] = 1

	// left to right
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	fmt.Println(res)

	right := 1
	//right to left
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right = nums[i] * right
	}

	return res

}
