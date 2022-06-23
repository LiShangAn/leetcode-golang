package leetcode

import "sort"

// 解法 1
func fourSum1(nums []int, target int) [][]int {

	result := [][]int{}
	count := map[int]int{}
	unique := []int{}

	for _, num := range nums {
		count[num]++
	}

	for key := range count {
		unique = append(unique, key)
	}

	sort.Ints(unique)
	for i := 0; i < len(unique); i++ {

		uniqueSelf := unique[i]

		if uniqueSelf*4 == target && count[uniqueSelf] >= 4 {
			result = append(result, []int{uniqueSelf, uniqueSelf, uniqueSelf, uniqueSelf})
		}

		for j := i + 1; j < len(unique); j++ {

			uniqueOther := unique[j]

			if uniqueSelf*3+uniqueOther == target && count[uniqueSelf] >= 3 && count[uniqueOther] >= 1 {
				result = append(result, []int{uniqueSelf, uniqueSelf, uniqueSelf, uniqueOther})
			}
			if uniqueSelf*2+uniqueOther*2 == target && count[uniqueSelf] >= 2 && count[uniqueOther] >= 2 {
				result = append(result, []int{uniqueSelf, uniqueSelf, uniqueOther, uniqueOther})
			}
			if uniqueSelf+uniqueOther*3 == target && count[uniqueSelf] >= 1 && count[uniqueOther] >= 3 {
				result = append(result, []int{uniqueSelf, uniqueOther, uniqueOther, uniqueOther})
			}

			for k := j + 1; k < len(unique); k++ {

				uniqueThird := unique[k]

				if uniqueSelf*2+uniqueOther+uniqueThird == target && count[uniqueSelf] >= 2 && count[uniqueOther] >= 1 && count[uniqueThird] >= 1 {
					result = append(result, []int{uniqueSelf, uniqueSelf, uniqueOther, uniqueThird})
				}
				if uniqueSelf*1+uniqueOther*2+uniqueThird == target && count[uniqueSelf] >= 1 && count[uniqueOther] >= 2 && count[uniqueThird] >= 1 {
					result = append(result, []int{uniqueSelf, uniqueOther, uniqueOther, uniqueThird})
				}
				if uniqueSelf*1+uniqueOther*1+uniqueThird*2 == target && count[uniqueSelf] >= 1 && count[uniqueOther] >= 1 && count[uniqueThird] >= 2 {
					result = append(result, []int{uniqueSelf, uniqueOther, uniqueThird, uniqueThird})
				}

				outSide := target - uniqueSelf - uniqueOther - uniqueThird

				if outSide > uniqueThird && count[outSide] > 0 {
					result = append(result, []int{uniqueSelf, uniqueOther, uniqueThird, outSide})
				}
			}
		}
	}

	return result
}
