package leetcode

import (
	"fmt"
	"testing"
)

type question160 struct {
	para160
	ans160
}

type para160 struct {
	one     []int
	another []int
}

type ans160 struct {
	one []int
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

func Test_Question160(t *testing.T) {

	qs := []question160{
		// {
		// 	para160{[]int{3}, []int{1, 2, 3}},
		// 	ans160{[]int{3}},
		// },

		// {
		// 	para160{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		// 	ans160{[]int{1, 2, 3, 4}},
		// },

		{
			para160{[]int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}},
			ans160{[]int{8, 4, 5}},
		},

		// {
		// 	para160{[]int{0, 9, 1, 2, 4}, []int{3, 2, 4}},
		// 	ans160{[]int{2, 4}},
		// },
	}

	for _, q := range qs {
		_, p := q.ans160, q.para160
		fmt.Printf("【input】:%v 【output】:%v\n", p, getIntersectionNode(Ints2List(p.one), Ints2List(p.another)))
	}
	fmt.Printf("\n\n\n")
}
