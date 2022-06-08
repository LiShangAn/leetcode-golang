package leetcode

import (
	"fmt"
	"testing"
)

type question70 struct {
	para70
	ans70
}

type para70 struct {
	input int
}

type ans70 struct {
	result int
}

func Test_Question70(t *testing.T) {

	qs := []question70{

		{
			para70{2},
			ans70{2},
		},

		{
			para70{3},
			ans70{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans70, q.para70
		fmt.Printf("input:%v output:%v\n", p, climbStairs(p.input))
	}
	fmt.Printf("\n\n\n")
}
