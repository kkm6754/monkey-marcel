package main


// 树形数据结构
//	堆


//给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：
//
//二叉树的根是数组 nums 中的最大元素。
//左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
//右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
//返回有给定数组 nums 构建的 最大二叉树 。


// 单调栈
//从前往后依次遍历数组元素，对每个元素：
//
// 如果栈内有元素且元素小于当前元素，那么就需要依次弹出栈内比当前元素小的元素。 在这个过程中记录最后一个被弹出的元素，这个元素就是当前元素的左儿子。
// 如果栈内有元素且栈顶元素大于当前元素，那么当前元素就是栈顶元素的右儿子。
// 最后压入当前元素。


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//func constructMaximumBinaryTree(nums []int) *TreeNode {
//
//}