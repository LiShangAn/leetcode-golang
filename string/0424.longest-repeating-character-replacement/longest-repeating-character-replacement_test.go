package leetcode

import (
	"fmt"
	"testing"
)

type question424 struct {
	para424
	ans424
}

type para424 struct {
	s string
	k int
}

type ans424 struct {
	result int
}

func Test_Question424(t *testing.T) {

	qs := []question424{

		{
			para424{"AABABBA", 1},
			ans424{4},
		},
	}

	for _, q := range qs {
		_, p := q.ans424, q.para424
		fmt.Printf("input:%v output:%v\n", p, characterReplacement(p.s, p.k))
	}
}
