package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		res = append(res, queue[len(queue) - 1].Val)

		level := make([]*TreeNode, 0)
		for _, node := range queue {
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		queue = level
	}

	return res

}



func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: nil,
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
	}


	res := rightSideView(root)
	fmt.Println(res)
}
