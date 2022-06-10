package leetcode

import (
	"fmt"
	"testing"
)

type question238 struct {
	para238
	ans238
}

type para238 struct {
	input []int
}

type ans238 struct {
	result []int
}

func Test_Question238(t *testing.T) {

	qs := []question238{
		{
			para238{[]int{1, 2, 3, 4}},
			ans238{[]int{24, 12, 8, 6}},
		},
	}

	for _, q := range qs {
		fmt.Printf("input:%v output:%v\n", q.input, productExceptSelf(q.input))
	}
	fmt.Printf("\n\n\n")
}
