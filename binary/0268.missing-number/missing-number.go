package leetcode

func missingNumber(nums []int) int {

	start := 0

	for i := 1; i <= len(nums); i++ { // start:0  i:1,2,3   n=3: 0,1,2,3
		start = start ^ i ^ nums[i-1]
	}

	return start
}
