package leetcode404



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	var sumOfLeft func(node *TreeNode, sum *int)
	sumOfLeft = func(node *TreeNode, sum *int) {
		if node.Left == nil && node.Right == nil {
			return
		}

		// 只处理左子节点
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil{
			*sum = *sum + node.Left.Val
			sumOfLeft(node.Left, sum)
		}
		if node.Left != nil {
			sumOfLeft(node.Left, sum)
		}
		if node.Right != nil {
			sumOfLeft(node.Right, sum)
		}
	}

	sumOfLeft(root, &res)

	return res
}
