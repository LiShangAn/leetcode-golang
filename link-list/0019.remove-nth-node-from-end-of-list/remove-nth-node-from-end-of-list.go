package leetcode

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	count := 0

	tmp := head
	for tmp != nil {
		count++
		tmp = tmp.Next
	}

	result := new(ListNode)
	result.Next = head
	current := result

	for i := 0; i < (count - n); i++ {
		current = current.Next
	}
	current.Next = current.Next.Next

	return result.Next
}
