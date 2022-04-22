package leetcode637


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)
	if root == nil {
		return res
	}

	level := make([]*TreeNode, 0)
	level = append(level, root)
	levelVal := make([]int, 0)
	levelVal = append(levelVal, root.Val)

	for len(level) > 0 {
		res = append(res, avg(levelVal))

		nextLevel := make([]*TreeNode, 0)
		nextVal := make([]int, 0)
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
				nextVal = append(nextVal, node.Left.Val)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
				nextVal = append(nextVal, node.Right.Val)
			}
		}
		level = nextLevel
		levelVal = nextVal
	}

	return res
}

func avg(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	var sum int
	for _, v := range nums {
		sum = sum + v
	}
	return float64(sum) / float64(len(nums))
}
