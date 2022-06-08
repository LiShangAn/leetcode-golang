package leetcode

import (
	"fmt"
	"testing"
)

type question371 struct {
	para371
	ans371
}

type para371 struct {
	a int
	b int
}

type ans371 struct {
	result int
}

func Test_Question371(t *testing.T) {

	qs := []question371{
		{
			para371{5, 6},
			ans371{11},
		},
	}

	for _, q := range qs {
		_, p := q.ans371, q.para371
		fmt.Printf("input:%v output:%v\n", p, getSum(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
