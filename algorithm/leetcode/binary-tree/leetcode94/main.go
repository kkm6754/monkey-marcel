package leetcode94

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	cur := root
	stack := make([]*TreeNode, 0)

	for cur != nil || len(stack) > 0{
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack) - 1]
			res = append(res, cur.Val)
			stack = stack[: len(stack) - 1]
			cur = cur.Right
		}
	}

	return res
}