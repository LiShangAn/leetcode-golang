package leetcode

import (
	"fmt"
	"testing"
)

type question322 struct {
	para322
	ans322
}

type para322 struct {
	input  []int
	amount int
}

type ans322 struct {
	result int
}

func Test_Question322(t *testing.T) {

	qs := []question322{
		{
			para322{[]int{1, 2, 5}, 11},
			ans322{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans322, q.para322
		fmt.Printf("input:%v output:%v\n", p, coinChange(p.input, p.amount))
	}
	fmt.Printf("\n\n\n")
}
