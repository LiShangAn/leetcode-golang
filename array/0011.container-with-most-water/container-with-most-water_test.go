package leetcode

import (
	"fmt"
	"testing"
)

type question11 struct {
	para11
	ans11
}

type para11 struct {
	input []int
}

type ans11 struct {
	result int
}

func Test_Question11(t *testing.T) {

	qs := []question11{

		// {
		// 	para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		// 	ans11{49},
		// },

		{
			para11{[]int{1, 1}},
			ans11{1},
		},
	}

	for _, q := range qs {
		_, p := q.ans11, q.para11
		fmt.Printf("input:%v output:%v\n", p.input, maxArea(p.input))
	}
	fmt.Printf("\n\n\n")
}
