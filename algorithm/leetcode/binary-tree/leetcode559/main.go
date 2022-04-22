package leetcode559

type Node struct {
	Val int
	Children []*Node
}


func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	if root.Children != nil {
		for _, child := range root.Children {
			depth = max(depth, maxDepth(child))
		}
	}

	return depth + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

