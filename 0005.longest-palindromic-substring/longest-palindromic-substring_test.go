package leetcode

import (
	"fmt"
	"testing"
)

type question5 struct {
	para5
	ans5
}
type para5 struct {
	input1 string
}
type ans5 struct {
	result1 string
}

func Test_Question5(t *testing.T) {

	qs := []question5{
		// {
		// 	para5{"babad"},
		// 	ans5{"bab"},
		// },

		{
			para5{"cbabd"},
			ans5{"bb"},
		},
		// {
		// 	para5{"a"},
		// 	ans5{"a"},
		// },
		// {
		// 	para5{"2abcba3"},
		// 	ans5{"a"},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 5------------------------\n")
	for _, q := range qs {
		input := q.para5.input1
		expect := q.ans5.result1
		actual := longestPalindrome(input)
		fmt.Printf("【input】:%v\n【expect】:%v\n【actual】:%v\n", input, expect, actual)
	}
	fmt.Printf("\n\n\n")
}
