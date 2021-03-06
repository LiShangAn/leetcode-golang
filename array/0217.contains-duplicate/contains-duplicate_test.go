package leetcode

import (
	"fmt"
	"testing"
)

type question217 struct {
	para217
	ans217
}

type para217 struct {
	input []int
}

type ans217 struct {
	result bool
}

func Test_Question217(t *testing.T) {

	qs := []question217{

		{
			para217{[]int{1, 2, 3, 1}},
			ans217{true},
		},

		{
			para217{[]int{1, 2, 3, 4}},
			ans217{false},
		},

		{
			para217{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			ans217{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans217, q.para217
		fmt.Printf("input:%v output:%v\n", p, containsDuplicate(p.input))
	}
	fmt.Printf("\n\n\n")
}
