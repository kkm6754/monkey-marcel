package leetcode222

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		res++
		cur := stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	return res
}
