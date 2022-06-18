package leetcode

func mostWordsFound(sentences []string) int {

	max := 0

	for i := 0; i < len(sentences); i++ {
		count := 0

		for j := 0; j < len(sentences[i]); j++ {
			if sentences[i][j] == ' ' {
				count++
			}
		}

		if count > max {
			max = count
		}
	}
	return max + 1
}
