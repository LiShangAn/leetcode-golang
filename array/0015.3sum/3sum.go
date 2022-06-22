package leetcode

import "sort"

func threeSum(nums []int) [][]int {

	result := [][]int{}

	count := map[int]int{}

	for _, num := range nums {
		count[num]++
	}

	unique := []int{}
	for key := range count {
		unique = append(unique, key)
	}

	sort.Ints(unique)
	for i := 0; i < len(unique); i++ {

		uniqueSelf := unique[i]
		// 擁有自己3次
		if (uniqueSelf*3 == 0) && count[uniqueSelf] >= 3 {
			result = append(result, []int{uniqueSelf, uniqueSelf, uniqueSelf})
		}

		for j := i + 1; j < len(unique); j++ {
			uniqueOthers := unique[j]

			// 擁有自己2次 + 其他人1次
			if (uniqueSelf*2+uniqueOthers == 0) && count[uniqueSelf] >= 2 {
				result = append(result, []int{uniqueSelf, uniqueSelf, uniqueOthers})
			}
			// 擁有自己1次 + 其他人2次
			if (uniqueSelf+uniqueOthers*2 == 0) && count[uniqueOthers] >= 2 {
				result = append(result, []int{uniqueSelf, uniqueOthers, uniqueOthers})
			}

			// 擁有自己1次 + 其他人1次 + 境外移入1次
			outSide := 0 - uniqueSelf - uniqueOthers
			if outSide > uniqueOthers && count[outSide] > 0 {
				result = append(result, []int{uniqueSelf, uniqueOthers, outSide})
			}
		}
	}

	return result
}
