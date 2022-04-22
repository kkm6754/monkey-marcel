package main

import (
	"fmt"
)

// 二叉树遍历 DFS（深度优先遍历）
// 前序	根-》左-》右
// 后序 左-》右-》根
// 中序 左-》根-》右






// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if nil == root {
		return ret
	}

	traversal(root, &ret)
	return ret 
} 

// 递归 3要素
// 1. 定义
// 2. 出口
// 3. 大问题拆分成小问题递归


// 递归方式 不常考
func traversal(root *TreeNode, ret *[]int) {
	if nil == root {
		return
	}
	
	*ret = append(*ret, root.Val)
	traversal(root.Left, ret)
	traversal(root.Right, ret)
}



func main() {
	left := TreeNode {
		Val: 3,
		Left: nil,
		Right: nil,
	}
	right := TreeNode {
		Val: 2,
		Left: &left,
		Right: nil,
	}
	header := TreeNode {
		Val: 1,
		Left: nil,
		Right: &right,
	}
	ret := preorderTraversal(&header)
	fmt.Println(ret)
}