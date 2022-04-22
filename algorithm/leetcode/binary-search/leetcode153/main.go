package leetcode153

//153. 寻找旋转排序数组中的最小值

// 二分法
// 左闭右开 左边需要加一，因为左闭，包含最左的一个数

func findMin(nums []int) int {

	l, r := 0, len(nums)
	for l < r {
		m := l + (r - l)/2

		if nums[m] < nums[r] {
			r = m
		}else {
			l = m + 1
		}
	}

	return nums[l]

}