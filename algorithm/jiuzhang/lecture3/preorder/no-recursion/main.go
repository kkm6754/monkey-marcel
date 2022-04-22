package main

import "fmt"


// 前序遍历
// 根 左 右

// 非递归，利用栈


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

	stack := make([]*TreeNode, 0)
	stack = append(stack, root) // 先放根
	
	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		ret = append(ret, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {	// 先放右，后取
			stack = append(stack, node.Right)
		}
		if node.Left != nil { // 后放左，先去
			stack = append(stack, node.Left)
		}
	}

	return ret 
}


// 通过迭代模拟递归的先进后出来实现后序遍历
// 前中序遍历来说只需要改变当前节点和其左右子节点的入栈顺序即可
// 非常好的统一了三种遍历
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	// 先让根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top != nil {
			// 当前节点先入栈，这样出栈时就能保证比子节点后出栈
			// 并且对于已经处理过子节点的节点，会在其再次入栈的时候再压入一个 nil 作为标记
			// 下次遍历到 nil 就直接将 nil 下面的节点出栈
			stack = append(stack, top, nil)
			// 子节点入栈顺序是先右后左，保证出栈先左后右
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
		} else if len(stack) > 0 {
			// 已处理过子节点的节点直接出栈并加入输出
			top = stack[len(stack)-1]
			ans = append(ans, top.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return ans
}


func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode


	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}


func main() {
	left := TreeNode {
		Val: 2,
		Left: nil,
		Right: nil,
	}
	right := TreeNode {
		Val: 3,
		Left: nil,
		Right: nil,
	}
	header := TreeNode {
		Val: 1,
		Left: &left,
		Right: &right,
	}
	pre := preorderTraversal(&header)
	fmt.Println(pre)

	post := postorderTraversal(&header)
	fmt.Println(post)

	in := inorderTraversal(&header)
	fmt.Println(in)
}