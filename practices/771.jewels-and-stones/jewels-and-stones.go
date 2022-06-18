package leetcode

func numJewelsInStones1(jewels string, stones string) int {

	res := make(map[byte]bool)
	count := 0

	for i := 0; i < len(jewels); i++ {
		res[byte(jewels[i])] = true
	}

	for i := 0; i < len(stones); i++ {
		if res[byte(stones[i])] {
			count++
		}
	}
	return count
}

func numJewelsInStones2(jewels string, stones string) int {

	count := 0
	for i := 0; i < len(stones); i++ {

		for j := 0; j < len(jewels); j++ {
			if stones[i] == jewels[j] {
				count++
			}
		}
	}
	return count
}
