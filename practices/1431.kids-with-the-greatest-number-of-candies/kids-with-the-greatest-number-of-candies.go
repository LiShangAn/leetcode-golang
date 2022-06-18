package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {

	result := []bool{}

	max := 0
	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}

	for _, candy := range candies {
		if candy+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
