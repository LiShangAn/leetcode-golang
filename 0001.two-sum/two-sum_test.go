package golang

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

type para1 struct {
	nums   []int
	target int
}

type ans1 struct {
	one []int
}

func Test_Question1(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{2, 7, 11, 15}, 9},
			ans1{[]int{0, 1}},
		},
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},
		{
			para1{[]int{3, 3}, 6},
			ans1{[]int{0, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")
	for _, q := range qs {
		input := q.para1
		expect := q.ans1.one

		actual := twoSum(input.nums, input.target)
		fmt.Printf("【input】:%v\n【expect】:%v\n【actual】:%v\n", input, expect, actual)

	}
}
