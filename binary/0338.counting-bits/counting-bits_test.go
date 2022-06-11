package leetcode

import (
	"fmt"
	"testing"
)

type question338 struct {
	para338
	ans338
}

type para338 struct {
	input int
}

type ans338 struct {
	result []int
}

func Test_Question338(t *testing.T) {

	qs := []question338{

		{
			para338{5},
			ans338{[]int{0, 1, 1, 2, 1, 2}},
		},
	}

	for _, q := range qs {
		_, p := q.ans338, q.para338
		fmt.Printf("input:%v output:%v\n", p, countBits(p.input))
	}
	fmt.Printf("\n\n\n")
}
