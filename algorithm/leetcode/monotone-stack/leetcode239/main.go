package leetcode239

//239. 滑动窗口最大值


func maxSlidingWindow(nums []int, k int) []int {
	var stack []int
	var res []int

	for i, v := range nums {

		// 单调递减栈，最新位置的值比栈顶值要大时需要出栈
		for len(stack) > 0 && v >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

		// 最大值是否越过窗口
		if i-k+1 > stack[0] {
			stack = stack[1:]
		}

		// 窗口填满后，取栈底的最大值到结果中
		if i+1 >= k {
			res = append(res, nums[stack[0]])
		}
	}
	return res
}


//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
//1 [3  -1  -3] 5  3  6  7       3
//1  3 [-1  -3  5] 3  6  7       5
//1  3  -1 [-3  5  3] 6  7       5
//1  3  -1  -3 [5  3  6] 7       6
//1  3  -1  -3  5 [3  6  7]      7



