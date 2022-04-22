package leetcode101


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
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

	return check(p.Left, q.Right) && check(q.Left, p.Right)
}