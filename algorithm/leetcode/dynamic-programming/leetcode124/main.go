package leetcode124

import "math"

// 124. 二叉树中的最大路径和


// dp[C] 表示当前节点为最大上升节点的最大路径和（路径中最靠近根节点，最高的节点）
// 初始值 从叶子节点开始计算
// dp[C] = max(max(dp[L], 0), max(dp[R], 0)) + C.val   C,L,R 分别代指 当前节点、左子节点、右子节点。
// 结果 当前节点为根节点的最大路径和




type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	dfs(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}



func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 后续遍历，先求左右子树
	// 小于0的节点，舍弃
	leftSubTreeSum := max(0, dfs(node.Left))
	rightSubTreeSum := max(0, dfs(node.Right))

	// node节点计算
	res = max(res, node.Val + leftSubTreeSum + rightSubTreeSum)


	// node节点计算
	return node.Val + max(leftSubTreeSum, rightSubTreeSum)
}