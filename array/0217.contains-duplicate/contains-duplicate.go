package leetcode

func containsDuplicate(nums []int) bool {

	record := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		_, exist := record[nums[i]]

		if exist {
			return true
		}
		record[nums[i]] = true
	}
	return false
}
