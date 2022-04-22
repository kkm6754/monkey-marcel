package leetcode107


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	levelNode := make([][]*TreeNode, 0)
	level := make([]*TreeNode, 0)
	level = append(level, root)
	for len(level) > 0 {
		levelNode = append(levelNode, level)
		nextLevel := make([]*TreeNode, 0)
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		level = nextLevel
	}

	for i := len(levelNode) - 1; i >= 0; i-- {
		level := levelNode[i]
		levelRes := make([]int, 0)
		for _, node := range level {
			levelRes = append(levelRes, node.Val)
		}
		res = append(res, levelRes)
	}

	return res
}