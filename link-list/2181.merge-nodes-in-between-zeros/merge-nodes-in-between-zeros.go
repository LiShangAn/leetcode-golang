package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {

	// init
	result := &ListNode{}
	h := result

	if head.Next == nil {
		return result
	}

	current := head
	sum := 0

	for current.Next != nil {
		if current.Next.Val != 0 {
			sum += current.Next.Val
		} else {
			h.Next = &ListNode{Val: sum, Next: nil}
			h = h.Next
			sum = 0
		}
		current = current.Next
	}

	return result.Next
}
