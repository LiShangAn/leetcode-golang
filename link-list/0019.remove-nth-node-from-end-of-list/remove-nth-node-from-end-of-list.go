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

	dummyNode := new(ListNode)
	dummyNode.Next = head
	currentNode := dummyNode

	count := 0
	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		count++
	}

	for i := 0; i < (count - n); i++ {
		currentNode = currentNode.Next
	}
	currentNode.Next = currentNode.Next.Next

	return dummyNode.Next
}
