package leetcode

import (
	"fmt"
	"testing"
)

type question152 struct {
	para152
	ans152
}
type para152 struct {
	input []int
}

type ans152 struct {
	result int
}

func Test_Question152(t *testing.T) {

	qs := []question152{
		{
			para152{[]int{2, 3, -2, 4}},
			ans152{6},
		},
	}

	for _, q := range qs {
		_, p := q.ans152, q.para152
		fmt.Printf("input:%v output:%v\n", p, maxProduct(p.input))
	}
	fmt.Printf("\n\n\n")
}
