package leetcode

func findMin1(nums []int) int {

	low := 0

	high := len(nums) - 1

	for low < high {
		median := (low + high) / 2

		if nums[low] < nums[high] {
			return nums[low]
		}

		if nums[median] > nums[high] {
			low = median + 1
		} else {
			high = median
		}
	}

	return nums[low]

}

func findMin2(nums []int) int {

	min := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		}
	}

	return min
}
