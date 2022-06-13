package leetcode

import (
	"fmt"
	"testing"
)

type question19 struct {
	para19
	ans19
}

type para19 struct {
	input []int
	n     int
}

type ans19 struct {
	result []int
}

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

func Test_Question19(t *testing.T) {

	qs := []question19{

		{
			para19{[]int{1, 2}, 2},
			ans19{[]int{2}},
		},

		// {
		// 	para19{[]int{1, 2, 3, 4, 5}, 2},
		// 	ans19{[]int{1, 2, 3, 5}},
		// },
	}

	for _, q := range qs {
		_, p := q.ans19, q.para19
		fmt.Printf("input:%v output:%v\n", p, removeNthFromEnd(Ints2List(p.input), p.n))
	}
	fmt.Printf("\n\n\n")
}
