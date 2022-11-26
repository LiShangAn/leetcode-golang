package leetcode

import "fmt"

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {

	// 方法二
	// dummyNode := new(ListNode)
	// preNode := dummyNode

	// 方法一
	// var dummyNode ListNode
	// preNode := &dummyNode

	// 方法三
	dummyNode := ListNode{}
	preNode := &dummyNode

	fmt.Println(dummyNode)
	fmt.Println(*preNode)

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			preNode.Next = list1
			list1 = list1.Next
		} else {
			preNode.Next = list2
			list2 = list2.Next
		}
		preNode = preNode.Next
	}
	if list1 == nil {
		preNode.Next = list2
	} else {
		preNode.Next = list1
	}

	return dummyNode.Next
}
