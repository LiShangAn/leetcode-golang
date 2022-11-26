package golang

import "fmt"

func countBoat(input [][]byte) [3]int {

	result := [3]int{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {

			if input[i][j] == '#' {
				total := dfs(&input, i, j)

				if total >= 1 {
					result[total-1]++
				}
			}
		}
	}

	fmt.Println(input)

	return result
}

// func isInGrid(input *[][]byte, i, j int) bool {
// 	return i >= 0 && j >= 0 && i < len(*input) && j < len((*input)[i])
// }

func dfs(input *[][]byte, i, j int) int {

	// if !isInGrid(input, i, j) || string((*input)[i][j]) == "." {
	// 	return 0
	// }

	if !(i >= 0 && j >= 0 && i < len(*input) && j < len((*input)[i])) || ((*input)[i][j] != '#') {
		return 0
	}

	(*input)[i][j] = '.'
	// (*input)[i][j] = byte(count)

	total := 1

	total += dfs(input, i+1, j)
	total += dfs(input, i, j+1)
	total += dfs(input, i-1, j)
	total += dfs(input, i, j-1)

	return total
}
