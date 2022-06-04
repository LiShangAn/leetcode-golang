package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	para3
	ans3
}

type para3 struct {
	input1 string
}

type ans3 struct {
	result1 int
}

func Test_Question3(t *testing.T) {
	qs := []question3{
		{
			para3{"abcabcbb"},
			ans3{3},
		},
		{
			para3{"bbbbb"},
			ans3{1},
		},
		{
			para3{"pwwkew"},
			ans3{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")
	for _, q := range qs {
		input := q.para3.input1
		expect := q.ans3.result1
		actual := lengthOfLongestSubstring(input)
		fmt.Printf("【input】:%v\n【expect】:%v\n【actual】:%v\n", input, expect, actual)
	}
}
