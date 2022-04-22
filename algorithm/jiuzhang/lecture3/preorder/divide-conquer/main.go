package main

import "fmt"


// 前序遍历
// 分治 递归


// 二叉树对问题，多可以拆分为左右子树上的问题
// 分治 可行性高


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	ret = append(ret, root.Val)
	ret = append(ret, left...)
	ret = append(ret, right...)

	return	ret
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