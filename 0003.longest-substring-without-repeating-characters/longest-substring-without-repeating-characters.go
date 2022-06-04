package leetcode

func lengthOfLongestSubstring(s string) int {

	i, j := 0, 0
	maxLength := 0

	hastTable := make(map[byte]int, len(s))

	for j < len(s) {
		if idx, ok := hastTable[s[j]]; ok && idx >= i {
			i = idx + 1
		}
		hastTable[s[j]] = j
		j++
		maxLength = max(maxLength, (j - i))
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
