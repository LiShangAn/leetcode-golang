package leetcode

import (
	"fmt"
	"testing"
)

type question21 struct {
	para21
	ans21
}

type para21 struct {
	input1 []int
	input2 []int
}

type ans21 struct {
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

func Test_Question21(t *testing.T) {

	qs := []question21{
		{
			para21{[]int{1, 2, 4}, []int{1, 3, 4}},
			ans21{[]int{1, 1, 2, 3, 4, 4}},
		},
	}

	for _, q := range qs {
		_, p := q.ans21, q.para21
		fmt.Printf("input:%v output:%v\n", p, mergeTwoLists(Ints2List(p.input1), Ints2List(p.input2)))
	}
	fmt.Printf("\n\n\n")
}
