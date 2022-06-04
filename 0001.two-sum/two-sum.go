package golang

import "fmt"

func twoSum(nums []int, target int) []int {

	fmt.Println(nums)
	storage := make(map[int]int, len(nums))

	for i, num := range nums {
		if idx, ok := storage[target-num]; ok {
			return []int{idx, i}
		}
		storage[num] = i
	}
	return []int{0, 0}
}
