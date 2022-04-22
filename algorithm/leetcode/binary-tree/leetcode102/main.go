package leetcode102


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelRes := make([]int, 0)
		nextLevel := make([]*TreeNode, 0)

		for _, node := range queue {
			levelRes = append(levelRes, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		res = append(res, levelRes)
		queue = nextLevel
	}

	return res
}
