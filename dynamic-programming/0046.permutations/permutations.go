package leetcode

func permute(nums []int) [][]int {

	result := [][]int{}

	used := make([]bool, len(nums))

	path := []int{}

	n := len(nums)
	if n == 0 {
		return result
	}

	result = dfs(nums, n, 0, path, used, result)

	return result
}

func dfs(nums []int, length int, depth int, path []int, used []bool, result [][]int) [][]int {

	if length == depth {
		result = append(result, path)
		return result
	}

	for i := 0; i < length; i++ {
		if used[i] {
			continue
		} else {
			newPath := append(path, nums[i])
			used[i] = true
			result = dfs(nums, length, depth+1, newPath, used, result)
			used[i] = false
		}
	}

	return result
}
