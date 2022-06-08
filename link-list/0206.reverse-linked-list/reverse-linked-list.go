package leetcode

// ListNode define
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var result *ListNode

	/*
		  [1,2,3,4,5]
			head = 1
			nextNode = 2
	*/

	for head != nil {
		nextNode := head.Next
		head.Next = result
		result = head
		head = nextNode
	}

	return result
}
