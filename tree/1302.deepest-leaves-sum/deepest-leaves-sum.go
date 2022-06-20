package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {

	sum, maxDeep := 0, 0

	deepestValue(root, &sum, 0, &maxDeep)

	return sum
}

func deepestValue(root *TreeNode, sum *int, deep int, maxDeep *int) {

	if root == nil {
		return
	}

	if deep > *maxDeep {
		*maxDeep = deep
		*sum = root.Val
	} else if deep == *maxDeep {
		*sum += root.Val
	}

	deepestValue(root.Right, sum, deep+1, maxDeep)
	deepestValue(root.Left, sum, deep+1, maxDeep)
}
