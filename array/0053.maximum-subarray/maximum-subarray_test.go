package leetcode

import (
	"fmt"
	"testing"
)

type question53 struct {
	para53
	ans53
}

type para53 struct {
	input []int
}

type ans53 struct {
	result int
}

func Test_Question53(t *testing.T) {

	qs := []question53{

		{
			para53{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans53{6},
		},
	}

	for _, q := range qs {
		_, p := q.ans53, q.para53
		fmt.Printf("input:%v output:%v\n", p, maxSubArray(p.input))
	}
	fmt.Printf("\n\n\n")
}
