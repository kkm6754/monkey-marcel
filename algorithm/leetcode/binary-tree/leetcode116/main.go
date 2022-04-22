package leetcode116

type Node struct {
     Val int
     Left *Node
     Right *Node
     Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	level := make([]*Node, 0)
	level = append(level, root)
	for len(level) > 0 {
		nextLevel := make([]*Node, 0)

		for i := 0; i < len(level); i++ {
			if i < len(level) - 1 {
				level[i].Next = level[i+1]
			}else {
				level[i].Next = nil
			}
			if level[i].Left != nil {
				nextLevel = append(nextLevel, level[i].Left)
			}
			if level[i].Right != nil {
				nextLevel = append(nextLevel, level[i].Right)
			}
		}
		level = nextLevel
	}

	return root
}