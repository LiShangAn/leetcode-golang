package leetcode

func smallerNumbersThanCurrent(nums []int) []int {

	result := []int{}
	for _, num := range nums {

		count := 0
		for _, compare := range nums {
			if num > compare {
				count++
			}
		}
		result = append(result, count)
	}

	return result
}
