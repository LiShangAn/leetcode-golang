package leetcode

import (
	"fmt"
	"testing"
)

type question206 struct {
	para206
	ans206
}

type para206 struct {
	input []int
}

type ans206 struct {
	result []int
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

func Test_Question206(t *testing.T) {

	qs := []question206{

		{
			para206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans206, q.para206
		fmt.Printf("input:%v output:%v\n", p, reverseList(Ints2List(p.input)))
	}
}
