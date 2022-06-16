package leetcode

import "fmt"

func isAnagram(s string, t string) bool {
	tmp := map[rune]int{}

	for _, v := range s {
		tmp[v]++
	}

	for _, v := range t {
		tmp[v]--
	}

	for _, v := range tmp {
		fmt.Println(v)
		if v != 0 {
			return false
		}
	}

	return true
}
