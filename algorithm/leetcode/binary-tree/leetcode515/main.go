package leetcode515


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	level := make([]*TreeNode, 0)
	level = append(level, root)
	levelVal := make([]int, 0)
	levelVal = append(levelVal, root.Val)

	for len(level) > 0 {
		res = append(res, max(levelVal))

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


func max(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}