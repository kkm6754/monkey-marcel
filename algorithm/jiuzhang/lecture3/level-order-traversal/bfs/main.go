package main


import "fmt"


// 二叉树 层次遍历
// bfs 
// 队列实现


// 层次遍历 返回二维数据
// bfs 一维数据

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue)>0 {
		levelNode := make([]*TreeNode, 0)
		level := make([]int, 0)

		for _, treeNode := range queue {
			level = append(level, treeNode.Val)			
			if treeNode.Left != nil {
				levelNode = append(levelNode, treeNode.Left)
			}
			if treeNode.Right != nil {
				levelNode = append(levelNode, treeNode.Right)
			}
		}
		ret = append(ret, level)
		if len(levelNode) == 0 {
			break
		}
		queue = levelNode
	}

	return ret
}




func main() {
	left2 := TreeNode {
		Val: 4,
		Left: nil,
		Right: nil,
	}
	right := TreeNode {
		Val: 3,
		Left: &left2,
		Right: nil,
	}
	left := TreeNode {
		Val: 2,
		Left: nil,
		Right: nil,
	}
	header := TreeNode {
		Val: 1,
		Left: &left,
		Right: &right,
	}
	ret := levelOrder(&header)
	fmt.Println(ret)
}