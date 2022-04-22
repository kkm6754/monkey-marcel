package leetcode429

type Node struct {
    Val int
    Children []*Node
}


func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	level := make([]*Node, 0)
	level = append(level, root)
	for len(level) > 0 {
		levelVal := make([]int, 0)
		nextLevel := make([]*Node, 0)
		for _, node := range level {
			levelVal = append(levelVal, node.Val)
			if node.Children != nil && len(node.Children) > 0 {
				nextLevel = append(nextLevel, node.Children...)
			}
		}
		res = append(res, levelVal)
		level = nextLevel
	}

	return res

}
