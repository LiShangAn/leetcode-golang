package leetcode

import (
	"fmt"
	"testing"
)

type question191 struct {
	para191
	ans191
}

type para191 struct {
	input uint32
}

type ans191 struct {
	result int
}

func Test_Question191(t *testing.T) {

	qs := []question191{

		{
			para191{5},
			ans191{1},
		},
	}

	for _, q := range qs {
		_, p := q.ans191, q.para191
		fmt.Printf("input:%v output:%v\n", p, hammingWeight1(p.input))
		fmt.Printf("input:%v output:%v\n", p, hammingWeight2(p.input))

	}
	fmt.Printf("\n\n\n")
}
