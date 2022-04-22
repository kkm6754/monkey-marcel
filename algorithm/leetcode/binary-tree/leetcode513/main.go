package leetcode513


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func findBottomLeftValue(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		res = queue[0].Val
		nextLevel := make([]*TreeNode, 0)
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
	}

	return res
}


func findBottom(node *TreeNode, bottom *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return
	}
}