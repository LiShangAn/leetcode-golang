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

func numComponents(head *ListNode, nums []int) int {

	cur := head

	exist := make(map[int]bool)

	for _, num := range nums {
		exist[num] = true
	}

	for cur != nil {
		if exist[cur.Val] {
			cur.Val = 1
		} else {
			cur.Val = 0
		}
		cur = cur.Next
	}

	cur = head

	result := 0

	for cur != nil {
		if cur.Val != 0 && (cur.Next != nil) {
			if cur.Next.Val == 0 {
				result++
			}
		} else if cur.Val != 0 && (cur.Next == nil) {
			result++
		}

		cur = cur.Next
	}

	return result
}
