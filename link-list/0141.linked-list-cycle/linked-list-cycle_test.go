package leetcode

import (
	"fmt"
	"testing"
)

type question141 struct {
	para141
	ans141
}

type para141 struct {
	input []int
}

type ans141 struct {
	result bool
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l

	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func Test_Question141(t *testing.T) {

	qs := []question141{

		{
			para141{[]int{}},
			ans141{false},
		},
		// {
		// 	para141{[]int{3, 2, 0, -4}},
		// 	ans141{false},
		// },
	}

	for _, q := range qs {
		_, p := q.ans141, q.para141
		fmt.Printf("input:%v output:%v\n", p, hasCycle(Ints2List(p.input)))
	}
	fmt.Printf("\n\n\n")
}
