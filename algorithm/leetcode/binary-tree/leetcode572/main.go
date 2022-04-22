package leetcode572

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		if check(cur, subRoot) {
			return true
		}
		stack = stack[: len(stack) - 1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}

	return false
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return check(p.Left, q.Left) && check(p.Right, q.Right)
}